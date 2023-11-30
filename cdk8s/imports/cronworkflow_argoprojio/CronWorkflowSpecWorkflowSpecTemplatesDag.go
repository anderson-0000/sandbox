package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDag struct {
	Tasks *[]*CronWorkflowSpecWorkflowSpecTemplatesDagTasks `field:"required" json:"tasks" yaml:"tasks"`
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	Target *string `field:"optional" json:"target" yaml:"target"`
}

