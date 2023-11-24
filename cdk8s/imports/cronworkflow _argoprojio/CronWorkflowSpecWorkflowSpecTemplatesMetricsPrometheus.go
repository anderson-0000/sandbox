package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

