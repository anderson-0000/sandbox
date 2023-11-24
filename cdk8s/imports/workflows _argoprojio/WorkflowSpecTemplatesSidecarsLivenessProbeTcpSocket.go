package workflows _argoprojio


type WorkflowSpecTemplatesSidecarsLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesSidecarsLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

