package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesCsi struct {
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	NodePublishSecretRef *CronWorkflowSpecWorkflowSpecVolumesCsiNodePublishSecretRef `field:"optional" json:"nodePublishSecretRef" yaml:"nodePublishSecretRef"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

