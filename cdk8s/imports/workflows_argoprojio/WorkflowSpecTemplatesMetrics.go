package workflows_argoprojio


type WorkflowSpecTemplatesMetrics struct {
	Prometheus *[]*WorkflowSpecTemplatesMetricsPrometheus `field:"required" json:"prometheus" yaml:"prometheus"`
}

