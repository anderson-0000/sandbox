package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3EncryptionOptions struct {
	EnableEncryption *bool `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	KmsEncryptionContext *string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	ServerSideCustomerKeySecret *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret `field:"optional" json:"serverSideCustomerKeySecret" yaml:"serverSideCustomerKeySecret"`
}

