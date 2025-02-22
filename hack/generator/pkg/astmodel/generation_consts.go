/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

var (
	CodeGenerationComments []string = []string{
		// Note that this format is actually an official specification in go:  https://github.com/golang/go/issues/13560
		"// Code generated by k8s-infra. DO NOT EDIT.", // TODO: Update this when the generated is moved/renamed
	}
)

const (
	// CodeGeneratedFileSuffix is used to identify generated files (note there is no file extension here)
	CodeGeneratedFileSuffix = "_gen"

	// ARMReferenceTag is the tag ID used for specifying references to other ARM resources on properties.
	ARMReferenceTag = "armReference"
)
