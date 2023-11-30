package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

