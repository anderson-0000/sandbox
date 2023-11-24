package workflows _argoprojio


type WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

