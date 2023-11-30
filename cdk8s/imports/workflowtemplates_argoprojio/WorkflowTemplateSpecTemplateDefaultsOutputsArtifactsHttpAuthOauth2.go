package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

