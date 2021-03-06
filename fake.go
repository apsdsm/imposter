//    Copyright 2017 Nick del Pozo
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package imposter

import (
	"fmt"
	"reflect"
)

// A Fake is an embeddable struct that makes it easier for fake classes to record method calls.
type Fake struct {
	calls   map[string][]interface{}
	returns map[string][]interface{}
}

// Received checks to see if the specified method was called with the given parameters. It
// returns true if it finds an exact match, and false in all other cases.
func (f *Fake) Received(method string, signature ...interface{}) bool {
	if f.calls == nil {
		panic("no calls were made to this object.")
	}

	checksiglen := len(signature)

	if checksiglen > 1 {
		panic("to check multiple calls to this method, call 'Received' multiple times.")
	}

	if sig, ok := f.calls[method]; ok {
		for _, s := range sig {
			if (checksiglen == 0 && s == nil) || s == signature[0] {
				return true
			}
		}
	} else {
		panic("did not get any calls to " + method)
	}

	rep := getSignatureString(signature)
	res := fmt.Sprintf("\n\ndid not receive calls to %s with: \n# %s \n\n", method, rep)
	res += fmt.Sprint("received calls were:")

	for i, s := range f.calls[method] {
		rep = getSignatureString(s)
		res += fmt.Sprintf("\n%d %s", i, rep)
	}

	panic(res)
}

// DidNotReceive returns true if no calls were registered against the specified method.
func (f *Fake) DidNotReceive(method string, signature ...interface{}) bool {

	checksiglen := len(signature)

	if checksiglen > 1 {
		panic("to check multiple calls to this method, call 'Received' multiple times.")
	}

	if f.calls == nil {
		return true
	}

	if sig, ok := f.calls[method]; ok {
		for _, s := range sig {
			if (checksiglen == 0 && s == nil) || s == signature[0] {
				panic("did not expect to receive to receive calls to " + method + " but one was received.")
			}
		}
	}

	return true
}

// getSignatureString returns a string representation of a signature (or any interface)
// and its contained values.
func getSignatureString(signature interface{}) (rep string) {
	sValue := reflect.ValueOf(signature)
	sTypes := sValue.Type()

	var c string

	for i := 0; i < sValue.NumField(); i++ {
		rep += c + fmt.Sprintf("%s %s = %v", sTypes.Field(i).Name, sValue.Field(i).Type(), sValue.Field(i).Interface())
		c = ", "
	}

	rep = "{ " + rep + " }"

	return rep
}

// SetCall is used to record a new call to the specified method, recording the parameters
// that were passed.
func (f *Fake) SetCall(method string, signature ...interface{}) {
	if f.calls == nil {
		f.calls = make(map[string][]interface{})
	}

	if len(signature) == 0 {
		f.calls[method] = append(f.calls[method], nil)
	} else {
		for _, s := range signature {
			f.calls[method] = append(f.calls[method], s)
		}
	}
}
