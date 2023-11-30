package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerResources struct {
	Limits *map[string]WorkflowSpecTemplateDefaultsContainerResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplateDefaultsContainerResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

