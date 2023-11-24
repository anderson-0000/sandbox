package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesScriptEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplatesScriptEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

