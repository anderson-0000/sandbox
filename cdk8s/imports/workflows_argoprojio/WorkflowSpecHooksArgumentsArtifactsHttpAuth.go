package workflows_argoprojio


type WorkflowSpecHooksArgumentsArtifactsHttpAuth struct {
	BasicAuth *WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecHooksArgumentsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

