package executor

import (
	"github.com/golang/glog"
	types "github.com/owainlewis/kcd/pkg/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

// outputDir is the location that build assets are moved to
// typically this will be a shared Kubernetes volume so that
// assets can be shared between jobs
const outputDir = "/output"

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
	glog.Infof("Executing job: %s", job.Name)
	template := e.NewJobExecutionPod(job)

	// TODO which namespace to run in (must be configurable)
	pod, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(template)
	if err != nil {
		glog.Infof("Failed to create pod")
		return nil, err
	}

	return pod, nil
}

func env(k, v string) v1.EnvVar {
	return v1.EnvVar{Name: k, Value: v}
}

// NewJobExecutionPod will construct the pod template in which a job will run
//
// The pod has two primary roles
//
// 1. A sidecar container will check out the project from a VCS
// 2. A primary pod is created with the correct shared directories and environment
//
func (e Executor) NewJobExecutionPod(job *types.Job) *v1.Pod {
	defaultEnv := []v1.EnvVar{
		env("WORKSPACE", job.Workspace),
		env("OUTPUT_DIR", outputDir),
	}
	primary := v1.Container{
		Name:       "primary",
		Image:      job.Image,
		WorkingDir: job.Workspace,
		Env:        defaultEnv,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "workspace",
				MountPath: job.Workspace,
			},
			{
				Name:      "output",
				MountPath: outputDir,
			}},
		Command: append([]string{job.Command.Cmd}, job.Command.Args...),
	}

	// When a source is declared as part of a job, we use an init container
	// to go and fetch that source code from a VCS such as github.com
	var initContainers []v1.Container
	if job.Source != nil {

		buildEnv := []v1.EnvVar{
			env("GIT_URL", job.Source.GitURL),
			env("GIT_BRANCH", job.Source.GitBranch),
		}

		sourceCloneContainer := v1.Container{
			Name:  "setup",
			Image: "alpine/git",
			Env:   append(defaultEnv, buildEnv...),
			VolumeMounts: []v1.VolumeMount{{
				Name:      "workspace",
				MountPath: job.Workspace,
			}},
			Command: []string{
				"ash", "-c", "git clone -b + " + job.Source.GitURL + " $WORKSPACE",
			}}

		initContainers = append(initContainers, sourceCloneContainer)
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"owner": "kcd",
			},
		},
		Spec: v1.PodSpec{
			Containers:     []v1.Container{primary},
			InitContainers: initContainers,
			RestartPolicy:  v1.RestartPolicyNever,
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
