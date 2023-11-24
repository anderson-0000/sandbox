//go:build !no_runtime_type_checking

package workflowtemplates _argoprojio

import (
	"fmt"
)

func validateWorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

