package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

