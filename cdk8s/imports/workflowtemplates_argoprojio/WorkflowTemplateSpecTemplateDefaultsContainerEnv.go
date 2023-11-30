package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

