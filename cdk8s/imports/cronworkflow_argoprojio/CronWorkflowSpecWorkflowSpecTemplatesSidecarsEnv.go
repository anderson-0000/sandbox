package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

