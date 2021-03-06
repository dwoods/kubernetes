// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package testing

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_testing_EmbeddedTest, InType: reflect.TypeOf(&EmbeddedTest{})},
		{Fn: DeepCopy_testing_EmbeddedTestExternal, InType: reflect.TypeOf(&EmbeddedTestExternal{})},
		{Fn: DeepCopy_testing_ExtensionA, InType: reflect.TypeOf(&ExtensionA{})},
		{Fn: DeepCopy_testing_ExtensionB, InType: reflect.TypeOf(&ExtensionB{})},
		{Fn: DeepCopy_testing_ExternalComplex, InType: reflect.TypeOf(&ExternalComplex{})},
		{Fn: DeepCopy_testing_ExternalExtensionType, InType: reflect.TypeOf(&ExternalExtensionType{})},
		{Fn: DeepCopy_testing_ExternalInternalSame, InType: reflect.TypeOf(&ExternalInternalSame{})},
		{Fn: DeepCopy_testing_ExternalOptionalExtensionType, InType: reflect.TypeOf(&ExternalOptionalExtensionType{})},
		{Fn: DeepCopy_testing_ExternalSimple, InType: reflect.TypeOf(&ExternalSimple{})},
		{Fn: DeepCopy_testing_ExternalTestType1, InType: reflect.TypeOf(&ExternalTestType1{})},
		{Fn: DeepCopy_testing_ExternalTestType2, InType: reflect.TypeOf(&ExternalTestType2{})},
		{Fn: DeepCopy_testing_InternalComplex, InType: reflect.TypeOf(&InternalComplex{})},
		{Fn: DeepCopy_testing_InternalExtensionType, InType: reflect.TypeOf(&InternalExtensionType{})},
		{Fn: DeepCopy_testing_InternalOptionalExtensionType, InType: reflect.TypeOf(&InternalOptionalExtensionType{})},
		{Fn: DeepCopy_testing_InternalSimple, InType: reflect.TypeOf(&InternalSimple{})},
		{Fn: DeepCopy_testing_ObjectTest, InType: reflect.TypeOf(&ObjectTest{})},
		{Fn: DeepCopy_testing_ObjectTestExternal, InType: reflect.TypeOf(&ObjectTestExternal{})},
		{Fn: DeepCopy_testing_TestType1, InType: reflect.TypeOf(&TestType1{})},
		{Fn: DeepCopy_testing_TestType2, InType: reflect.TypeOf(&TestType2{})},
		{Fn: DeepCopy_testing_UnknownType, InType: reflect.TypeOf(&UnknownType{})},
		{Fn: DeepCopy_testing_UnversionedType, InType: reflect.TypeOf(&UnversionedType{})},
	}
}

// DeepCopy_testing_EmbeddedTest is an autogenerated deepcopy function.
func DeepCopy_testing_EmbeddedTest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EmbeddedTest)
		out := out.(*EmbeddedTest)
		*out = *in
		// in.Object is kind 'Interface'
		if in.Object != nil {
			if newVal, err := c.DeepCopy(&in.Object); err != nil {
				return err
			} else {
				out.Object = *newVal.(*runtime.Object)
			}
		}
		// in.EmptyObject is kind 'Interface'
		if in.EmptyObject != nil {
			if newVal, err := c.DeepCopy(&in.EmptyObject); err != nil {
				return err
			} else {
				out.EmptyObject = *newVal.(*runtime.Object)
			}
		}
		return nil
	}
}

// DeepCopy_testing_EmbeddedTestExternal is an autogenerated deepcopy function.
func DeepCopy_testing_EmbeddedTestExternal(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EmbeddedTestExternal)
		out := out.(*EmbeddedTestExternal)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Object); err != nil {
			return err
		} else {
			out.Object = *newVal.(*runtime.RawExtension)
		}
		if newVal, err := c.DeepCopy(&in.EmptyObject); err != nil {
			return err
		} else {
			out.EmptyObject = *newVal.(*runtime.RawExtension)
		}
		return nil
	}
}

// DeepCopy_testing_ExtensionA is an autogenerated deepcopy function.
func DeepCopy_testing_ExtensionA(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExtensionA)
		out := out.(*ExtensionA)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_ExtensionB is an autogenerated deepcopy function.
func DeepCopy_testing_ExtensionB(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExtensionB)
		out := out.(*ExtensionB)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_ExternalComplex is an autogenerated deepcopy function.
func DeepCopy_testing_ExternalComplex(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalComplex)
		out := out.(*ExternalComplex)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_ExternalExtensionType is an autogenerated deepcopy function.
func DeepCopy_testing_ExternalExtensionType(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalExtensionType)
		out := out.(*ExternalExtensionType)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Extension); err != nil {
			return err
		} else {
			out.Extension = *newVal.(*runtime.RawExtension)
		}
		return nil
	}
}

