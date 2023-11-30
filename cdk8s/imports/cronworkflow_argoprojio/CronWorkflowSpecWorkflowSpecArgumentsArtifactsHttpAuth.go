package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuth struct {
	BasicAuth *CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

