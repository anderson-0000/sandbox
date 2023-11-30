package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

