package workflows _argoprojio


type WorkflowSpecPodDisruptionBudget struct {
	MaxUnavailable WorkflowSpecPodDisruptionBudgetMaxUnavailable `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	MinAvailable WorkflowSpecPodDisruptionBudgetMinAvailable `field:"optional" json:"minAvailable" yaml:"minAvailable"`
	Selector *WorkflowSpecPodDisruptionBudgetSelector `field:"optional" json:"selector" yaml:"selector"`
}

