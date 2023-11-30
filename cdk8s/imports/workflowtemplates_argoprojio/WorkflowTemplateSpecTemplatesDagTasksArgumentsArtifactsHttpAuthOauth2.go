package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

