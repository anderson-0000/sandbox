package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

