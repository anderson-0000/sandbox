package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

