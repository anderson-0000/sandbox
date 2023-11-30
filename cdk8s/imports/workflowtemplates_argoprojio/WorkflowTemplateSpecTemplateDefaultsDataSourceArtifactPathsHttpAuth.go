package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth struct {
	BasicAuth *WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

