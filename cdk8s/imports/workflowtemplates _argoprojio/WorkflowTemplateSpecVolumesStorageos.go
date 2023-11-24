package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumesStorageos struct {
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *WorkflowTemplateSpecVolumesStorageosSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
	VolumeNamespace *string `field:"optional" json:"volumeNamespace" yaml:"volumeNamespace"`
}

