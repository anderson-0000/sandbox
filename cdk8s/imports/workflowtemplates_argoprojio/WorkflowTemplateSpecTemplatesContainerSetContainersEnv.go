package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerSetContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

