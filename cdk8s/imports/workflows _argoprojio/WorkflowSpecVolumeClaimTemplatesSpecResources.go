package workflows _argoprojio


type WorkflowSpecVolumeClaimTemplatesSpecResources struct {
	Limits *map[string]WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

