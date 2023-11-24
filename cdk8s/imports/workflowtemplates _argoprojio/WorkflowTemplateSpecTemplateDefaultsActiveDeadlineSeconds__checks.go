//go:build !no_runtime_type_checking

package workflowtemplates _argoprojio

import (
	"fmt"
)

func validateWorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

