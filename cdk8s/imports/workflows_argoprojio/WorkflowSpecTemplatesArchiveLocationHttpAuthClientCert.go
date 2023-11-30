package workflows_argoprojio


type WorkflowSpecTemplatesArchiveLocationHttpAuthClientCert struct {
	ClientCertSecret *WorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientCertSecret `field:"optional" json:"clientCertSecret" yaml:"clientCertSecret"`
	ClientKeySecret *WorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientKeySecret `field:"optional" json:"clientKeySecret" yaml:"clientKeySecret"`
}

