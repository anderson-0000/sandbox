package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

