package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsInitContainersEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

