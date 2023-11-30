package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersResources struct {
	Limits *map[string]WorkflowSpecTemplateDefaultsInitContainersResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplateDefaultsInitContainersResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

