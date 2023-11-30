package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptResources struct {
	Limits *map[string]WorkflowSpecTemplateDefaultsScriptResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplateDefaultsScriptResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

