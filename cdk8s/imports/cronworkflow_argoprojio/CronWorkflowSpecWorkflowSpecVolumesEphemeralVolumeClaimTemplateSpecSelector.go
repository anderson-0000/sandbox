package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

