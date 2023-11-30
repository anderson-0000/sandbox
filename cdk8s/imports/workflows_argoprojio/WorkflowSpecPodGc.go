package workflows_argoprojio


type WorkflowSpecPodGc struct {
	DeleteDelayDuration *string `field:"optional" json:"deleteDelayDuration" yaml:"deleteDelayDuration"`
	LabelSelector *WorkflowSpecPodGcLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

