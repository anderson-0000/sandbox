package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptions struct {
	EnableEncryption *bool `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	KmsEncryptionContext *string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	ServerSideCustomerKeySecret *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret `field:"optional" json:"serverSideCustomerKeySecret" yaml:"serverSideCustomerKeySecret"`
}

