package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesVolumesCinder struct {
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplatesVolumesCinderSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

