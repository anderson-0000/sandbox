package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources struct {
	Limits *map[string]CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

