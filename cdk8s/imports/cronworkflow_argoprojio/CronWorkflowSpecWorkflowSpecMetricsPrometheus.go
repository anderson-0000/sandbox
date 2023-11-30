package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecMetricsPrometheus struct {
	Help *string `field:"required" json:"help" yaml:"help"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Counter *CronWorkflowSpecWorkflowSpecMetricsPrometheusCounter `field:"optional" json:"counter" yaml:"counter"`
	Gauge *CronWorkflowSpecWorkflowSpecMetricsPrometheusGauge `field:"optional" json:"gauge" yaml:"gauge"`
	Histogram *CronWorkflowSpecWorkflowSpecMetricsPrometheusHistogram `field:"optional" json:"histogram" yaml:"histogram"`
	Labels *[]*CronWorkflowSpecWorkflowSpecMetricsPrometheusLabels `field:"optional" json:"labels" yaml:"labels"`
	When *string `field:"optional" json:"when" yaml:"when"`
}

