package workflows _argoprojio


type WorkflowSpecTemplatesContainerEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplatesContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

