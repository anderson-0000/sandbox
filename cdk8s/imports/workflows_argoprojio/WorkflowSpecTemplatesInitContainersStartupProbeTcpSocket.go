package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersStartupProbeTcpSocket struct {
	Port WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

