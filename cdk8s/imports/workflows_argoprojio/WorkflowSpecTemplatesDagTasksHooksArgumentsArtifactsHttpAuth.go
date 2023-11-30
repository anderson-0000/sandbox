package workflows_argoprojio


type WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth struct {
	BasicAuth *WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

