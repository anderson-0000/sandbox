package workflows_argoprojio


type WorkflowSpecTemplatesSidecarsReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesSidecarsReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

