/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/dave/dst"
)

// Function represents something that is an (unnamed) Go function
type Function interface {
	// The unique name of this function
	// (You can't have two functions with the same name on the same object or resource)
	Name() string

	RequiredPackageReferences() *PackageReferenceSet

	// References returns the set of types to which this function refers.
	// Should *not* include the receiver of this function
	References() TypeNameSet

	// AsFunc renders the current instance as a Go abstract syntax tree
	AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl

	// Equals determines if this Function is equal to another one
	Equals(f Function) bool
}

var _ Function = &objectFunction{}

type objectFunctionHandler func(f *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl

// objectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
type objectFunction struct {
	name             string
	o                *ObjectType
	idFactory        IdentifierFactory
	asFunc           objectFunctionHandler
	requiredPackages *PackageReferenceSet
}

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (k *objectFunction) Name() string {
	return k.name
}

// RequiredPackageReferences returns the set of required packages for this function
func (k *objectFunction) RequiredPackageReferences() *PackageReferenceSet {
	return k.requiredPackages
}

// References returns the TypeName's referenced by this function
func (k *objectFunction) References() TypeNameSet {
	return k.o.References()
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (k *objectFunction) AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl {
	return k.asFunc(k, codeGenerationContext, receiver, k.name)
}

// Equals checks if this function is equal to the passed in function
func (k *objectFunction) Equals(f Function) bool {
	typedF, ok := f.(*objectFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	return k.o.Equals(typedF.o) && k.name == typedF.name
}

type resourceFunctionHandler func(f *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl

// resourceFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the resource they are operating on
type resourceFunction struct {
	name             string
	resource         *ResourceType
	idFactory        IdentifierFactory
	asFunc           resourceFunctionHandler
	requiredPackages *PackageReferenceSet
}

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (r *resourceFunction) Name() string {
	return r.name
}

// RequiredPackageReferences returns the set of required packages for this function
func (r *resourceFunction) RequiredPackageReferences() *PackageReferenceSet {
	return r.requiredPackages
}

// References returns the TypeName's referenced by this function
func (r *resourceFunction) References() TypeNameSet {
	return r.resource.References()
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (r *resourceFunction) AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl {
	return r.asFunc(r, codeGenerationContext, receiver, r.name)
}

// Equals determines if this Function is equal to another one
func (r *resourceFunction) Equals(f Function) bool {
	typedF, ok := f.(*resourceFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	return r.resource.Equals(typedF.resource) && r.name == typedF.name
}
