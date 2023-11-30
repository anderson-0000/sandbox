package workflows_argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocationHttpAuth struct {
	BasicAuth *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

