package workflows _argoprojio


type WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources struct {
	Limits *map[string]WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

