// +build !ignore_autogenerated

/*
Copyright 2020 Red Hat, Inc.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	operatorsv1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	operators "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AppLink)(nil), (*operators.AppLink)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AppLink_To_operators_AppLink(a.(*AppLink), b.(*operators.AppLink), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.AppLink)(nil), (*AppLink)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_AppLink_To_v1alpha1_AppLink(a.(*operators.AppLink), b.(*AppLink), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CSVDescription)(nil), (*operators.CSVDescription)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CSVDescription_To_operators_CSVDescription(a.(*CSVDescription), b.(*operators.CSVDescription), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.CSVDescription)(nil), (*CSVDescription)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_CSVDescription_To_v1alpha1_CSVDescription(a.(*operators.CSVDescription), b.(*CSVDescription), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Icon)(nil), (*operators.Icon)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Icon_To_operators_Icon(a.(*Icon), b.(*operators.Icon), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.Icon)(nil), (*Icon)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_Icon_To_v1alpha1_Icon(a.(*operators.Icon), b.(*Icon), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageChannel)(nil), (*operators.PackageChannel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageChannel_To_operators_PackageChannel(a.(*PackageChannel), b.(*operators.PackageChannel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.PackageChannel)(nil), (*PackageChannel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_PackageChannel_To_v1alpha1_PackageChannel(a.(*operators.PackageChannel), b.(*PackageChannel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageManifest)(nil), (*operators.PackageManifest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageManifest_To_operators_PackageManifest(a.(*PackageManifest), b.(*operators.PackageManifest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.PackageManifest)(nil), (*PackageManifest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_PackageManifest_To_v1alpha1_PackageManifest(a.(*operators.PackageManifest), b.(*PackageManifest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageManifestList)(nil), (*operators.PackageManifestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageManifestList_To_operators_PackageManifestList(a.(*PackageManifestList), b.(*operators.PackageManifestList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.PackageManifestList)(nil), (*PackageManifestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_PackageManifestList_To_v1alpha1_PackageManifestList(a.(*operators.PackageManifestList), b.(*PackageManifestList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageManifestSpec)(nil), (*operators.PackageManifestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageManifestSpec_To_operators_PackageManifestSpec(a.(*PackageManifestSpec), b.(*operators.PackageManifestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.PackageManifestSpec)(nil), (*PackageManifestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_PackageManifestSpec_To_v1alpha1_PackageManifestSpec(a.(*operators.PackageManifestSpec), b.(*PackageManifestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PackageManifestStatus)(nil), (*operators.PackageManifestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PackageManifestStatus_To_operators_PackageManifestStatus(a.(*PackageManifestStatus), b.(*operators.PackageManifestStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*operators.PackageManifestStatus)(nil), (*PackageManifestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_operators_PackageManifestStatus_To_v1alpha1_PackageManifestStatus(a.(*operators.PackageManifestStatus), b.(*PackageManifestStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AppLink_To_operators_AppLink(in *AppLink, out *operators.AppLink, s conversion.Scope) error {
	out.Name = in.Name
	out.URL = in.URL
	return nil
}

// Convert_v1alpha1_AppLink_To_operators_AppLink is an autogenerated conversion function.
func Convert_v1alpha1_AppLink_To_operators_AppLink(in *AppLink, out *operators.AppLink, s conversion.Scope) error {
	return autoConvert_v1alpha1_AppLink_To_operators_AppLink(in, out, s)
}

func autoConvert_operators_AppLink_To_v1alpha1_AppLink(in *operators.AppLink, out *AppLink, s conversion.Scope) error {
	out.Name = in.Name
	out.URL = in.URL
	return nil
}

// Convert_operators_AppLink_To_v1alpha1_AppLink is an autogenerated conversion function.
func Convert_operators_AppLink_To_v1alpha1_AppLink(in *operators.AppLink, out *AppLink, s conversion.Scope) error {
	return autoConvert_operators_AppLink_To_v1alpha1_AppLink(in, out, s)
}

func autoConvert_v1alpha1_CSVDescription_To_operators_CSVDescription(in *CSVDescription, out *operators.CSVDescription, s conversion.Scope) error {
	out.DisplayName = in.DisplayName
	out.Icon = *(*[]operators.Icon)(unsafe.Pointer(&in.Icon))
	out.Version = in.Version
	if err := Convert_v1alpha1_AppLink_To_operators_AppLink(&in.Provider, &out.Provider, s); err != nil {
		return err
	}
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.LongDescription = in.LongDescription
	out.InstallModes = *(*[]operatorsv1alpha1.InstallMode)(unsafe.Pointer(&in.InstallModes))
	out.CustomResourceDefinitions = in.CustomResourceDefinitions
	out.APIServiceDefinitions = in.APIServiceDefinitions
	return nil
}

// Convert_v1alpha1_CSVDescription_To_operators_CSVDescription is an autogenerated conversion function.
func Convert_v1alpha1_CSVDescription_To_operators_CSVDescription(in *CSVDescription, out *operators.CSVDescription, s conversion.Scope) error {
	return autoConvert_v1alpha1_CSVDescription_To_operators_CSVDescription(in, out, s)
}

func autoConvert_operators_CSVDescription_To_v1alpha1_CSVDescription(in *operators.CSVDescription, out *CSVDescription, s conversion.Scope) error {
	out.DisplayName = in.DisplayName
	out.Icon = *(*[]Icon)(unsafe.Pointer(&in.Icon))
	out.Version = in.Version
	if err := Convert_operators_AppLink_To_v1alpha1_AppLink(&in.Provider, &out.Provider, s); err != nil {
		return err
	}
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.LongDescription = in.LongDescription
	out.InstallModes = *(*[]operatorsv1alpha1.InstallMode)(unsafe.Pointer(&in.InstallModes))
	out.CustomResourceDefinitions = in.CustomResourceDefinitions
	out.APIServiceDefinitions = in.APIServiceDefinitions
	return nil
}

// Convert_operators_CSVDescription_To_v1alpha1_CSVDescription is an autogenerated conversion function.
func Convert_operators_CSVDescription_To_v1alpha1_CSVDescription(in *operators.CSVDescription, out *CSVDescription, s conversion.Scope) error {
	return autoConvert_operators_CSVDescription_To_v1alpha1_CSVDescription(in, out, s)
}

func autoConvert_v1alpha1_Icon_To_operators_Icon(in *Icon, out *operators.Icon, s conversion.Scope) error {
	out.Base64Data = in.Base64Data
	out.Mediatype = in.Mediatype
	return nil
}

// Convert_v1alpha1_Icon_To_operators_Icon is an autogenerated conversion function.
func Convert_v1alpha1_Icon_To_operators_Icon(in *Icon, out *operators.Icon, s conversion.Scope) error {
	return autoConvert_v1alpha1_Icon_To_operators_Icon(in, out, s)
}

func autoConvert_operators_Icon_To_v1alpha1_Icon(in *operators.Icon, out *Icon, s conversion.Scope) error {
	out.Base64Data = in.Base64Data
	out.Mediatype = in.Mediatype
	return nil
}

// Convert_operators_Icon_To_v1alpha1_Icon is an autogenerated conversion function.
func Convert_operators_Icon_To_v1alpha1_Icon(in *operators.Icon, out *Icon, s conversion.Scope) error {
	return autoConvert_operators_Icon_To_v1alpha1_Icon(in, out, s)
}

func autoConvert_v1alpha1_PackageChannel_To_operators_PackageChannel(in *PackageChannel, out *operators.PackageChannel, s conversion.Scope) error {
	out.Name = in.Name
	out.CurrentCSV = in.CurrentCSV
	if err := Convert_v1alpha1_CSVDescription_To_operators_CSVDescription(&in.CurrentCSVDesc, &out.CurrentCSVDesc, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PackageChannel_To_operators_PackageChannel is an autogenerated conversion function.
func Convert_v1alpha1_PackageChannel_To_operators_PackageChannel(in *PackageChannel, out *operators.PackageChannel, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageChannel_To_operators_PackageChannel(in, out, s)
}

func autoConvert_operators_PackageChannel_To_v1alpha1_PackageChannel(in *operators.PackageChannel, out *PackageChannel, s conversion.Scope) error {
	out.Name = in.Name
	out.CurrentCSV = in.CurrentCSV
	if err := Convert_operators_CSVDescription_To_v1alpha1_CSVDescription(&in.CurrentCSVDesc, &out.CurrentCSVDesc, s); err != nil {
		return err
	}
	return nil
}

// Convert_operators_PackageChannel_To_v1alpha1_PackageChannel is an autogenerated conversion function.
func Convert_operators_PackageChannel_To_v1alpha1_PackageChannel(in *operators.PackageChannel, out *PackageChannel, s conversion.Scope) error {
	return autoConvert_operators_PackageChannel_To_v1alpha1_PackageChannel(in, out, s)
}

func autoConvert_v1alpha1_PackageManifest_To_operators_PackageManifest(in *PackageManifest, out *operators.PackageManifest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PackageManifestSpec_To_operators_PackageManifestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PackageManifestStatus_To_operators_PackageManifestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PackageManifest_To_operators_PackageManifest is an autogenerated conversion function.
func Convert_v1alpha1_PackageManifest_To_operators_PackageManifest(in *PackageManifest, out *operators.PackageManifest, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageManifest_To_operators_PackageManifest(in, out, s)
}

func autoConvert_operators_PackageManifest_To_v1alpha1_PackageManifest(in *operators.PackageManifest, out *PackageManifest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_operators_PackageManifestSpec_To_v1alpha1_PackageManifestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_operators_PackageManifestStatus_To_v1alpha1_PackageManifestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_operators_PackageManifest_To_v1alpha1_PackageManifest is an autogenerated conversion function.
func Convert_operators_PackageManifest_To_v1alpha1_PackageManifest(in *operators.PackageManifest, out *PackageManifest, s conversion.Scope) error {
	return autoConvert_operators_PackageManifest_To_v1alpha1_PackageManifest(in, out, s)
}

func autoConvert_v1alpha1_PackageManifestList_To_operators_PackageManifestList(in *PackageManifestList, out *operators.PackageManifestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]operators.PackageManifest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PackageManifestList_To_operators_PackageManifestList is an autogenerated conversion function.
func Convert_v1alpha1_PackageManifestList_To_operators_PackageManifestList(in *PackageManifestList, out *operators.PackageManifestList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageManifestList_To_operators_PackageManifestList(in, out, s)
}

func autoConvert_operators_PackageManifestList_To_v1alpha1_PackageManifestList(in *operators.PackageManifestList, out *PackageManifestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PackageManifest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_operators_PackageManifestList_To_v1alpha1_PackageManifestList is an autogenerated conversion function.
func Convert_operators_PackageManifestList_To_v1alpha1_PackageManifestList(in *operators.PackageManifestList, out *PackageManifestList, s conversion.Scope) error {
	return autoConvert_operators_PackageManifestList_To_v1alpha1_PackageManifestList(in, out, s)
}

func autoConvert_v1alpha1_PackageManifestSpec_To_operators_PackageManifestSpec(in *PackageManifestSpec, out *operators.PackageManifestSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_PackageManifestSpec_To_operators_PackageManifestSpec is an autogenerated conversion function.
func Convert_v1alpha1_PackageManifestSpec_To_operators_PackageManifestSpec(in *PackageManifestSpec, out *operators.PackageManifestSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageManifestSpec_To_operators_PackageManifestSpec(in, out, s)
}

func autoConvert_operators_PackageManifestSpec_To_v1alpha1_PackageManifestSpec(in *operators.PackageManifestSpec, out *PackageManifestSpec, s conversion.Scope) error {
	return nil
}

// Convert_operators_PackageManifestSpec_To_v1alpha1_PackageManifestSpec is an autogenerated conversion function.
func Convert_operators_PackageManifestSpec_To_v1alpha1_PackageManifestSpec(in *operators.PackageManifestSpec, out *PackageManifestSpec, s conversion.Scope) error {
	return autoConvert_operators_PackageManifestSpec_To_v1alpha1_PackageManifestSpec(in, out, s)
}

func autoConvert_v1alpha1_PackageManifestStatus_To_operators_PackageManifestStatus(in *PackageManifestStatus, out *operators.PackageManifestStatus, s conversion.Scope) error {
	out.CatalogSource = in.CatalogSource
	out.CatalogSourceDisplayName = in.CatalogSourceDisplayName
	out.CatalogSourcePublisher = in.CatalogSourcePublisher
	out.CatalogSourceNamespace = in.CatalogSourceNamespace
	if err := Convert_v1alpha1_AppLink_To_operators_AppLink(&in.Provider, &out.Provider, s); err != nil {
		return err
	}
	out.PackageName = in.PackageName
	out.Channels = *(*[]operators.PackageChannel)(unsafe.Pointer(&in.Channels))
	out.DefaultChannel = in.DefaultChannel
	return nil
}

// Convert_v1alpha1_PackageManifestStatus_To_operators_PackageManifestStatus is an autogenerated conversion function.
func Convert_v1alpha1_PackageManifestStatus_To_operators_PackageManifestStatus(in *PackageManifestStatus, out *operators.PackageManifestStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackageManifestStatus_To_operators_PackageManifestStatus(in, out, s)
}

func autoConvert_operators_PackageManifestStatus_To_v1alpha1_PackageManifestStatus(in *operators.PackageManifestStatus, out *PackageManifestStatus, s conversion.Scope) error {
	out.CatalogSource = in.CatalogSource
	out.CatalogSourceDisplayName = in.CatalogSourceDisplayName
	out.CatalogSourcePublisher = in.CatalogSourcePublisher
	out.CatalogSourceNamespace = in.CatalogSourceNamespace
	if err := Convert_operators_AppLink_To_v1alpha1_AppLink(&in.Provider, &out.Provider, s); err != nil {
		return err
	}
	out.PackageName = in.PackageName
	out.Channels = *(*[]PackageChannel)(unsafe.Pointer(&in.Channels))
	out.DefaultChannel = in.DefaultChannel
	return nil
}

// Convert_operators_PackageManifestStatus_To_v1alpha1_PackageManifestStatus is an autogenerated conversion function.
func Convert_operators_PackageManifestStatus_To_v1alpha1_PackageManifestStatus(in *operators.PackageManifestStatus, out *PackageManifestStatus, s conversion.Scope) error {
	return autoConvert_operators_PackageManifestStatus_To_v1alpha1_PackageManifestStatus(in, out, s)
}
