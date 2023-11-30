package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

