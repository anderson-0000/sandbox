package workflows_argoprojio


type WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

