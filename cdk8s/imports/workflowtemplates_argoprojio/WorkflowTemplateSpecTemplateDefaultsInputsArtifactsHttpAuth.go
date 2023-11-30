package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuth struct {
	BasicAuth *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

