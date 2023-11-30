package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

