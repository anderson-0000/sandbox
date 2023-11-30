package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecPodDisruptionBudgetSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecPodDisruptionBudgetSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

