package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

