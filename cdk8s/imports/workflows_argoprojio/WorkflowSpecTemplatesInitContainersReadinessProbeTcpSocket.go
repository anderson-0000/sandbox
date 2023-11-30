package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesInitContainersReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

