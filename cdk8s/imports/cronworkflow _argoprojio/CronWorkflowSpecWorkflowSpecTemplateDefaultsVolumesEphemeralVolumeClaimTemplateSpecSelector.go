package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector struct {
	MatchExpressions *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

