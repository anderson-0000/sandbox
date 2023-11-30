package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

