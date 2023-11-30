package workflows_argoprojio


type WorkflowSpecTemplatesContainerSetContainersResources struct {
	Limits *map[string]WorkflowSpecTemplatesContainerSetContainersResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplatesContainerSetContainersResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

