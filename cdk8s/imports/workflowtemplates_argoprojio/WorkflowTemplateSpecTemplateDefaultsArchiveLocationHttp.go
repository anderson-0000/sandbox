package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

