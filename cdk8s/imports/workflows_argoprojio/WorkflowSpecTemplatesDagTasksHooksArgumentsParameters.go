package workflows_argoprojio


type WorkflowSpecTemplatesDagTasksHooksArgumentsParameters struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Default *string `field:"optional" json:"default" yaml:"default"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Enum *[]*string `field:"optional" json:"enum" yaml:"enum"`
	GlobalName *string `field:"optional" json:"globalName" yaml:"globalName"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

