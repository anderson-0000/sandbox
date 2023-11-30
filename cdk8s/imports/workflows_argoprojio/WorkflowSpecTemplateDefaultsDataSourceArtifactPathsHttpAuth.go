package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth struct {
	BasicAuth *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

