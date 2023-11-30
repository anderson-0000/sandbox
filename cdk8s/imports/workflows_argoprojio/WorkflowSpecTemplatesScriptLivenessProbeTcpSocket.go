package workflows_argoprojio


type WorkflowSpecTemplatesScriptLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesScriptLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

