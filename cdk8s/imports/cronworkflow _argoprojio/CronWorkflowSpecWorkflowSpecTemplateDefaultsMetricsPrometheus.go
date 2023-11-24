package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

