package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3 struct {
	AccessKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3AccessKeySecret `field:"optional" json:"accessKeySecret" yaml:"accessKeySecret"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	CaSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CaSecret `field:"optional" json:"caSecret" yaml:"caSecret"`
	CreateBucketIfNotPresent *CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CreateBucketIfNotPresent `field:"optional" json:"createBucketIfNotPresent" yaml:"createBucketIfNotPresent"`
	EncryptionOptions *CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptions `field:"optional" json:"encryptionOptions" yaml:"encryptionOptions"`
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	Key *string `field:"optional" json:"key" yaml:"key"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	SecretKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3SecretKeySecret `field:"optional" json:"secretKeySecret" yaml:"secretKeySecret"`
	UseSdkCreds *bool `field:"optional" json:"useSdkCreds" yaml:"useSdkCreds"`
}

