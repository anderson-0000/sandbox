package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumeClaimTemplatesSpecResources struct {
	Limits *map[string]WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

