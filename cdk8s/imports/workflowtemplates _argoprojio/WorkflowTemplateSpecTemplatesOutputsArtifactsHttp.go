package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesOutputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

