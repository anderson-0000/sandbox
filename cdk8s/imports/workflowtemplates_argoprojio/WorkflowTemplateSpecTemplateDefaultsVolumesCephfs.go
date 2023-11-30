package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumesCephfs struct {
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	SecretRef *WorkflowTemplateSpecTemplateDefaultsVolumesCephfsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	User *string `field:"optional" json:"user" yaml:"user"`
}

