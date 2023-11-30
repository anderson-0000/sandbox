package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesPhotonPersistentDisk struct {
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

