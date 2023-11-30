package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesAwsElasticBlockStore struct {
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

