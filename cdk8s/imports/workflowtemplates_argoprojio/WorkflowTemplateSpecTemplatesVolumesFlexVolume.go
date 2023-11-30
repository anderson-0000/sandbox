package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesVolumesFlexVolume struct {
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *WorkflowTemplateSpecTemplatesVolumesFlexVolumeSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

