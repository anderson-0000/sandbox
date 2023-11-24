package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

