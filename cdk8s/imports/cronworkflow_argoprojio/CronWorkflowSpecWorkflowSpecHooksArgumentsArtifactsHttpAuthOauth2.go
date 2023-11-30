package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

