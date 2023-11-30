package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesCephfs struct {
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	SecretRef *WorkflowSpecTemplateDefaultsVolumesCephfsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	User *string `field:"optional" json:"user" yaml:"user"`
}

