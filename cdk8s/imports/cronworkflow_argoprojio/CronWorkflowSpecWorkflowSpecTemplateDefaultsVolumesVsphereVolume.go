package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesVsphereVolume struct {
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	StoragePolicyId *string `field:"optional" json:"storagePolicyId" yaml:"storagePolicyId"`
	StoragePolicyName *string `field:"optional" json:"storagePolicyName" yaml:"storagePolicyName"`
}

