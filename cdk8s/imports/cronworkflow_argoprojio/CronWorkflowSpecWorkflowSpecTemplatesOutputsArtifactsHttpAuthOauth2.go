package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

