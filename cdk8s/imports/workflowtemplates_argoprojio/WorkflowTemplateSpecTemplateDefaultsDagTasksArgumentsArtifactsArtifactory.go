package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

