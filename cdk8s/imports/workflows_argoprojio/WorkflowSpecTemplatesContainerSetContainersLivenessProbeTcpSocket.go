package workflows_argoprojio


type WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

