package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

