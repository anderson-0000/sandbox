package workflows_argoprojio


type WorkflowSpecTemplatesContainerLifecyclePreStopTcpSocket struct {
	Port WorkflowSpecTemplatesContainerLifecyclePreStopTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

