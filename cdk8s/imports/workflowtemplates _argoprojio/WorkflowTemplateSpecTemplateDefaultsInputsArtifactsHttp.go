package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

