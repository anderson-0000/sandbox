package workflows_argoprojio


type WorkflowSpecVolumeClaimTemplatesStatus struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	AllocatedResources *map[string]WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources `field:"optional" json:"allocatedResources" yaml:"allocatedResources"`
	Capacity *map[string]WorkflowSpecVolumeClaimTemplatesStatusCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	Conditions *[]*WorkflowSpecVolumeClaimTemplatesStatusConditions `field:"optional" json:"conditions" yaml:"conditions"`
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	ResizeStatus *string `field:"optional" json:"resizeStatus" yaml:"resizeStatus"`
}

