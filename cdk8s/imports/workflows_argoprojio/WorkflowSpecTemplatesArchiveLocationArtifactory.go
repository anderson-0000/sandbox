package workflows_argoprojio


type WorkflowSpecTemplatesArchiveLocationArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowSpecTemplatesArchiveLocationArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplatesArchiveLocationArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

