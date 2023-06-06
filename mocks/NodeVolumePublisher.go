/*
 *
 * Copyright © 2021-2023 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	fs "github.com/dell/csi-powerstore/v2/pkg/common/fs"

	logrus "github.com/sirupsen/logrus"

	mock "github.com/stretchr/testify/mock"
)

// NodeVolumePublisher is an autogenerated mock type for the NodeVolumePublisher type
type NodeVolumePublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, logFields, _a2, cap, isRO, targetPath, stagingPath
func (_m *NodeVolumePublisher) Publish(ctx context.Context, logFields logrus.Fields, _a2 fs.Interface, cap *csi.VolumeCapability, isRO bool, targetPath string, stagingPath string) (*csi.NodePublishVolumeResponse, error) {
	ret := _m.Called(ctx, logFields, _a2, cap, isRO, targetPath, stagingPath)

	var r0 *csi.NodePublishVolumeResponse
	if rf, ok := ret.Get(0).(func(context.Context, logrus.Fields, fs.Interface, *csi.VolumeCapability, bool, string, string) *csi.NodePublishVolumeResponse); ok {
		r0 = rf(ctx, logFields, _a2, cap, isRO, targetPath, stagingPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*csi.NodePublishVolumeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, logrus.Fields, fs.Interface, *csi.VolumeCapability, bool, string, string) error); ok {
		r1 = rf(ctx, logFields, _a2, cap, isRO, targetPath, stagingPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
