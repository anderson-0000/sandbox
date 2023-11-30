package workflows_argoprojio


type WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuth struct {
	BasicAuth *WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

