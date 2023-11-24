package workflows _argoprojio


type WorkflowSpecTemplatesDagTasksArgumentsArtifactsAzure struct {
	Blob *string `field:"required" json:"blob" yaml:"blob"`
	Container *string `field:"required" json:"container" yaml:"container"`
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	AccountKeySecret *WorkflowSpecTemplatesDagTasksArgumentsArtifactsAzureAccountKeySecret `field:"optional" json:"accountKeySecret" yaml:"accountKeySecret"`
	UseSdkCreds *bool `field:"optional" json:"useSdkCreds" yaml:"useSdkCreds"`
}

