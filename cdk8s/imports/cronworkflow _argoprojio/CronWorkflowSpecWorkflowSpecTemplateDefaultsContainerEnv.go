package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

