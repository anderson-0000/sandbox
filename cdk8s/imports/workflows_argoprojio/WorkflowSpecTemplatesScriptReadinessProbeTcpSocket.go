package workflows_argoprojio


type WorkflowSpecTemplatesScriptReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesScriptReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

