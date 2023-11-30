package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

