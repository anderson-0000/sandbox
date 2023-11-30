package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector struct {
	MatchExpressions *[]*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

