package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplatesInitContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

