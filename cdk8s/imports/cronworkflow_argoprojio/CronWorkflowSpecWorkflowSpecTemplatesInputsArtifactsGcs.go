package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

