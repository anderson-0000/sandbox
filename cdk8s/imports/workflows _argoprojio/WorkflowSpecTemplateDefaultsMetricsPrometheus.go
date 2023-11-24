package workflows _argoprojio


type WorkflowSpecTemplateDefaultsMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *WorkflowSpecTemplateDefaultsMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *WorkflowSpecTemplateDefaultsMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *WorkflowSpecTemplateDefaultsMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*WorkflowSpecTemplateDefaultsMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

