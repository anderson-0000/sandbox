package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDag struct {
	Tasks *[]*WorkflowSpecTemplateDefaultsDagTasks `field:"required" json:"tasks" yaml:"tasks"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	Target *string `field:"optional" json:"target" yaml:"target"`
}

