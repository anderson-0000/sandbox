package workflows_argoprojio


type WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector struct {
	MatchExpressions *[]*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

