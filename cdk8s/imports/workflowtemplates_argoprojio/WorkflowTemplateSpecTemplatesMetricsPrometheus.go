package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *WorkflowTemplateSpecTemplatesMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *WorkflowTemplateSpecTemplatesMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *WorkflowTemplateSpecTemplatesMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*WorkflowTemplateSpecTemplatesMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

