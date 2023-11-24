package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplateDefaultsContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

