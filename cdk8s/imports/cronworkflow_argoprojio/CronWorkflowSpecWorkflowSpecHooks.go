package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecHooks struct {
	Arguments *CronWorkflowSpecWorkflowSpecHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *CronWorkflowSpecWorkflowSpecHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

