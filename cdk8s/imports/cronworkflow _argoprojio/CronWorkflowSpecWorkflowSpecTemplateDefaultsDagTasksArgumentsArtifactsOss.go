package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOss struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	AccessKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssAccessKeySecret `field:"optional" json:"accessKeySecret" yaml:"accessKeySecret"`
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	CreateBucketIfNotPresent *bool `field:"optional" json:"createBucketIfNotPresent" yaml:"createBucketIfNotPresent"`
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	LifecycleRule *CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssLifecycleRule `field:"optional" json:"lifecycleRule" yaml:"lifecycleRule"`
	SecretKeySecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssSecretKeySecret `field:"optional" json:"secretKeySecret" yaml:"secretKeySecret"`
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	UseSdkCreds *bool `field:"optional" json:"useSdkCreds" yaml:"useSdkCreds"`
}

