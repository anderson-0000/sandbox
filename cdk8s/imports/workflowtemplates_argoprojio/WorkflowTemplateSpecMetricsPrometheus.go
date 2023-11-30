package workflowtemplates_argoprojio


type WorkflowTemplateSpecMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *WorkflowTemplateSpecMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *WorkflowTemplateSpecMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *WorkflowTemplateSpecMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*WorkflowTemplateSpecMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

