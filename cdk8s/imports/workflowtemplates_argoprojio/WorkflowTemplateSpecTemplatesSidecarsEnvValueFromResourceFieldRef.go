package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

