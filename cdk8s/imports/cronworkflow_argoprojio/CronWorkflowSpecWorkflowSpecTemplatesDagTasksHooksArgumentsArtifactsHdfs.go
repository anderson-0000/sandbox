package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	Addresses *[]*string `field:"optional" json:"addresses" yaml:"addresses"`
	Force *bool `field:"optional" json:"force" yaml:"force"`
	HdfsUser *string `field:"optional" json:"hdfsUser" yaml:"hdfsUser"`
	KrbCCacheSecret *CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret `field:"optional" json:"krbCCacheSecret" yaml:"krbCCacheSecret"`
	KrbConfigConfigMap *CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap `field:"optional" json:"krbConfigConfigMap" yaml:"krbConfigConfigMap"`
	KrbKeytabSecret *CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret `field:"optional" json:"krbKeytabSecret" yaml:"krbKeytabSecret"`
	KrbRealm *string `field:"optional" json:"krbRealm" yaml:"krbRealm"`
	KrbServicePrincipalName *string `field:"optional" json:"krbServicePrincipalName" yaml:"krbServicePrincipalName"`
	KrbUsername *string `field:"optional" json:"krbUsername" yaml:"krbUsername"`
}

