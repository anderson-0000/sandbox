package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

