package workflows_argoprojio


type WorkflowSpecTemplatesVolumesCephfs struct {
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	SecretRef *WorkflowSpecTemplatesVolumesCephfsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	User *string `field:"optional" json:"user" yaml:"user"`
}

