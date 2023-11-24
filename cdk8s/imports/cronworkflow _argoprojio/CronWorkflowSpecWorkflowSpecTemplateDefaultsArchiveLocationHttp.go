package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

