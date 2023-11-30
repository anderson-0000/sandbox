package workflowtemplates_argoprojio


type WorkflowTemplateSpecPodDisruptionBudget struct {
	MaxUnavailable WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	MinAvailable WorkflowTemplateSpecPodDisruptionBudgetMinAvailable `field:"optional" json:"minAvailable" yaml:"minAvailable"`
	Selector *WorkflowTemplateSpecPodDisruptionBudgetSelector `field:"optional" json:"selector" yaml:"selector"`
}

