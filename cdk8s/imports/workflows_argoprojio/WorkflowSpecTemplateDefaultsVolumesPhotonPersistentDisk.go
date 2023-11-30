package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesPhotonPersistentDisk struct {
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

