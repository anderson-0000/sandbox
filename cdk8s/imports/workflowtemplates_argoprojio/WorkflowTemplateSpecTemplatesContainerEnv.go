package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplatesContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

