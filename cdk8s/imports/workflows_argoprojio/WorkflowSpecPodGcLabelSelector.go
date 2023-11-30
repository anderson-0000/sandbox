package workflows_argoprojio


type WorkflowSpecPodGcLabelSelector struct {
	MatchExpressions *[]*WorkflowSpecPodGcLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

