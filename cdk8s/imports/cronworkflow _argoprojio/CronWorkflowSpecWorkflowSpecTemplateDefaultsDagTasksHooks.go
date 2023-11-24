package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooks struct {
	Arguments *CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
}

