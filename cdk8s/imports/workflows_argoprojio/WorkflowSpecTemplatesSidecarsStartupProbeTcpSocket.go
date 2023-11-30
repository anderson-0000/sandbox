package workflows_argoprojio


type WorkflowSpecTemplatesSidecarsStartupProbeTcpSocket struct {
	Port WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

