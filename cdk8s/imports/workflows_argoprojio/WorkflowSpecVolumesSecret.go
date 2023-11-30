package workflows_argoprojio


type WorkflowSpecVolumesSecret struct {
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	Items *[]*WorkflowSpecVolumesSecretItems `field:"optional" json:"items" yaml:"items"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

