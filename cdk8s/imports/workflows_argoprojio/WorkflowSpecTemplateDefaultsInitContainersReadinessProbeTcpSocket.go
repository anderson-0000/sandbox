package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocket struct {
	Port WorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

