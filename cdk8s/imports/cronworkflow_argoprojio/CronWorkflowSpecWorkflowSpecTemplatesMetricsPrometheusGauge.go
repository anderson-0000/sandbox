package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusGauge struct {
	Realtime *bool `field:"required" json:"realtime" yaml:"realtime"`
	Value *string `field:"required" json:"value" yaml:"value"`
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
}

