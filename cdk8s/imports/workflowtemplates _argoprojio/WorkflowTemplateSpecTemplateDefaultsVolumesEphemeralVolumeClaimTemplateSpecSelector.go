package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector struct {
	MatchExpressions *[]*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

