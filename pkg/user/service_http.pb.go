// Code generated by protoc-gen-go-http. DO NOT EDIT.

package user

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path:              "/keyauth.user.UserService/QueryAccount",
				FunctionName:      "QueryAccount",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "list"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/DescribeAccount",
				FunctionName:      "DescribeAccount",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "*", "action": "get"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/CreateAccount",
				FunctionName:      "CreateAccount",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "create"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/BlockAccount",
				FunctionName:      "BlockAccount",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/UnBlockAccount",
				FunctionName:      "UnBlockAccount",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/DeleteAccount",
				FunctionName:      "DeleteAccount",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "delete"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/UpdateAccountProfile",
				FunctionName:      "UpdateAccountProfile",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/UpdateAccountPassword",
				FunctionName:      "UpdateAccountPassword",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				Path:              "/keyauth.user.UserService/GeneratePassword",
				FunctionName:      "GeneratePassword",
				Method:            "",
				Resource:          "",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          false,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "*"},
				Extension:         map[string]string{},
			},
		},
	}
	return set
}
