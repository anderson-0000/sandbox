package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumesFc struct {
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	TargetWwNs *[]*string `field:"optional" json:"targetWwNs" yaml:"targetWwNs"`
	Wwids *[]*string `field:"optional" json:"wwids" yaml:"wwids"`
}

