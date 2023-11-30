package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesVolumesGitRepo struct {
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

