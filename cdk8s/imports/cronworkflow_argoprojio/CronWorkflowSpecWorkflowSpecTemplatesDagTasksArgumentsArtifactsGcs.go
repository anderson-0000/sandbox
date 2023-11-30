package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

