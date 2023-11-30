package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth struct {
	BasicAuth *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	ClientCert *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	Oauth2 *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

