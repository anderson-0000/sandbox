package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuth struct {
	BasicAuth *WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

