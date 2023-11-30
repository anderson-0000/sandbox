package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesInitContainersResources struct {
	Limits *map[string]WorkflowTemplateSpecTemplatesInitContainersResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowTemplateSpecTemplatesInitContainersResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

