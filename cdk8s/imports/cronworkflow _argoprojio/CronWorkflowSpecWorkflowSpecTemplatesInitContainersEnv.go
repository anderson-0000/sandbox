package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

