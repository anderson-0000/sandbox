package workflows _argoprojio


type WorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

