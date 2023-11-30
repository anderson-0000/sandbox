package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

