package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatus struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	AllocatedResources *map[string]CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources `field:"optional" json:"allocatedResources" yaml:"allocatedResources"`
	Capacity *map[string]CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	Conditions *[]*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusConditions `field:"optional" json:"conditions" yaml:"conditions"`
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	ResizeStatus *string `field:"optional" json:"resizeStatus" yaml:"resizeStatus"`
}

