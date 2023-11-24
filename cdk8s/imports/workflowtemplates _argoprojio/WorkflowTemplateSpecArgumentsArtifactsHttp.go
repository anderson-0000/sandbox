package workflowtemplates _argoprojio


type WorkflowTemplateSpecArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

