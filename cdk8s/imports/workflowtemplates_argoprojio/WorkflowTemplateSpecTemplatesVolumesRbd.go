package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesVolumesRbd struct {
	Image *string `field:"required" json:"image" yaml:"image"`
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	Keyring *string `field:"optional" json:"keyring" yaml:"keyring"`
	Pool *string `field:"optional" json:"pool" yaml:"pool"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *WorkflowTemplateSpecTemplatesVolumesRbdSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	User *string `field:"optional" json:"user" yaml:"user"`
}

