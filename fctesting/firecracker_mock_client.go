// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package fctesting

import (
	ops "github.com/valyentdev/firecracker-go-sdk/client/operations"
)

type MockClient struct {
	CreateSnapshotFn                 func(params *ops.CreateSnapshotParams) (*ops.CreateSnapshotNoContent, error)
	CreateSyncActionFn               func(params *ops.CreateSyncActionParams) (*ops.CreateSyncActionNoContent, error)
	DescribeBalloonConfigFn          func(params *ops.DescribeBalloonConfigParams) (*ops.DescribeBalloonConfigOK, error)
	DescribeBalloonStatsFn           func(params *ops.DescribeBalloonStatsParams) (*ops.DescribeBalloonStatsOK, error)
	DescribeInstanceFn               func(params *ops.DescribeInstanceParams) (*ops.DescribeInstanceOK, error)
	GetExportVMConfigFn              func(params *ops.GetExportVMConfigParams) (*ops.GetExportVMConfigOK, error)
	GetFirecrackerVersionFn          func(params *ops.GetFirecrackerVersionParams) (*ops.GetFirecrackerVersionOK, error)
	GetMachineConfigurationFn        func(params *ops.GetMachineConfigurationParams) (*ops.GetMachineConfigurationOK, error)
	GetMmdsFn                        func(params *ops.GetMmdsParams) (*ops.GetMmdsOK, error)
	LoadSnapshotFn                   func(params *ops.LoadSnapshotParams) (*ops.LoadSnapshotNoContent, error)
	PatchBalloonFn                   func(params *ops.PatchBalloonParams) (*ops.PatchBalloonNoContent, error)
	PatchBalloonStatsIntervalFn      func(params *ops.PatchBalloonStatsIntervalParams) (*ops.PatchBalloonStatsIntervalNoContent, error)
	PatchGuestDriveByIDFn            func(params *ops.PatchGuestDriveByIDParams) (*ops.PatchGuestDriveByIDNoContent, error)
	PatchGuestNetworkInterfaceByIDFn func(params *ops.PatchGuestNetworkInterfaceByIDParams) (*ops.PatchGuestNetworkInterfaceByIDNoContent, error)
	PatchMachineConfigurationFn      func(params *ops.PatchMachineConfigurationParams) (*ops.PatchMachineConfigurationNoContent, error)
	PatchMmdsFn                      func(params *ops.PatchMmdsParams) (*ops.PatchMmdsNoContent, error)
	PatchVMFn                        func(params *ops.PatchVMParams) (*ops.PatchVMNoContent, error)
	PutBalloonFn                     func(params *ops.PutBalloonParams) (*ops.PutBalloonNoContent, error)
	PutCPUConfigurationFn            func(params *ops.PutCPUConfigurationParams) (*ops.PutCPUConfigurationNoContent, error)
	PutEntropyDeviceFn               func(params *ops.PutEntropyDeviceParams) (*ops.PutEntropyDeviceNoContent, error)
	PutGuestBootSourceFn             func(params *ops.PutGuestBootSourceParams) (*ops.PutGuestBootSourceNoContent, error)
	PutGuestDriveByIDFn              func(params *ops.PutGuestDriveByIDParams) (*ops.PutGuestDriveByIDNoContent, error)
	PutGuestNetworkInterfaceByIDFn   func(params *ops.PutGuestNetworkInterfaceByIDParams) (*ops.PutGuestNetworkInterfaceByIDNoContent, error)
	PutGuestVsockFn                  func(params *ops.PutGuestVsockParams) (*ops.PutGuestVsockNoContent, error)
	PutLoggerFn                      func(params *ops.PutLoggerParams) (*ops.PutLoggerNoContent, error)
	PutMachineConfigurationFn        func(params *ops.PutMachineConfigurationParams) (*ops.PutMachineConfigurationNoContent, error)
	PutMetricsFn                     func(params *ops.PutMetricsParams) (*ops.PutMetricsNoContent, error)
	PutMmdsFn                        func(params *ops.PutMmdsParams) (*ops.PutMmdsNoContent, error)
	PutMmdsConfigFn                  func(params *ops.PutMmdsConfigParams) (*ops.PutMmdsConfigNoContent, error)
}

func (c *MockClient) CreateSnapshot(params *ops.CreateSnapshotParams) (*ops.CreateSnapshotNoContent, error) {
	if c.CreateSnapshotFn != nil {
		return c.CreateSnapshotFn(params)
	}

	return nil, nil
}

func (c *MockClient) CreateSyncAction(params *ops.CreateSyncActionParams) (*ops.CreateSyncActionNoContent, error) {
	if c.CreateSyncActionFn != nil {
		return c.CreateSyncActionFn(params)
	}

	return nil, nil
}

func (c *MockClient) DescribeBalloonConfig(params *ops.DescribeBalloonConfigParams) (*ops.DescribeBalloonConfigOK, error) {
	if c.DescribeBalloonConfigFn != nil {
		return c.DescribeBalloonConfigFn(params)
	}

	return nil, nil
}

func (c *MockClient) DescribeBalloonStats(params *ops.DescribeBalloonStatsParams) (*ops.DescribeBalloonStatsOK, error) {
	if c.DescribeBalloonStatsFn != nil {
		return c.DescribeBalloonStatsFn(params)
	}

	return nil, nil
}

func (c *MockClient) DescribeInstance(params *ops.DescribeInstanceParams) (*ops.DescribeInstanceOK, error) {
	if c.DescribeInstanceFn != nil {
		return c.DescribeInstanceFn(params)
	}

	return nil, nil
}

