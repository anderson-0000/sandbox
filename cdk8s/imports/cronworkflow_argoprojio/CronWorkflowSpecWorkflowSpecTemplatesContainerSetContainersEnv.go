package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

