// Copyright The OpenTelemetry Authors
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

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "paging"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	SystemPagingFaults     MetricIntf
	SystemPagingOperations MetricIntf
	SystemPagingUsage      MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"system.paging.faults",
		"system.paging.operations",
		"system.paging.usage",
	}
}

var metricsByName = map[string]MetricIntf{
	"system.paging.faults":     Metrics.SystemPagingFaults,
	"system.paging.operations": Metrics.SystemPagingOperations,
	"system.paging.usage":      Metrics.SystemPagingUsage,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"system.paging.faults",
		func(metric pdata.Metric) {
			metric.SetName("system.paging.faults")
			metric.SetDescription("The number of page faults.")
			metric.SetUnit("{faults}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.paging.operations",
		func(metric pdata.Metric) {
			metric.SetName("system.paging.operations")
			metric.SetDescription("The number of paging operations.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.paging.usage",
		func(metric pdata.Metric) {
			metric.SetName("system.paging.usage")
			metric.SetDescription("Swap (unix) or pagefile (windows) usage.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// Device (Name of the page file.)
	Device string
	// Direction (Page In or Page Out.)
	Direction string
	// State (Breakdown of paging usage by type.)
	State string
	// Type (Type of fault.)
	Type string
}{
	"device",
	"direction",
	"state",
	"type",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels

// LabelDirection are the possible values that the label "direction" can have.
var LabelDirection = struct {
	PageIn  string
	PageOut string
}{
	"page_in",
	"page_out",
}

// LabelState are the possible values that the label "state" can have.
var LabelState = struct {
	Cached string
	Free   string
	Used   string
}{
	"cached",
	"free",
	"used",
}

// LabelType are the possible values that the label "type" can have.
var LabelType = struct {
	Major string
	Minor string
}{
	"major",
	"minor",
}