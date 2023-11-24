package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecPodGc struct {
	DeleteDelayDuration *string `field:"optional" json:"deleteDelayDuration" yaml:"deleteDelayDuration"`
	LabelSelector *CronWorkflowSpecWorkflowSpecPodGcLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

