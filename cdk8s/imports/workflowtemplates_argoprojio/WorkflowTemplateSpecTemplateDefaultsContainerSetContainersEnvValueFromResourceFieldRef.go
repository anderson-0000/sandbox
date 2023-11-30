package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

