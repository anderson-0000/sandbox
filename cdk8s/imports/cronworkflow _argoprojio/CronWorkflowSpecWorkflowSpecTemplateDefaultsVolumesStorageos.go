package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesStorageos struct {
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesStorageosSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
	VolumeNamespace *string `field:"optional" json:"volumeNamespace" yaml:"volumeNamespace"`
}

