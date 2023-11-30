package workflows_argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocationGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *WorkflowSpecTemplateDefaultsArchiveLocationGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

