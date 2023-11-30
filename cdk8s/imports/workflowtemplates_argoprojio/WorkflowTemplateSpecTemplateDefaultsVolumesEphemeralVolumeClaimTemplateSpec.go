package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpec struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	DataSource *WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	DataSourceRef *WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef `field:"optional" json:"dataSourceRef" yaml:"dataSourceRef"`
	Resources *WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources `field:"optional" json:"resources" yaml:"resources"`
	Selector *WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

