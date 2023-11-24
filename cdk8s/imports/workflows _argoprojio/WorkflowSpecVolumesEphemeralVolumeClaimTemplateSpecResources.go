package workflows _argoprojio


type WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResources struct {
	Limits *map[string]WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

