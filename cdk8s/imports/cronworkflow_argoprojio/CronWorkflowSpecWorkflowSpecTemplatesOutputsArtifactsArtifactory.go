package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

