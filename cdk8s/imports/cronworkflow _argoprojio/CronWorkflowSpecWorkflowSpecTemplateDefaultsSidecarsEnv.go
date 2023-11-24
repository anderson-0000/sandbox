package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

