package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

