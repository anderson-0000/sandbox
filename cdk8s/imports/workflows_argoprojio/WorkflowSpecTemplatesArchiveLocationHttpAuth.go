package workflows_argoprojio


type WorkflowSpecTemplatesArchiveLocationHttpAuth struct {
	BasicAuth *WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplatesArchiveLocationHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

