package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

