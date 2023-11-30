package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesArchiveLocationGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *WorkflowTemplateSpecTemplatesArchiveLocationGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

