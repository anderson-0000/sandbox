package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuth struct {
	BasicAuth *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

