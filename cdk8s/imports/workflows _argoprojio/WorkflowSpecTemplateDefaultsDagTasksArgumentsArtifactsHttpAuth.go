package workflows _argoprojio


type WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth struct {
	BasicAuth *WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

