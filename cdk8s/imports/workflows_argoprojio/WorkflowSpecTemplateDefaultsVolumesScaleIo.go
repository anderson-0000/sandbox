package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesScaleIo struct {
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	SecretRef *WorkflowSpecTemplateDefaultsVolumesScaleIoSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
	System *string `field:"required" json:"system" yaml:"system"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ProtectionDomain *string `field:"optional" json:"protectionDomain" yaml:"protectionDomain"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SslEnabled *bool `field:"optional" json:"sslEnabled" yaml:"sslEnabled"`
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	StoragePool *string `field:"optional" json:"storagePool" yaml:"storagePool"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

