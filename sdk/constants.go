//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package sdk

const host = "https://partner.microsoft.com/v1"

type Get0ItemsItem string

const (
	Get0ItemsItemNone Get0ItemsItem = "none"
	Get0ItemsItemRunning Get0ItemsItem = "running"
	Get0ItemsItemScheduled Get0ItemsItem = "scheduled"
	Get0ItemsItemSuccess Get0ItemsItem = "success"
)

// PossibleGet0ItemsItemValues returns the possible values for the Get0ItemsItem const type.
func PossibleGet0ItemsItemValues() []Get0ItemsItem {
	return []Get0ItemsItem{	
		Get0ItemsItemNone,
		Get0ItemsItemRunning,
		Get0ItemsItemScheduled,
		Get0ItemsItemSuccess,
	}
}

