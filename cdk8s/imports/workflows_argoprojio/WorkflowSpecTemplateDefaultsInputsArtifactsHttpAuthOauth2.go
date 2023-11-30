package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

