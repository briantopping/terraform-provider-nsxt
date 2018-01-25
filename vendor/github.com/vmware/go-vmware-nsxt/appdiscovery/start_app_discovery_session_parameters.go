/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package appdiscovery

// Parameters to start an application discovery session. It can have NSGroup Ids as well as the App Profile Ids.
type StartAppDiscoverySessionParameters struct {

	// App Profile Ids
	AppProfileIds []string `json:"app_profile_ids,omitempty"`

	// NSGroup Ids
	NsGroupIds []string `json:"ns_group_ids"`
}