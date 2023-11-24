package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2 struct {
	ClientIdSecret *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret `field:"optional" json:"clientIdSecret" yaml:"clientIdSecret"`
	ClientSecretSecret *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret `field:"optional" json:"clientSecretSecret" yaml:"clientSecretSecret"`
	EndpointParams *[]*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams `field:"optional" json:"endpointParams" yaml:"endpointParams"`
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	TokenUrlSecret *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret `field:"optional" json:"tokenUrlSecret" yaml:"tokenUrlSecret"`
}

