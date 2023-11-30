package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumeClaimTemplatesSpecSelector struct {
	MatchExpressions *[]*WorkflowTemplateSpecVolumeClaimTemplatesSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

