package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptions struct {
	EnableEncryption *bool `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	KmsEncryptionContext *string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	ServerSideCustomerKeySecret *CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret `field:"optional" json:"serverSideCustomerKeySecret" yaml:"serverSideCustomerKeySecret"`
}

