package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpec struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	DataSource *CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	DataSourceRef *CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef `field:"optional" json:"dataSourceRef" yaml:"dataSourceRef"`
	Resources *CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResources `field:"optional" json:"resources" yaml:"resources"`
	Selector *CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

