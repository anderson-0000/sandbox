package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

