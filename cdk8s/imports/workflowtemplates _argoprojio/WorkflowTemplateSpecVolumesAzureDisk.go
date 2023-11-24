package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumesAzureDisk struct {
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	DiskUri *string `field:"required" json:"diskUri" yaml:"diskUri"`
	CachingMode *string `field:"optional" json:"cachingMode" yaml:"cachingMode"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

