package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

