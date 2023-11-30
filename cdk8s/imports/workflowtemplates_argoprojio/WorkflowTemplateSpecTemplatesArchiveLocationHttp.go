package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesArchiveLocationHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplatesArchiveLocationHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplatesArchiveLocationHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

