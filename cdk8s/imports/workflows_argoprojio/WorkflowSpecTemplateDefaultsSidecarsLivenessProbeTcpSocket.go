package workflows_argoprojio


type WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

