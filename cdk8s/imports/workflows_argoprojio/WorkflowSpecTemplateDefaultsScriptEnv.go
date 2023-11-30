package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplateDefaultsScriptEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

