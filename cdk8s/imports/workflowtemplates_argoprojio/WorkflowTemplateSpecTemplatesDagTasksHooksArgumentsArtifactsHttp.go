package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

