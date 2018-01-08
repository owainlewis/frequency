package executor

import (
	"github.com/golang/glog"
	types "github.com/owainlewis/kcd/pkg/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

// Should be configurable
const shell = "/bin/bash"

// Executor controls how jobs are executed inside Kubernetes
type Executor struct {
	Client kubernetes.Interface
}

// NewExecutor creates a properly configured Job Executor
func NewExecutor(clientset kubernetes.Interface) Executor {
	return Executor{Client: clientset}
}

// Execute will execute a single job
func (e Executor) Execute(job *types.Job) (*v1.Pod, error) {

	job.EnsureDefaults()

	glog.Infof("Executing job: ", job.Name)
	template := e.NewPod(job.Workspace, job.Image, job.Commands)

	// TODO which namespace to run in (must be configurable)
	pod, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(template)
	if err != nil {
		glog.Infof("Failed to create pod")
		return nil, err
	}

	return pod, nil
}

// FormatCommands will format the commands passed for execution in a shell
// Hacky way to do this. TODO can we refine here?
func (e Executor) FormatCommands(commands []string) []string {
	cmdStr := ""
	for _, c := range commands {
		cmdStr = cmdStr + c + ";"
	}

	return []string{shell, "-c", cmdStr}
}

// NewPod will construct the pod template to be created to run a job
func (e Executor) NewPod(workspace string, image string, commands []string) *v1.Pod {
	defaultEnv := []v1.EnvVar{
		{
			Name:  "WORKSPACE",
			Value: workspace,
		},
		{
			Name:  "OUTPUT_DIR",
			Value: "/output",
		},
	}
	primary := v1.Container{
		Name:       "primary",
		Image:      image,
		WorkingDir: workspace,
		Env:        defaultEnv,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "workspace",
				MountPath: workspace,
			},
			{
				Name:      "output",
				MountPath: "/output",
			}},
		Command: e.FormatCommands(commands),
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"owner": "kcd",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{primary},
			InitContainers: []v1.Container{{
				Name:  "setup",
				Image: "alpine/git",
				Env:   defaultEnv,
				VolumeMounts: []v1.VolumeMount{{
					Name:      "workspace",
					MountPath: workspace,
				}},
				Command: []string{
					"ash", "-c", "git clone https://github.com/owainlewis/hello-spinnaker.git $WORKSPACE",
				},
			}},
			RestartPolicy: v1.RestartPolicyNever,
			Volumes: []v1.Volume{
				{
					Name: "workspace",
				},
				{
					Name: "output",
				},
			},
		},
	}

	pod.SetGenerateName("kcd-")

	return pod

}
