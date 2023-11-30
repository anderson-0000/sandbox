package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesIscsi struct {
	Iqn *string `field:"required" json:"iqn" yaml:"iqn"`
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	TargetPortal *string `field:"required" json:"targetPortal" yaml:"targetPortal"`
	ChapAuthDiscovery *bool `field:"optional" json:"chapAuthDiscovery" yaml:"chapAuthDiscovery"`
	ChapAuthSession *bool `field:"optional" json:"chapAuthSession" yaml:"chapAuthSession"`
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	InitiatorName *string `field:"optional" json:"initiatorName" yaml:"initiatorName"`
	IscsiInterface *string `field:"optional" json:"iscsiInterface" yaml:"iscsiInterface"`
	Portals *[]*string `field:"optional" json:"portals" yaml:"portals"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	SecretRef *WorkflowTemplateSpecVolumesIscsiSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

