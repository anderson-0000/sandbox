package workflows_argoprojio


type WorkflowSpecTemplatesContainerReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplatesContainerReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

