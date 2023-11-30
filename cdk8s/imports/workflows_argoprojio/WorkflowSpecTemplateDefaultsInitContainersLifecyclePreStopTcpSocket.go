package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

