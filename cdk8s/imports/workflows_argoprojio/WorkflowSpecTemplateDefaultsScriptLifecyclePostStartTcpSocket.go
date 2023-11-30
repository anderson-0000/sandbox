package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

