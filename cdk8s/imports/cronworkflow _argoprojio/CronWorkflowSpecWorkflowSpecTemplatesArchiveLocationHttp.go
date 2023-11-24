package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

