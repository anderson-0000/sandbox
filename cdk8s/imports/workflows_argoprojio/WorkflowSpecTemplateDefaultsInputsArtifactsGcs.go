package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInputsArtifactsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *WorkflowSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

