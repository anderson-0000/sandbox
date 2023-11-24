package workflows _argoprojio


type WorkflowSpecTemplatesContainerLifecyclePostStartTcpSocket struct {
	Port WorkflowSpecTemplatesContainerLifecyclePostStartTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

