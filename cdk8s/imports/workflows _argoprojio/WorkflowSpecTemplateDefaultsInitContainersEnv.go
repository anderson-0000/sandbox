package workflows _argoprojio


type WorkflowSpecTemplateDefaultsInitContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplateDefaultsInitContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

