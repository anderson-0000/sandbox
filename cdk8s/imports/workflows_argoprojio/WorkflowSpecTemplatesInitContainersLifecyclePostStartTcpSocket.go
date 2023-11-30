package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocket struct {
	Port WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

