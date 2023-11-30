package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecPodDisruptionBudget struct {
	MaxUnavailable CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	MinAvailable CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable `field:"optional" json:"minAvailable" yaml:"minAvailable"`
	Selector *CronWorkflowSpecWorkflowSpecPodDisruptionBudgetSelector `field:"optional" json:"selector" yaml:"selector"`
}

