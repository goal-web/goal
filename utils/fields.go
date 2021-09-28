package utils

import "github.com/qbhy/goal/contracts"

func MergeFields(fields contracts.Fields, finalFields contracts.Fields) {
	for key, value := range finalFields {
		fields[key] = value
	}
}
