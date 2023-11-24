package workflows _argoprojio


type WorkflowSpecTemplateDefaultsDagTasks struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Arguments *WorkflowSpecTemplateDefaultsDagTasksArguments `field:"optional" json:"arguments" yaml:"arguments"`
	ContinueOn *WorkflowSpecTemplateDefaultsDagTasksContinueOn `field:"optional" json:"continueOn" yaml:"continueOn"`
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	Depends *string `field:"optional" json:"depends" yaml:"depends"`
	Hooks *map[string]*WorkflowSpecTemplateDefaultsDagTasksHooks `field:"optional" json:"hooks" yaml:"hooks"`
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	OnExit *string `field:"optional" json:"onExit" yaml:"onExit"`
	Template *string `field:"optional" json:"template" yaml:"template"`
	TemplateRef *WorkflowSpecTemplateDefaultsDagTasksTemplateRef `field:"optional" json:"templateRef" yaml:"templateRef"`
	When *string `field:"optional" json:"when" yaml:"when"`
	WithItems *[]interface{} `field:"optional" json:"withItems" yaml:"withItems"`
	WithParam *string `field:"optional" json:"withParam" yaml:"withParam"`
	WithSequence *WorkflowSpecTemplateDefaultsDagTasksWithSequence `field:"optional" json:"withSequence" yaml:"withSequence"`
}

