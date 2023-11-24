package workflows _argoprojio


type WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuth struct {
	BasicAuth *WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

