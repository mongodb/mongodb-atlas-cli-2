// Copyright 2020 MongoDB Inc
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

package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mcli/internal/fixtures"
	"github.com/mongodb/mcli/internal/mocks"
)

func TestAtlasClustersCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterCreator(ctrl)

	defer ctrl.Finish()

	expected := fixtures.Cluster()

	createOpts := &atlasClustersCreateOpts{
		globalOpts:   newGlobalOpts(),
		name:         "ProjectBar",
		region:       "US",
		instanceSize: atlasM2,
		members:      3,
		diskSizeGB:   10,
		backup:       false,
		mdbVersion:   currentMDBVersion,
		store:        mockStore,
	}

	mockStore.
		EXPECT().
		CreateCluster(createOpts.newCluster()).Return(expected, nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