// DeepCopy_testing_ExternalInternalSame is an autogenerated deepcopy function.
func DeepCopy_testing_ExternalInternalSame(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalInternalSame)
		out := out.(*ExternalInternalSame)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_ExternalOptionalExtensionType is an autogenerated deepcopy function.
func DeepCopy_testing_ExternalOptionalExtensionType(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalOptionalExtensionType)
		out := out.(*ExternalOptionalExtensionType)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Extension); err != nil {
			return err
		} else {
			out.Extension = *newVal.(*runtime.RawExtension)
		}
		return nil
	}
}

// DeepCopy_testing_ExternalSimple is an autogenerated deepcopy function.
func DeepCopy_testing_ExternalSimple(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalSimple)
		out := out.(*ExternalSimple)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_ExternalTestType1 is an autogenerated deepcopy function.
func DeepCopy_testing_ExternalTestType1(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalTestType1)
		out := out.(*ExternalTestType1)
		*out = *in
		if in.M != nil {
			in, out := &in.M, &out.M
			*out = make(map[string]int)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.N != nil {
			in, out := &in.N, &out.N
			*out = make(map[string]ExternalTestType2)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.O != nil {
			in, out := &in.O, &out.O
			*out = new(ExternalTestType2)
			**out = **in
		}
		if in.P != nil {
			in, out := &in.P, &out.P
			*out = make([]ExternalTestType2, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_testing_ExternalTestType2 is an autogenerated deepcopy function.
func DeepCopy_testing_ExternalTestType2(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalTestType2)
		out := out.(*ExternalTestType2)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_InternalComplex is an autogenerated deepcopy function.
func DeepCopy_testing_InternalComplex(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InternalComplex)
		out := out.(*InternalComplex)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_InternalExtensionType is an autogenerated deepcopy function.
func DeepCopy_testing_InternalExtensionType(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InternalExtensionType)
		out := out.(*InternalExtensionType)
		*out = *in
		// in.Extension is kind 'Interface'
		if in.Extension != nil {
			if newVal, err := c.DeepCopy(&in.Extension); err != nil {
				return err
			} else {
				out.Extension = *newVal.(*runtime.Object)
			}
		}
		return nil
	}
}

// DeepCopy_testing_InternalOptionalExtensionType is an autogenerated deepcopy function.
func DeepCopy_testing_InternalOptionalExtensionType(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InternalOptionalExtensionType)
		out := out.(*InternalOptionalExtensionType)
		*out = *in
		// in.Extension is kind 'Interface'
		if in.Extension != nil {
			if newVal, err := c.DeepCopy(&in.Extension); err != nil {
				return err
			} else {
				out.Extension = *newVal.(*runtime.Object)
			}
		}
		return nil
	}
}

// DeepCopy_testing_InternalSimple is an autogenerated deepcopy function.
func DeepCopy_testing_InternalSimple(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InternalSimple)
		out := out.(*InternalSimple)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_ObjectTest is an autogenerated deepcopy function.
func DeepCopy_testing_ObjectTest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ObjectTest)
		out := out.(*ObjectTest)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]runtime.Object, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*runtime.Object)
				}
			}
		}
		return nil
	}
}

// DeepCopy_testing_ObjectTestExternal is an autogenerated deepcopy function.
func DeepCopy_testing_ObjectTestExternal(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ObjectTestExternal)
		out := out.(*ObjectTestExternal)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]runtime.RawExtension, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*runtime.RawExtension)
				}
			}
		}
		return nil
	}
}

// DeepCopy_testing_TestType1 is an autogenerated deepcopy function.
func DeepCopy_testing_TestType1(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TestType1)
		out := out.(*TestType1)
		*out = *in
		if in.M != nil {
			in, out := &in.M, &out.M
			*out = make(map[string]int)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.N != nil {
			in, out := &in.N, &out.N
			*out = make(map[string]TestType2)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.O != nil {
			in, out := &in.O, &out.O
			*out = new(TestType2)
			**out = **in
		}
		if in.P != nil {
			in, out := &in.P, &out.P
			*out = make([]TestType2, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_testing_TestType2 is an autogenerated deepcopy function.
func DeepCopy_testing_TestType2(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TestType2)
		out := out.(*TestType2)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_UnknownType is an autogenerated deepcopy function.
func DeepCopy_testing_UnknownType(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UnknownType)
		out := out.(*UnknownType)
		*out = *in
		return nil
	}
}

// DeepCopy_testing_UnversionedType is an autogenerated deepcopy function.
func DeepCopy_testing_UnversionedType(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UnversionedType)
		out := out.(*UnversionedType)
		*out = *in
		return nil
	}
}
