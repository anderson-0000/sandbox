package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesInitContainersLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

