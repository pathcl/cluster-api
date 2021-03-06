/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import (
	kubeadmbootstrapv1alpha4 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this KubeadmConfig to the Hub version (v1alpha4).
func (src *KubeadmConfig) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfig)
	return Convert_v1alpha3_KubeadmConfig_To_v1alpha4_KubeadmConfig(src, dst, nil)
}

// ConvertFrom converts from the KubeadmConfig Hub version (v1alpha4) to this version.
func (dst *KubeadmConfig) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfig)
	return Convert_v1alpha4_KubeadmConfig_To_v1alpha3_KubeadmConfig(src, dst, nil)
}

// ConvertTo converts this KubeadmConfigList to the Hub version (v1alpha4).
func (src *KubeadmConfigList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfigList)
	return Convert_v1alpha3_KubeadmConfigList_To_v1alpha4_KubeadmConfigList(src, dst, nil)
}

// ConvertFrom converts from the KubeadmConfigList Hub version (v1alpha4) to this version.
func (dst *KubeadmConfigList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfigList)
	return Convert_v1alpha4_KubeadmConfigList_To_v1alpha3_KubeadmConfigList(src, dst, nil)
}

// ConvertTo converts this KubeadmConfigTemplate to the Hub version (v1alpha4).
func (src *KubeadmConfigTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfigTemplate)
	return Convert_v1alpha3_KubeadmConfigTemplate_To_v1alpha4_KubeadmConfigTemplate(src, dst, nil)
}

// ConvertFrom converts from the KubeadmConfigTemplate Hub version (v1alpha4) to this version.
func (dst *KubeadmConfigTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfigTemplate)
	return Convert_v1alpha4_KubeadmConfigTemplate_To_v1alpha3_KubeadmConfigTemplate(src, dst, nil)
}

// ConvertTo converts this KubeadmConfigTemplateList to the Hub version (v1alpha3).
func (src *KubeadmConfigTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfigTemplateList)
	return Convert_v1alpha3_KubeadmConfigTemplateList_To_v1alpha4_KubeadmConfigTemplateList(src, dst, nil)
}

// ConvertFrom converts from the KubeadmConfigTemplateList Hub version (v1alpha3) to this version.
func (dst *KubeadmConfigTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*kubeadmbootstrapv1alpha4.KubeadmConfigTemplateList)
	return Convert_v1alpha4_KubeadmConfigTemplateList_To_v1alpha3_KubeadmConfigTemplateList(src, dst, nil)
}
