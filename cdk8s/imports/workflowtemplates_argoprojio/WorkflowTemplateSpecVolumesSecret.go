package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesSecret struct {
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	Items *[]*WorkflowTemplateSpecVolumesSecretItems `field:"optional" json:"items" yaml:"items"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

