package workflows _argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocationHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplateDefaultsArchiveLocationHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

