package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesMetricsPrometheusGauge struct {
	Realtime *bool `field:"required" json:"realtime" yaml:"realtime"`
	Value *string `field:"required" json:"value" yaml:"value"`
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
}

