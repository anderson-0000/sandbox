package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerResources struct {
	Limits *map[string]WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowTemplateSpecTemplateDefaultsContainerResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

