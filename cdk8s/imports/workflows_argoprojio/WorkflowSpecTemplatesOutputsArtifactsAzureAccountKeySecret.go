package workflows_argoprojio


type WorkflowSpecTemplatesOutputsArtifactsAzureAccountKeySecret struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

