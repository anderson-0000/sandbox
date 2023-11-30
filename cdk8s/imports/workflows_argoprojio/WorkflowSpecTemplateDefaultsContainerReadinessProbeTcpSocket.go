package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

