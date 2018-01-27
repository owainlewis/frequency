package executor

// import (
// 	"github.com/golang/glog"
// 	v1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	kubernetes "k8s.io/client-go/kubernetes"

// 	tasks "github.com/owainlewis/frequency/pkg/tasks"
// )

// // PodTaskExecutor ...
// type PodTaskExecutor struct {
// 	Client kubernetes.Interface
// }

// // NewPodTaskExecutor creates a properly configured PodTaskExecutor
// func NewPodTaskExecutor(clientset kubernetes.Interface) PodTaskExecutor {
// 	return PodTaskExecutor{Client: clientset}
// }

// // Execute will execute a single job
// func (e PodTaskExecutor) Execute(task tasks.PodTask) error {
// 	glog.Infof("Executing Pod task: %+v", task)
// 	taskPOd := e.newPod(task)

// 	// TODO which namespace to run in (must be configurable)
// 	pod, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(taskPod)
// 	if err != nil {
// 		glog.Infof("Failed to create Pod: %s", err)
// 		return err
// 	}

// 	return nil
// }

// func env(k, v string) v1.EnvVar {
// 	return v1.EnvVar{Name: k, Value: v}
// }

// func (e Executor) newPod(task *tasks.PodTask) *v1.Pod {
// 	primary := v1.Container{
// 		Name:       "primary",
// 		Image:      task.Image,
// 		WorkingDir: task.Workspace,
// 		Env:        task.Env,
// 		Command:    task.Command,
// 		Args:       task.Args,
// 	}

// 	// // When a source is declared as part of a job, we use an init container
// 	// // to go and fetch that source code from a VCS such as github.com
// 	// var initContainers []v1.Container
// 	// if job.Source != nil {
// 	// 	command := buildGitCloneCommmand(job.Source.GitURL, job.Source.GitBranch)

// 	// 	sourceCloneContainer := v1.Container{
// 	// 		Name:  "setup",
// 	// 		Image: "alpine/git",
// 	// 		Env:   environment,
// 	// 		VolumeMounts: []v1.VolumeMount{{
// 	// 			Name:      "workspace",
// 	// 			MountPath: job.Workspace,
// 	// 		}},
// 	// 		Command: []string{"ash", "-c", command}}

// 	// 	initContainers = append(initContainers, sourceCloneContainer)
// 	// }

// 	pod := &v1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Labels: map[string]string{},
// 		},
// 		Spec: v1.PodSpec{
// 			Containers: []v1.Container{primary},
// 			//InitContainers: initContainers,
// 			RestartPolicy: v1.RestartPolicyNever,
// 		},
// 	}

// 	pod.SetGenerateName("frequency-")

// 	return pod

// }

// // func podEnvironmentForJob(job *types.Job) []v1.EnvVar {
// // 	environment := []v1.EnvVar{
// // 		env("WORKSPACE", job.Workspace),
// // 		env("OUTPUT_DIR", outputDir),
// // 	}

// // 	if job.Source != nil {
// // 		buildEnv := []v1.EnvVar{
// // 			env("GIT_URL", job.Source.GitURL),
// // 			env("GIT_BRANCH", job.Source.GitBranch),
// // 		}
// // 		environment = append(environment, buildEnv...)
// // 	}

// // 	return append(environment, job.Env...)
// // }

// // func buildGitCloneCommmand(gitURL string, gitBranch string) string {
// // 	return fmt.Sprintf("git clone %s $WORKSPACE", gitURL)
// // }
