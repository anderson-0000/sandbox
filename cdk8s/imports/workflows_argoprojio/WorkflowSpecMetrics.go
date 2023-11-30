package workflows_argoprojio


type WorkflowSpecMetrics struct {
	Prometheus *[]*WorkflowSpecMetricsPrometheus `field:"required" json:"prometheus" yaml:"prometheus"`
}

