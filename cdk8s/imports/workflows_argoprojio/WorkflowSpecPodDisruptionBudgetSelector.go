package workflows_argoprojio


type WorkflowSpecPodDisruptionBudgetSelector struct {
	MatchExpressions *[]*WorkflowSpecPodDisruptionBudgetSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

