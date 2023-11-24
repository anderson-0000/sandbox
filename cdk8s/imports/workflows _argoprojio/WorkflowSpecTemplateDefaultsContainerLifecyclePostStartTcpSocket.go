package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

