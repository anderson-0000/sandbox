package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

