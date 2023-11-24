package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesHttpHeaders struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplatesHttpHeadersValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

