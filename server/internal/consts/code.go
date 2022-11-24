// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package consts

// 全局状态码
const (
	CodeNil                      = -1  // No error code specified.
	CodeOK                       = 0   // It is OK.
	CodeInternalError            = 50  // An error occurred internally.
	CodeValidationFailed         = 51  // Data validation failed.
	CodeDbOperationError         = 52  // Database operation error.
	CodeInvalidParameter         = 53  // The given parameter for current operation is invalid.
	CodeMissingParameter         = 54  // Parameter for current operation is missing.
	CodeInvalidOperation         = 55  // The function cannot be used like this.
	CodeInvalidConfiguration     = 56  // The configuration is invalid for current operation.
	CodeMissingConfiguration     = 57  // The configuration is missing for current operation.
	CodeNotImplemented           = 58  // The operation is not implemented yet.
	CodeNotSupported             = 59  // The operation is not supported yet.
	CodeOperationFailed          = 60  // I tried, but I cannot give you what you want.
	CodeNotAuthorized            = 61  // Not Authorized.
	CodeSecurityReason           = 62  // Security Reason.
	CodeServerBusy               = 63  // Server is busy, please try again later.
	CodeUnknown                  = 64  // Unknown error.
	CodeNotFound                 = 65  // Resource does not exist.
	CodeInvalidRequest           = 66  // Invalid request.
	CodeBusinessValidationFailed = 300 // Business validation failed.
)
