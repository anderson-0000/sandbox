package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesSidecarsEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplatesSidecarsEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

