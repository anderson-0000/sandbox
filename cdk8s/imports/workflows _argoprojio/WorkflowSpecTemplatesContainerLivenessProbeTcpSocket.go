package workflows _argoprojio


type WorkflowSpecTemplatesContainerLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

