package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsPorts struct {
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	HostIp *string `field:"optional" json:"hostIp" yaml:"hostIp"`
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

