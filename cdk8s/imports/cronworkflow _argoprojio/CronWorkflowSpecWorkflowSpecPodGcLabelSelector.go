package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecPodGcLabelSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecPodGcLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

