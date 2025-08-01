// Copyright (c) 2020-2025 Tigera, Inc. All rights reserved.

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

// This file is here so that we can export a constructor to be used by the tests in the clusterconnection_test package, but since
// this is an _test file it will only be available for when running tests.

package clusterconnection

import (
	"context"

	"github.com/tigera/operator/pkg/controller/utils"

	operatorv1 "github.com/tigera/operator/api/v1"
	"github.com/tigera/operator/pkg/controller/options"
	"github.com/tigera/operator/pkg/controller/status"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func NewReconcilerWithShims(
	cli client.Client,
	schema *runtime.Scheme,
	status status.StatusManager,
	provider operatorv1.Provider,
	tierWatchReady *utils.ReadyFlag,
	clusterInfoWatchReady *utils.ReadyFlag,
) reconcile.Reconciler {
	opts := options.AddOptions{
		ShutdownContext: context.Background(),
	}

	return newReconciler(cli, schema, status, provider, tierWatchReady, clusterInfoWatchReady, opts)
}
