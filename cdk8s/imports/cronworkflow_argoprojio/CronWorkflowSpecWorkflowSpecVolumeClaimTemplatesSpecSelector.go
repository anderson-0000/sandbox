package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

