package workflows_argoprojio


type WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector struct {
	MatchExpressions *[]*WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

