package workflows _argoprojio


type WorkflowSpecVolumeClaimTemplatesSpecSelector struct {
	MatchExpressions *[]*WorkflowSpecVolumeClaimTemplatesSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

