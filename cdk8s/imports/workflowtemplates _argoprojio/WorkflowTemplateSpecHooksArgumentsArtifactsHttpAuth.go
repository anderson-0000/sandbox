package workflowtemplates _argoprojio


type WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuth struct {
	BasicAuth *WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

