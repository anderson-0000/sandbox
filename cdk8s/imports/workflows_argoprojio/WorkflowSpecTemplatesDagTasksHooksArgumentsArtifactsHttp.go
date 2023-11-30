package workflows_argoprojio


type WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

