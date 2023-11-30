package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

