package workflows_argoprojio


type WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

