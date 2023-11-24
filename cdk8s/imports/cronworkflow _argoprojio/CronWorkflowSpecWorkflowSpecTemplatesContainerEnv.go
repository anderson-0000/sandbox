package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

