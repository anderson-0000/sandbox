package workflows_argoprojio


type WorkflowSpecTemplatesMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *WorkflowSpecTemplatesMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *WorkflowSpecTemplatesMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *WorkflowSpecTemplatesMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*WorkflowSpecTemplatesMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

