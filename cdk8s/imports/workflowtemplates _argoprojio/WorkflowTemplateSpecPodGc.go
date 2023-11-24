package workflowtemplates _argoprojio


type WorkflowTemplateSpecPodGc struct {
	DeleteDelayDuration *string `field:"optional" json:"deleteDelayDuration" yaml:"deleteDelayDuration"`
	LabelSelector *WorkflowTemplateSpecPodGcLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

