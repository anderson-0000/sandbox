package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsDag struct {
	Tasks *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasks `field:"required" json:"tasks" yaml:"tasks"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	Target *string `field:"optional" json:"target" yaml:"target"`
}

