package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

