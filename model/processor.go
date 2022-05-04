// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package model

import (
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// Processor identifies an event type, and is used for routing events
// to the appropriate data stream or index.
//
// TODO(axw) this should be replaced with ECS event categorisation fields.
type Processor struct {
	Name  string
	Event string
}

func (p *Processor) fields() mapstr.M {
	var fields mapStr
	fields.maybeSetString("name", p.Name)
	fields.maybeSetString("event", p.Event)
	return mapstr.M(fields)
}
