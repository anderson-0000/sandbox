package workflows _argoprojio


type WorkflowSpecTemplatesSidecarsResources struct {
	Limits *map[string]WorkflowSpecTemplatesSidecarsResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplatesSidecarsResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

