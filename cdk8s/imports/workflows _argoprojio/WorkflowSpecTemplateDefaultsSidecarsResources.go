package workflows _argoprojio


type WorkflowSpecTemplateDefaultsSidecarsResources struct {
	Limits *map[string]WorkflowSpecTemplateDefaultsSidecarsResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplateDefaultsSidecarsResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

