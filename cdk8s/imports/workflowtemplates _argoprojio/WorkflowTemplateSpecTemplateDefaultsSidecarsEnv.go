package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsSidecarsEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

