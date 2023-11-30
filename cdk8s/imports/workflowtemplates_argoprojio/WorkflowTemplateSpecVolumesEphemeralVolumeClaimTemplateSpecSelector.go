package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelector struct {
	MatchExpressions *[]*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

