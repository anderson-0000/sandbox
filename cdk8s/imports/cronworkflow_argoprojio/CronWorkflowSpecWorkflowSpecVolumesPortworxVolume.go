package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesPortworxVolume struct {
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

