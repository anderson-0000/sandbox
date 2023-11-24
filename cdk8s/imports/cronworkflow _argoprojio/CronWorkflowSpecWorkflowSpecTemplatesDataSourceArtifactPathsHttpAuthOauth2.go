package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2 struct {
	ClientIdSecret *CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

