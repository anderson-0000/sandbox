package workflows_argoprojio


type WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

