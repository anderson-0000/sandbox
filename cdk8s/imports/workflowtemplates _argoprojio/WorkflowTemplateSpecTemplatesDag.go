package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesDag struct {
	Tasks *[]*WorkflowTemplateSpecTemplatesDagTasks `field:"required" json:"tasks" yaml:"tasks"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	Target *string `field:"optional" json:"target" yaml:"target"`
}

