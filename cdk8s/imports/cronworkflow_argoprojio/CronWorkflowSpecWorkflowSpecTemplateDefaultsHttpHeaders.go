package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeaders struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeadersValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

