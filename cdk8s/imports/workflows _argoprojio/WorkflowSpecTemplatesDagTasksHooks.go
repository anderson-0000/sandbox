package workflows _argoprojio


type WorkflowSpecTemplatesDagTasksHooks struct {
	Arguments *WorkflowSpecTemplatesDagTasksHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *WorkflowSpecTemplatesDagTasksHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

