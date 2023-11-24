package workflows _argoprojio


type WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

