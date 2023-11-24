package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesArchiveLocationHttpAuth struct {
	BasicAuth *WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

