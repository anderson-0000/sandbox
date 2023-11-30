package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocket struct {
	Port WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

