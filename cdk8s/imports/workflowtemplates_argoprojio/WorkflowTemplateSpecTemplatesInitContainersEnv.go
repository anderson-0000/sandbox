package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesInitContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplatesInitContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

