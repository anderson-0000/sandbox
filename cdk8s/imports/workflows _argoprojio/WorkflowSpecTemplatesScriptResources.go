package workflows _argoprojio


type WorkflowSpecTemplatesScriptResources struct {
	Limits *map[string]WorkflowSpecTemplatesScriptResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]WorkflowSpecTemplatesScriptResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

