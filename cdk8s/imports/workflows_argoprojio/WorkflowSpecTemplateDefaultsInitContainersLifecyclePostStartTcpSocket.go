package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

