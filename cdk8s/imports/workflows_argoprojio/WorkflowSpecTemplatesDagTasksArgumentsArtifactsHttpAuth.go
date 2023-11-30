package workflows_argoprojio


type WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuth struct {
	BasicAuth *WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

