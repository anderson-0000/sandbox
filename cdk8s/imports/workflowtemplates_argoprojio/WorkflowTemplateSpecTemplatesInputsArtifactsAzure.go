package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesInputsArtifactsAzure struct {
	Blob *string `field:"required" json:"blob" yaml:"blob"`
	Container *string `field:"required" json:"container" yaml:"container"`
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	AccountKeySecret *WorkflowTemplateSpecTemplatesInputsArtifactsAzureAccountKeySecret `field:"optional" json:"accountKeySecret" yaml:"accountKeySecret"`
	UseSdkCreds *bool `field:"optional" json:"useSdkCreds" yaml:"useSdkCreds"`
}
