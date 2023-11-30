package workflows_argoprojio


type WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector struct {
	MatchExpressions *[]*WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

