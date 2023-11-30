package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

