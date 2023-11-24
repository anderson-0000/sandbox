//go:build !no_runtime_type_checking

package workflows _argoprojio

import (
	"fmt"
)

func validateWorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

