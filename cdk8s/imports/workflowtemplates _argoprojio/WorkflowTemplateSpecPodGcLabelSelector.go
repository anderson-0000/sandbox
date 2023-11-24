package workflowtemplates _argoprojio


type WorkflowTemplateSpecPodGcLabelSelector struct {
	MatchExpressions *[]*WorkflowTemplateSpecPodGcLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

