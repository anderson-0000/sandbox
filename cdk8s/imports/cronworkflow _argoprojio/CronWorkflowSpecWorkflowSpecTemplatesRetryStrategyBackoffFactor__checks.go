//go:build !no_runtime_type_checking

package cronworkflow _argoprojio

import (
	"fmt"
)

func validateCronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

