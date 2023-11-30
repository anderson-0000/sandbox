package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesArchiveLocationArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowTemplateSpecTemplatesArchiveLocationArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowTemplateSpecTemplatesArchiveLocationArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

