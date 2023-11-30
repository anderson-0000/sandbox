package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

