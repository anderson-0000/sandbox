package workflows_argoprojio


type WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector struct {
	MatchExpressions *[]*WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

