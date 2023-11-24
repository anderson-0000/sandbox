package workflows _argoprojio


type WorkflowSpecTemplatesContainerStartupProbeTcpSocket struct {
	Port WorkflowSpecTemplatesContainerStartupProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

