package workflows_argoprojio


type WorkflowSpecTemplatesOutputsArtifactsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

