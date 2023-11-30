package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsScriptResources struct {
	Limits *map[string]WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

