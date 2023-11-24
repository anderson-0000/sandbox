package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

