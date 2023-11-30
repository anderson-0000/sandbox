package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector struct {
	MatchExpressions *[]*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

