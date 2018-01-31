package executor

// import (
// 	"github.com/golang/glog"
// 	v1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	kubernetes "k8s.io/client-go/kubernetes"

// 	tasks "github.com/owainlewis/frequency/pkg/tasks"
// )

// type CITaskExecutor struct {
// 	Client kubernetes.Interface
// }

// // NewCITaskExecutor creates a properly configured CITaskExecutor
// func NewCITaskExecutor(clientset kubernetes.Interface) CITaskExecutor {
// 	return CITaskExecutor{Client: clientset}
// }

// // Execute will execute a single job
// func (e CITaskExecutor) Execute(task tasks.CITask) error {
// 	glog.Infof("Executing CI task: %+v", task)
// 	taskPod := e.newPod(task)

// 	// TODO which namespace to run in (must be configurable)
// 	_, err := e.Client.CoreV1().Pods(v1.NamespaceDefault).Create(taskPod)
// 	return err
// }

// func env(k, v string) v1.EnvVar {
// 	return v1.EnvVar{Name: k, Value: v}
// }

// func (e CITaskExecutor) newPod(task tasks.CITask) *v1.Pod {
// 	primary := v1.Container{
// 		Name:       "primary",
// 		Image:      task.Image,
// 		WorkingDir: task.Workspace,
// 		Env:        task.Env,
// 		Command:    task.Run.Command,
// 		Args:       task.Run.Args,
// 	}

// 	// When a source is declared as part of a job, we use an init container
// 	// // to go and fetch that source code from a VCS such as github.com
// 	var initContainers []v1.Container
// 	// if job.Source != nil {
// 	// 	command := buildGitCloneCommmand(job.Source.GitURL, job.Source.GitBranch)

// 	sourceCloneContainer := v1.Container{
// 		Name:  "source",
// 		Image: "alpine/git",
// 		// 		Env:   environment,
// 		VolumeMounts: []v1.VolumeMount{{
// 			Name:      "workspace",
// 			MountPath: task.Workspace,
// 		}},
// 		Command: []string{"ash", "-c", "ls -la"}}

// 	initContainers = append(initContainers, sourceCloneContainer)

// 	pod := &v1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Labels: map[string]string{},
// 		},
// 		Spec: v1.PodSpec{
// 			Containers:     []v1.Container{primary},
// 			InitContainers: initContainers,
// 			RestartPolicy:  v1.RestartPolicyNever,
// 		},
// 	}

// 	pod.SetGenerateName("task-")

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
