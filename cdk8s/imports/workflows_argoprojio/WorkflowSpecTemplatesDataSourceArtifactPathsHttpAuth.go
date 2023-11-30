package workflows_argoprojio


type WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuth struct {
	BasicAuth *WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

