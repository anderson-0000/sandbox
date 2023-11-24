// cronworkflow _argoprojio
package cronworkflow _argoprojio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflow",
		reflect.TypeOf((*CronWorkflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CronWorkflow{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowProps",
		reflect.TypeOf((*CronWorkflowProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpec",
		reflect.TypeOf((*CronWorkflowSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArguments",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecArtifactRepositoryRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecArtifactRepositoryRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecDnsConfig",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecDnsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecDnsConfigOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecDnsConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecExecutor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooks",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArguments",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHooksTemplateRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecHostAliases",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecImagePullSecrets",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecMetrics",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecMetricsPrometheus",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecMetricsPrometheusCounter",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecMetricsPrometheusGauge",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecMetricsPrometheusHistogram",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecMetricsPrometheusLabels",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudget",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodDisruptionBudget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodDisruptionBudgetSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodDisruptionBudgetSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodGcLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodGcLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodGcLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodGcLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyBackoff",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyLimit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSecurityContextSysctls",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSynchronization",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSynchronizationMutex",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSynchronizationSemaphore",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaults",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaults)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsActiveDeadlineSeconds",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsActiveDeadlineSeconds)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsActiveDeadlineSeconds{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocation",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainer",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainers",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDag",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasks",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArguments",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksContinueOn",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksContinueOn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooks",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArguments",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksTemplateRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksTemplateRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequence",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequence)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceCount",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceCount{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsData",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPaths",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDataTransformation",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataTransformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsExecutor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsHostAliases",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpBodyFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpBodyFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeadersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeadersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeadersValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsHttpHeadersValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainers",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMemoize",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMemoize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMemoizeCache",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMemoizeCache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMemoizeCacheConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMemoizeCacheConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMetrics",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheus",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusCounter",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusGauge",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusHistogram",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusLabels",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifact",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoff",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyLimit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScript",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScript)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSysctls",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecars",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSuspend",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSuspend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronization",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronizationMutex",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronizationSemaphore",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsTolerations",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumes",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAwsElasticBlockStore",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAzureDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAzureFile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCephfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCephfsSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCinder",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCinderSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesConfigMapItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCsi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEmptyDir",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeral",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlexVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlocker",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGcePersistentDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGitRepo",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGlusterfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesHostPath",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesIscsi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesIscsiSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesNfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPersistentVolumeClaim",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPhotonPersistentDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPortworxVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjected",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesQuobyte",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesRbd",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesRbdSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesScaleIo",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesScaleIoSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesSecretItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesStorageos",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesStorageosSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesVsphereVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplates",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplates)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocation",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainer",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainers",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerSetVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesContainerVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDag",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasks",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArguments",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksContinueOn",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksContinueOn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooks",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArguments",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksTemplateRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksTemplateRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequence",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequence)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesData",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPaths",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDataTransformation",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesDataTransformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesExecutor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesHostAliases",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesHttpBodyFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesHttpBodyFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesHttpHeadersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesHttpHeadersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesHttpHeadersValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesHttpHeadersValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainers",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesInputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMemoize",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMemoize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMemoizeCache",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMemoizeCache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMemoizeCacheConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMemoizeCacheConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMetrics",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheus",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusCounter",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusGauge",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusHistogram",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusLabels",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifacts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsParameters",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsParametersValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesOutputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesOutputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifact",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArchive",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArchiveTar",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcPodMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactory",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactAzure",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactAzureAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGcs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGcsServiceAccountKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGitPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGitSshPrivateKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGitUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbCCacheSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbConfigConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbKeytabSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttp",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuth",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCert",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOss",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssAccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssLifecycleRule",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssSecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactRaw",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3AccessKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3CaSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3CreateBucketIfNotPresent",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3SecretKeySecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyAffinity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoff",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScript",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScript)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesScriptVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSecurityContextSysctls",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecars",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnv",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFromConfigMapRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFromSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromSecretKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecycle",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStart",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStop",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsPorts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContext",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextCapabilities",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextSeccompProfile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextWindowsOptions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbe",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeExec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeGrpc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGet",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeTcpSocket",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsVolumeDevices",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsVolumeMounts",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSidecarsVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSuspend",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSuspend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSynchronization",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSynchronizationMutex",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSynchronizationSemaphore",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesTolerations",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumes",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesAwsElasticBlockStore",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesAzureDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesAzureFile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesCephfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesCephfsSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesCinder",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesCinderSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesConfigMapItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesCsi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDir",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeral",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesFc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesFlexVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesFlocker",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesGcePersistentDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesGitRepo",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesGlusterfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesHostPath",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesIscsi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesIscsiSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesNfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesPersistentVolumeClaim",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesPhotonPersistentDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesPortworxVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjected",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesQuobyte",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesRbd",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesRbdSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesScaleIo",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesScaleIoSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesSecretItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesStorageos",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesStorageosSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesVsphereVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTemplatesVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTolerations",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTtlStrategy",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecTtlStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimGc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplates",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecDataSource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecDataSourceRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatus",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusConditions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumes",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesAwsElasticBlockStore",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesAzureDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesAzureFile",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesCephfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesCephfsSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesCinder",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesCinderSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesConfigMapItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesCsi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesDownwardApi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesDownwardApiItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEmptyDir",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeral",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesFc",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesFlexVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesFlocker",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesGcePersistentDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesGitRepo",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesGlusterfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesHostPath",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesIscsi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesIscsiSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesNfs",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesPersistentVolumeClaim",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesPhotonPersistentDisk",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesPortworxVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjected",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSources",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesQuobyte",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesRbd",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesRbdSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesScaleIo",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesScaleIoSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesSecret",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesSecretItems",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesStorageos",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesStorageosSecretRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumesVsphereVolume",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecWorkflowMetadata",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecWorkflowMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecWorkflowMetadataLabelsFrom",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecWorkflowMetadataLabelsFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecWorkflowTemplateRef",
		reflect.TypeOf((*CronWorkflowSpecWorkflowSpecWorkflowTemplateRef)(nil)).Elem(),
	)
}
