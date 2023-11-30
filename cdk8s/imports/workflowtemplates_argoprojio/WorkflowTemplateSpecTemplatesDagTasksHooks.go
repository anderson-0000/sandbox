package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesDagTasksHooks struct {
	Arguments *WorkflowTemplateSpecTemplatesDagTasksHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *WorkflowTemplateSpecTemplatesDagTasksHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

