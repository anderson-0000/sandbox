package workflows_argoprojio


type WorkflowSpecArgumentsArtifactsHttpAuth struct {
	BasicAuth *WorkflowSpecArgumentsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecArgumentsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecArgumentsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

