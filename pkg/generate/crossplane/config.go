// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package crossplane

import "github.com/aws-controllers-k8s/code-generator/pkg/generate/config"

// DefaultConfig is the default config object for Crossplane controllers.
var DefaultConfig = config.Config{
	PrefixConfig: config.PrefixConfig{
		SpecField:   ".Spec.ForProvider",
		StatusField: ".Status.AtProvider",
	},
	IncludeACKMetadata:             false,
	SetManyOutputNotFoundErrReturn: "return cr",
}
