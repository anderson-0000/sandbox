package workflows _argoprojio


type WorkflowSpecTemplatesScriptEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplatesScriptEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