func (c *MockClient) GetExportVMConfig(params *ops.GetExportVMConfigParams) (*ops.GetExportVMConfigOK, error) {
	if c.GetExportVMConfigFn != nil {
		return c.GetExportVMConfigFn(params)
	}

	return nil, nil
}

func (c *MockClient) GetFirecrackerVersion(params *ops.GetFirecrackerVersionParams) (*ops.GetFirecrackerVersionOK, error) {
	if c.GetFirecrackerVersionFn != nil {
		return c.GetFirecrackerVersionFn(params)
	}

	return nil, nil
}

func (c *MockClient) GetMachineConfiguration(params *ops.GetMachineConfigurationParams) (*ops.GetMachineConfigurationOK, error) {
	if c.GetMachineConfigurationFn != nil {
		return c.GetMachineConfigurationFn(params)
	}

	return nil, nil
}

func (c *MockClient) GetMmds(params *ops.GetMmdsParams) (*ops.GetMmdsOK, error) {
	if c.GetMmdsFn != nil {
		return c.GetMmdsFn(params)
	}

	return nil, nil
}

func (c *MockClient) LoadSnapshot(params *ops.LoadSnapshotParams) (*ops.LoadSnapshotNoContent, error) {
	if c.LoadSnapshotFn != nil {
		return c.LoadSnapshotFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchBalloon(params *ops.PatchBalloonParams) (*ops.PatchBalloonNoContent, error) {
	if c.PatchBalloonFn != nil {
		return c.PatchBalloonFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchBalloonStatsInterval(params *ops.PatchBalloonStatsIntervalParams) (*ops.PatchBalloonStatsIntervalNoContent, error) {
	if c.PatchBalloonStatsIntervalFn != nil {
		return c.PatchBalloonStatsIntervalFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchGuestDriveByID(params *ops.PatchGuestDriveByIDParams) (*ops.PatchGuestDriveByIDNoContent, error) {
	if c.PatchGuestDriveByIDFn != nil {
		return c.PatchGuestDriveByIDFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchGuestNetworkInterfaceByID(params *ops.PatchGuestNetworkInterfaceByIDParams) (*ops.PatchGuestNetworkInterfaceByIDNoContent, error) {
	if c.PatchGuestNetworkInterfaceByIDFn != nil {
		return c.PatchGuestNetworkInterfaceByIDFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchMachineConfiguration(params *ops.PatchMachineConfigurationParams) (*ops.PatchMachineConfigurationNoContent, error) {
	if c.PatchMachineConfigurationFn != nil {
		return c.PatchMachineConfigurationFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchMmds(params *ops.PatchMmdsParams) (*ops.PatchMmdsNoContent, error) {
	if c.PatchMmdsFn != nil {
		return c.PatchMmdsFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchVM(params *ops.PatchVMParams) (*ops.PatchVMNoContent, error) {
	if c.PatchVMFn != nil {
		return c.PatchVMFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutBalloon(params *ops.PutBalloonParams) (*ops.PutBalloonNoContent, error) {
	if c.PutBalloonFn != nil {
		return c.PutBalloonFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutCPUConfiguration(params *ops.PutCPUConfigurationParams) (*ops.PutCPUConfigurationNoContent, error) {
	if c.PutCPUConfigurationFn != nil {
		return c.PutCPUConfigurationFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutEntropyDevice(params *ops.PutEntropyDeviceParams) (*ops.PutEntropyDeviceNoContent, error) {
	if c.PutEntropyDeviceFn != nil {
		return c.PutEntropyDeviceFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestBootSource(params *ops.PutGuestBootSourceParams) (*ops.PutGuestBootSourceNoContent, error) {
	if c.PutGuestBootSourceFn != nil {
		return c.PutGuestBootSourceFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestDriveByID(params *ops.PutGuestDriveByIDParams) (*ops.PutGuestDriveByIDNoContent, error) {
	if c.PutGuestDriveByIDFn != nil {
		return c.PutGuestDriveByIDFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestNetworkInterfaceByID(params *ops.PutGuestNetworkInterfaceByIDParams) (*ops.PutGuestNetworkInterfaceByIDNoContent, error) {
	if c.PutGuestNetworkInterfaceByIDFn != nil {
		return c.PutGuestNetworkInterfaceByIDFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestVsock(params *ops.PutGuestVsockParams) (*ops.PutGuestVsockNoContent, error) {
	if c.PutGuestVsockFn != nil {
		return c.PutGuestVsockFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutLogger(params *ops.PutLoggerParams) (*ops.PutLoggerNoContent, error) {
	if c.PutLoggerFn != nil {
		return c.PutLoggerFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutMachineConfiguration(params *ops.PutMachineConfigurationParams) (*ops.PutMachineConfigurationNoContent, error) {
	if c.PutMachineConfigurationFn != nil {
		return c.PutMachineConfigurationFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutMetrics(params *ops.PutMetricsParams) (*ops.PutMetricsNoContent, error) {
	if c.PutMetricsFn != nil {
		return c.PutMetricsFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutMmds(params *ops.PutMmdsParams) (*ops.PutMmdsNoContent, error) {
	if c.PutMmdsFn != nil {
		return c.PutMmdsFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutMmdsConfig(params *ops.PutMmdsConfigParams) (*ops.PutMmdsConfigNoContent, error) {
	if c.PutMmdsConfigFn != nil {
		return c.PutMmdsConfigFn(params)
	}

	return nil, nil
}
