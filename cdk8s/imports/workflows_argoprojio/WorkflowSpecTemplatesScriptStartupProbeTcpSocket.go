package workflows_argoprojio


type WorkflowSpecTemplatesScriptStartupProbeTcpSocket struct {
	Port WorkflowSpecTemplatesScriptStartupProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

