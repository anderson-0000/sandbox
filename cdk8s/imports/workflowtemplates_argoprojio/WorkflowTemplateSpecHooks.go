package workflowtemplates_argoprojio


type WorkflowTemplateSpecHooks struct {
	Arguments *WorkflowTemplateSpecHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *WorkflowTemplateSpecHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

