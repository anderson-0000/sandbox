package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

