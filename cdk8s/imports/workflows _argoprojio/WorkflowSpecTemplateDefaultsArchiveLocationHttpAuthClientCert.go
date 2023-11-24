package workflows _argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCert struct {
	ClientCertSecret *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientCertSecret `field:"optional" json:"clientCertSecret" yaml:"clientCertSecret"`
	ClientKeySecret *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientKeySecret `field:"optional" json:"clientKeySecret" yaml:"clientKeySecret"`
}

