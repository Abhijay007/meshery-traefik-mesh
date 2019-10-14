// Copyright 2019 Layer5.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maesh

import "github.com/layer5io/meshery-maesh/meshes"

type supportedOperation struct {
	// a friendly name
	name string
	// the template file name
	templateName string
	opType       meshes.OpCategory
}

const (
	customOpCommand         = "custom"
	runVet                  = "maesh_vet"
	installMaeshCommand     = "maesh_install"
	installmTLSMaeshCommand = "maesh_mtls_install"
	installBookInfoCommand  = "install_book_info"
	cbCommand               = "cb1"
	installSMI              = "install_smi"
	installHTTPBin          = "install_http_bin"
)

var supportedOps = map[string]supportedOperation{
	installMaeshCommand: {
		name: "Latest Maesh without mTLS",
		// templateName: "install_maesh.tmpl",
		opType: meshes.OpCategory_INSTALL,
	},
	installmTLSMaeshCommand: {
		name:   "Latest Maesh with mTLS",
		opType: meshes.OpCategory_INSTALL,
	},
	installBookInfoCommand: {
		name: "Book Info Application",
		// templateName: "install_maesh.tmpl",
		opType: meshes.OpCategory_SAMPLE_APPLICATION,
	},
	runVet: {
		name:   "Run maesh-vet",
		opType: meshes.OpCategory_VALIDATE,
		// templateName: "maesh_vet.tmpl",
		// appLabel:     "maesh-vet",
		// returnLogs:   true,
	},
	cbCommand: {
		name:         "Configure circuit breaker with only one connection",
		opType:       meshes.OpCategory_CONFIGURE,
		templateName: "circuit_breaking.tmpl",
	},
	installSMI: {
		name:   "Service Mesh Interface (SMI) Maesh Adapter",
		opType: meshes.OpCategory_INSTALL,
	},
	installHTTPBin: {
		name:         "HTTPbin Application",
		templateName: "httpbin.yaml",
		opType:       meshes.OpCategory_SAMPLE_APPLICATION,
	},
	customOpCommand: {
		name:   "Custom YAML",
		opType: meshes.OpCategory_CUSTOM,
	},
}
