package workflows _argoprojio


type WorkflowSpecHooks struct {
	Arguments *WorkflowSpecHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *WorkflowSpecHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

