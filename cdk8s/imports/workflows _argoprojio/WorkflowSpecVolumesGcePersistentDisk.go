package workflows _argoprojio


type WorkflowSpecVolumesGcePersistentDisk struct {
	PdName *string `field:"required" json:"pdName" yaml:"pdName"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

