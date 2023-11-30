package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooks struct {
	Arguments *CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

