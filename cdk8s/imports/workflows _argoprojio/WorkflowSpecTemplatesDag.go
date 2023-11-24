package workflows _argoprojio


type WorkflowSpecTemplatesDag struct {
	Tasks *[]*WorkflowSpecTemplatesDagTasks `field:"required" json:"tasks" yaml:"tasks"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	Target *string `field:"optional" json:"target" yaml:"target"`
}

