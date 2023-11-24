package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumesCinder struct {
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *WorkflowTemplateSpecVolumesCinderSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

