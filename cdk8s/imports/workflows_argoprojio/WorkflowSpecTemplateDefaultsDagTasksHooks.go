package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDagTasksHooks struct {
	Arguments *WorkflowSpecTemplateDefaultsDagTasksHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *WorkflowSpecTemplateDefaultsDagTasksHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

