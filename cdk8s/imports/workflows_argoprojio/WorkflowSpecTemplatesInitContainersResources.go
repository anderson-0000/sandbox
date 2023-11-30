package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersResources struct {
	Limits *map[string]WorkflowSpecTemplatesInitContainersResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplatesInitContainersResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

