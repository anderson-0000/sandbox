package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

