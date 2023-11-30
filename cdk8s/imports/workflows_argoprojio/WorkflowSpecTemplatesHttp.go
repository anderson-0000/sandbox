package workflows_argoprojio


type WorkflowSpecTemplatesHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Body *string `field:"optional" json:"body" yaml:"body"`
	BodyFrom *WorkflowSpecTemplatesHttpBodyFrom `field:"optional" json:"bodyFrom" yaml:"bodyFrom"`
	Headers *[]*WorkflowSpecTemplatesHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
	InsecureSkipVerify *bool `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	Method *string `field:"optional" json:"method" yaml:"method"`
	SuccessCondition *string `field:"optional" json:"successCondition" yaml:"successCondition"`
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}
