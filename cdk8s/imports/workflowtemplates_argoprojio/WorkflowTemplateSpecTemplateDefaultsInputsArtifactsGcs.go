package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

