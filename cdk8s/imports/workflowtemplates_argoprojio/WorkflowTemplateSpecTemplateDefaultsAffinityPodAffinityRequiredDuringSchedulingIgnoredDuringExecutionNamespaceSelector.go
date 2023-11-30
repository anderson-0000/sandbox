package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector struct {
	MatchExpressions *[]*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

