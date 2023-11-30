package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesStorageos struct {
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *WorkflowSpecTemplateDefaultsVolumesStorageosSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
	VolumeNamespace *string `field:"optional" json:"volumeNamespace" yaml:"volumeNamespace"`
}

