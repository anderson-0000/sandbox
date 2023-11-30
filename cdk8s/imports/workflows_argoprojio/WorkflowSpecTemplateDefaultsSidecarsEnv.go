package workflows_argoprojio


type WorkflowSpecTemplateDefaultsSidecarsEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplateDefaultsSidecarsEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

