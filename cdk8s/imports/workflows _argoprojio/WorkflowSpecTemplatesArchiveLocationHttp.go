package workflows _argoprojio


type WorkflowSpecTemplatesArchiveLocationHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplatesArchiveLocationHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplatesArchiveLocationHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

