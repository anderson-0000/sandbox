package workflows_argoprojio


type WorkflowSpecSecurityContextSeccompProfile struct {
	Type *string `field:"required" json:"type" yaml:"type"`
	LocalhostProfile *string `field:"optional" json:"localhostProfile" yaml:"localhostProfile"`
}

