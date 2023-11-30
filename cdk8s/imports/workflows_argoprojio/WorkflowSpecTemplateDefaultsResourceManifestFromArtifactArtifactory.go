package workflows_argoprojio


type WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

