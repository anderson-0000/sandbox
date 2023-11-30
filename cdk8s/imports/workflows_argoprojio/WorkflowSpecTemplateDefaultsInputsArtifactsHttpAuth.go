package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuth struct {
	BasicAuth *WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

