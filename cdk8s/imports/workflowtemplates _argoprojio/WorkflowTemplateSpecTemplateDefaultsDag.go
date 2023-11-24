package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsDag struct {
	Tasks *[]*WorkflowTemplateSpecTemplateDefaultsDagTasks `field:"required" json:"tasks" yaml:"tasks"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	Target *string `field:"optional" json:"target" yaml:"target"`
}

