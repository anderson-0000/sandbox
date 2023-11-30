package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

