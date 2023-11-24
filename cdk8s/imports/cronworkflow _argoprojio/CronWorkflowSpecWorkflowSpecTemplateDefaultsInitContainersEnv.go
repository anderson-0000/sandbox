package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

