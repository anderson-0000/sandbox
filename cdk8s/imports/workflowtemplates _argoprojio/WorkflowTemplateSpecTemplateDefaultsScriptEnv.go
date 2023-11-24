package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsScriptEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

