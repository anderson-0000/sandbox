package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

