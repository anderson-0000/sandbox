package workflows _argoprojio


type WorkflowSpecTemplatesSidecarsEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplatesSidecarsEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

