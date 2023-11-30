package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcs struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	ServiceAccountKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret `field:"optional" json:"serviceAccountKeySecret" yaml:"serviceAccountKeySecret"`
}

