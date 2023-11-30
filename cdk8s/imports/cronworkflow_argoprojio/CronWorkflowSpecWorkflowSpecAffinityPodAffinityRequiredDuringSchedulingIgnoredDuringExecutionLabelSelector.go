package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

