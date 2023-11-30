package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

