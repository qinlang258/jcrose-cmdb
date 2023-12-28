package model

type K8sNodeListOutput struct {
	Name             string
	Status           string
	Age              string
	Version          string
	InternalIp       string
	OsImage          string
	KernelVersion    string
	ContainerRuntime string
}
