package workflows_argoprojio


type WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources struct {
	Limits *map[string]WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

