package workflows_argoprojio


type WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

