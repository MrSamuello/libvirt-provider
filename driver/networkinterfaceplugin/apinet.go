// Copyright 2023 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkinterfaceplugin

import (
	"fmt"

	virtletnetworkinterface "github.com/onmetal/libvirt-driver/pkg/plugins/networkinterface"
	"github.com/onmetal/libvirt-driver/pkg/plugins/networkinterface/apinet"
	"github.com/spf13/pflag"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

type apinetOptions struct {
	APInetNodeName string
}

func (o *apinetOptions) PluginName() string {
	return "apinet"
}

func (o *apinetOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.APInetNodeName, "apinet-node-name", "", "APInet node name")
}

func (o *apinetOptions) NetworkInterfacePlugin() (virtletnetworkinterface.Plugin, func(), error) {
	if o.APInetNodeName == "" {
		return nil, nil, fmt.Errorf("must specify apinet-node-name")
	}

	return apinet.NewPlugin(o.APInetNodeName), nil, nil
}

func init() {
	utilruntime.Must(DefaultPluginTypeRegistry.Register(&apinetOptions{}, 1))
}