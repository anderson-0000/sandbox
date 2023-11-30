package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

