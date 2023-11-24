package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesInputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplatesInputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

