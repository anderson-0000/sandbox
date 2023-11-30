package workflowtemplates_argoprojio


type WorkflowTemplateSpecPodDisruptionBudgetSelector struct {
	MatchExpressions *[]*WorkflowTemplateSpecPodDisruptionBudgetSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

