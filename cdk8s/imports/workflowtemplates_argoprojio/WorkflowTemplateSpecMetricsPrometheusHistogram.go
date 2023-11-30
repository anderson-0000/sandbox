package workflowtemplates_argoprojio


type WorkflowTemplateSpecMetricsPrometheusHistogram struct {
	Buckets *[]*float64 `field:"required" json:"buckets" yaml:"buckets"`
	Value *string `field:"required" json:"value" yaml:"value"`
}

