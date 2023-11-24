package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecVolumesCinder struct {
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *CronWorkflowSpecWorkflowSpecVolumesCinderSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

