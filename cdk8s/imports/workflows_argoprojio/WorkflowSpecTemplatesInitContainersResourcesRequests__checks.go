//go:build !no_runtime_type_checking

package workflows_argoprojio

import (
	"fmt"
)

func validateWorkflowSpecTemplatesInitContainersResourcesRequests_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowSpecTemplatesInitContainersResourcesRequests_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

