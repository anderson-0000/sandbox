package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerResources struct {
	Limits *map[string]CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	Requests *map[string]CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

