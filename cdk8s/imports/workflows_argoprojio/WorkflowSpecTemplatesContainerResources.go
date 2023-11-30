package workflows_argoprojio


type WorkflowSpecTemplatesContainerResources struct {
	Limits *map[string]WorkflowSpecTemplatesContainerResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplatesContainerResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

