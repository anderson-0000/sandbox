package workflows_argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptions struct {
	EnableEncryption *bool `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	KmsEncryptionContext *string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	ServerSideCustomerKeySecret *WorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret `field:"optional" json:"serverSideCustomerKeySecret" yaml:"serverSideCustomerKeySecret"`
}

