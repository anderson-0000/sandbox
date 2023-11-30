package workflows_argoprojio


type WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

