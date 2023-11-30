package workflows_argoprojio


type WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

