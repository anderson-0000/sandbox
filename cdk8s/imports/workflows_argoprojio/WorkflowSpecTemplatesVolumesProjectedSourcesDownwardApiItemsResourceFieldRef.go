package workflows_argoprojio


type WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

