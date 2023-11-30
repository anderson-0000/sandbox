package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersVolumeMounts struct {
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	Name *string `field:"required" json:"name" yaml:"name"`
	MountPropagation *string `field:"optional" json:"mountPropagation" yaml:"mountPropagation"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
	SubPathExpr *string `field:"optional" json:"subPathExpr" yaml:"subPathExpr"`
}

