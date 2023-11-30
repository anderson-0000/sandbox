package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

