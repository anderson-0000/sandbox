package workflows_argoprojio


type WorkflowSpecVolumesCinder struct {
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *WorkflowSpecVolumesCinderSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

