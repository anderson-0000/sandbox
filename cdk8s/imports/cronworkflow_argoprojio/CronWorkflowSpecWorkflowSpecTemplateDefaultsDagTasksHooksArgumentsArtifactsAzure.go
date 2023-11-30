package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzure struct {
	Blob *string `field:"required" json:"blob" yaml:"blob"`
	Container *string `field:"required" json:"container" yaml:"container"`
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	AccountKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzureAccountKeySecret `field:"optional" json:"accountKeySecret" yaml:"accountKeySecret"`
	UseSdkCreds *bool `field:"optional" json:"useSdkCreds" yaml:"useSdkCreds"`
}

