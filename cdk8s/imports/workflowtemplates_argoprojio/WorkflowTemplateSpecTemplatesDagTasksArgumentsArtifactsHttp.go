package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

