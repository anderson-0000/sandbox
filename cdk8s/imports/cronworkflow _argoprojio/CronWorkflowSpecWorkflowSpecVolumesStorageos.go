package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecVolumesStorageos struct {
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *CronWorkflowSpecWorkflowSpecVolumesStorageosSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
	VolumeNamespace *string `field:"optional" json:"volumeNamespace" yaml:"volumeNamespace"`
}

