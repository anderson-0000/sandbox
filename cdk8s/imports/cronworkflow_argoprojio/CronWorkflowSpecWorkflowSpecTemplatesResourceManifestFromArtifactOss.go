package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOss struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	AccessKeySecret *CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssAccessKeySecret `field:"optional" json:"accessKeySecret" yaml:"accessKeySecret"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	CreateBucketIfNotPresent *bool `field:"optional" json:"createBucketIfNotPresent" yaml:"createBucketIfNotPresent"`
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	LifecycleRule *CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssLifecycleRule `field:"optional" json:"lifecycleRule" yaml:"lifecycleRule"`
	SecretKeySecret *CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssSecretKeySecret `field:"optional" json:"secretKeySecret" yaml:"secretKeySecret"`
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	UseSdkCreds *bool `field:"optional" json:"useSdkCreds" yaml:"useSdkCreds"`
}

