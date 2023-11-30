package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResources struct {
	Limits *map[string]WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

