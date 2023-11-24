package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

