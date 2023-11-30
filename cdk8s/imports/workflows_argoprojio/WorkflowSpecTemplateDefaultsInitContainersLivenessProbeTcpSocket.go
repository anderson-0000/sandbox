package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

