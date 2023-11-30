package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerResources struct {
	Limits *map[string]WorkflowTemplateSpecTemplatesContainerResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowTemplateSpecTemplatesContainerResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

