package workflows _argoprojio


type WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

