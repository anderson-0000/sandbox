package workflowtemplates _argoprojio


type WorkflowTemplateSpecHooksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecHooksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

