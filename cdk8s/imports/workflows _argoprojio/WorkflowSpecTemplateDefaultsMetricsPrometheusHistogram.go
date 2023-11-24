package workflows _argoprojio


type WorkflowSpecTemplateDefaultsMetricsPrometheusHistogram struct {
	Buckets *[]*float64 `field:"required" json:"buckets" yaml:"buckets"`
	Value *string `field:"required" json:"value" yaml:"value"`
}

