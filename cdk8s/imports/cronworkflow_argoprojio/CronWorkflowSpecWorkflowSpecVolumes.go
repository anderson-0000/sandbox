package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *CronWorkflowSpecWorkflowSpecVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *CronWorkflowSpecWorkflowSpecVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *CronWorkflowSpecWorkflowSpecVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *CronWorkflowSpecWorkflowSpecVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *CronWorkflowSpecWorkflowSpecVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *CronWorkflowSpecWorkflowSpecVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *CronWorkflowSpecWorkflowSpecVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *CronWorkflowSpecWorkflowSpecVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *CronWorkflowSpecWorkflowSpecVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *CronWorkflowSpecWorkflowSpecVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *CronWorkflowSpecWorkflowSpecVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *CronWorkflowSpecWorkflowSpecVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *CronWorkflowSpecWorkflowSpecVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *CronWorkflowSpecWorkflowSpecVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *CronWorkflowSpecWorkflowSpecVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *CronWorkflowSpecWorkflowSpecVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *CronWorkflowSpecWorkflowSpecVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *CronWorkflowSpecWorkflowSpecVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *CronWorkflowSpecWorkflowSpecVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *CronWorkflowSpecWorkflowSpecVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *CronWorkflowSpecWorkflowSpecVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *CronWorkflowSpecWorkflowSpecVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *CronWorkflowSpecWorkflowSpecVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *CronWorkflowSpecWorkflowSpecVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *CronWorkflowSpecWorkflowSpecVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *CronWorkflowSpecWorkflowSpecVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *CronWorkflowSpecWorkflowSpecVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *CronWorkflowSpecWorkflowSpecVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *CronWorkflowSpecWorkflowSpecVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

