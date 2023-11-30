package workflows_argoprojio


type WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth struct {
	BasicAuth *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

