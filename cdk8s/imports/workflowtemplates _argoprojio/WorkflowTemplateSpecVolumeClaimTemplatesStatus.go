package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumeClaimTemplatesStatus struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	AllocatedResources *map[string]WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources `field:"optional" json:"allocatedResources" yaml:"allocatedResources"`
	Capacity *map[string]WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	Conditions *[]*WorkflowTemplateSpecVolumeClaimTemplatesStatusConditions `field:"optional" json:"conditions" yaml:"conditions"`
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	ResizeStatus *string `field:"optional" json:"resizeStatus" yaml:"resizeStatus"`
}

