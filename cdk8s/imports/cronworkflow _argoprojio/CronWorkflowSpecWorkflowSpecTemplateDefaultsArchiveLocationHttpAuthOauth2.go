package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2 struct {
	ClientIdSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

