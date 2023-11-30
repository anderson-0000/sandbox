package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3 struct {
	AccessKeySecret *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3AccessKeySecret `field:"optional" json:"accessKeySecret" yaml:"accessKeySecret"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	CaSecret *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3CaSecret `field:"optional" json:"caSecret" yaml:"caSecret"`
	CreateBucketIfNotPresent *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3CreateBucketIfNotPresent `field:"optional" json:"createBucketIfNotPresent" yaml:"createBucketIfNotPresent"`
	EncryptionOptions *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptions `field:"optional" json:"encryptionOptions" yaml:"encryptionOptions"`
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	Key *string `field:"optional" json:"key" yaml:"key"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	SecretKeySecret *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3SecretKeySecret `field:"optional" json:"secretKeySecret" yaml:"secretKeySecret"`
	UseSdkCreds *bool `field:"optional" json:"useSdkCreds" yaml:"useSdkCreds"`
}

