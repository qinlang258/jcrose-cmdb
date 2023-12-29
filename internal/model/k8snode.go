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

type K8sNodeDescribeOutput struct {
	Name              string
	Labels            []K8sNodeLabelOutput
	Annotations       []K8sNodeAnnotationsOutput
	CreationTimestamp string
	Taints            []K8sNodeTaintsOutput
	PodCIDR           string
}

type K8sNodeLabelOutput struct {
}

type K8sNodeAnnotationsOutput struct {
}

type K8sNodeTaintsOutput struct {
}

type K8sNodeConditionsOutput struct {
	Type               string
	Status             string
	LastHeartbeatTime  string
	LastTransitionTime string
	Reason             string
	Message            string
}

type K8sNodeAddressesOutput struct {
	InternalIP string
	Hostname   string
}

type K8sNodeAllocatableOutput struct {
	Cpu        int
	VolumeSize int
	Memory     int
	Pods       int
}

type K8sNodeSystemInfo struct {
}
