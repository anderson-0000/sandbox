package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

