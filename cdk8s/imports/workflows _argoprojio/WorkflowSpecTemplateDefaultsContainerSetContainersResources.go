package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersResources struct {
	Limits *map[string]WorkflowSpecTemplateDefaultsContainerSetContainersResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

