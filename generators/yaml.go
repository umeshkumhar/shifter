/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package generator

import (
	//"encoding/json"
	"fmt"
	runtime "k8s.io/apimachinery/pkg/runtime"
	//"log"
)

func Yaml(path string, objects []runtime.Object) {

	for k, v := range objects {

		fmt.Println(k, v)
		serializer(v)
	}

}