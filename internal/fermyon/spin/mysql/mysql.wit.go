// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package mysql represents the imported interface "fermyon:spin/mysql".
package mysql

import (
	rdbmstypes "github.com/rajatjindal/wasip2-string/internal/fermyon/spin/rdbms-types"
	"github.com/ydnar/wasm-tools-go/cm"
)

// MysqlError represents the variant "fermyon:spin/mysql#mysql-error".
//
// General purpose error.
//
//	variant mysql-error {
//		success,
//		connection-failed(string),
//		bad-parameter(string),
//		query-failed(string),
//		value-conversion-failed(string),
//		other-error(string),
//	}
type MysqlError cm.Variant[uint8, string, string]

// MysqlErrorSuccess returns a [MysqlError] of case "success".
func MysqlErrorSuccess() MysqlError {
	var data struct{}
	return cm.New[MysqlError](0, data)
}

// Success returns true if [MysqlError] represents the variant case "success".
func (self *MysqlError) Success() bool {
	return self.Tag() == 0
}

// MysqlErrorConnectionFailed returns a [MysqlError] of case "connection-failed".
func MysqlErrorConnectionFailed(data string) MysqlError {
	return cm.New[MysqlError](1, data)
}

// ConnectionFailed returns a non-nil *[string] if [MysqlError] represents the variant case "connection-failed".
func (self *MysqlError) ConnectionFailed() *string {
	return cm.Case[string](self, 1)
}

// MysqlErrorBadParameter returns a [MysqlError] of case "bad-parameter".
func MysqlErrorBadParameter(data string) MysqlError {
	return cm.New[MysqlError](2, data)
}

// BadParameter returns a non-nil *[string] if [MysqlError] represents the variant case "bad-parameter".
func (self *MysqlError) BadParameter() *string {
	return cm.Case[string](self, 2)
}

// MysqlErrorQueryFailed returns a [MysqlError] of case "query-failed".
func MysqlErrorQueryFailed(data string) MysqlError {
	return cm.New[MysqlError](3, data)
}

// QueryFailed returns a non-nil *[string] if [MysqlError] represents the variant case "query-failed".
func (self *MysqlError) QueryFailed() *string {
	return cm.Case[string](self, 3)
}

// MysqlErrorValueConversionFailed returns a [MysqlError] of case "value-conversion-failed".
func MysqlErrorValueConversionFailed(data string) MysqlError {
	return cm.New[MysqlError](4, data)
}

// ValueConversionFailed returns a non-nil *[string] if [MysqlError] represents the variant case "value-conversion-failed".
func (self *MysqlError) ValueConversionFailed() *string {
	return cm.Case[string](self, 4)
}

// MysqlErrorOtherError returns a [MysqlError] of case "other-error".
func MysqlErrorOtherError(data string) MysqlError {
	return cm.New[MysqlError](5, data)
}

// OtherError returns a non-nil *[string] if [MysqlError] represents the variant case "other-error".
func (self *MysqlError) OtherError() *string {
	return cm.Case[string](self, 5)
}

// Query represents the imported function "query".
//
// query the database: select
//
//	query: func(address: string, statement: string, params: list<parameter-value>)
//	-> result<row-set, mysql-error>
//
//go:nosplit
func Query(address string, statement string, params cm.List[rdbmstypes.ParameterValue]) (result cm.OKResult[rdbmstypes.RowSet, MysqlError]) {
	address0, address1 := cm.LowerString(address)
	statement0, statement1 := cm.LowerString(statement)
	params0, params1 := cm.LowerList(params)
	wasmimport_Query((*uint8)(address0), (uint32)(address1), (*uint8)(statement0), (uint32)(statement1), (*rdbmstypes.ParameterValue)(params0), (uint32)(params1), &result)
	return
}

//go:wasmimport fermyon:spin/mysql query
//go:noescape
func wasmimport_Query(address0 *uint8, address1 uint32, statement0 *uint8, statement1 uint32, params0 *rdbmstypes.ParameterValue, params1 uint32, result *cm.OKResult[rdbmstypes.RowSet, MysqlError])

// Execute represents the imported function "execute".
//
// execute command to the database: insert, update, delete
//
//	execute: func(address: string, statement: string, params: list<parameter-value>)
//	-> result<_, mysql-error>
//
//go:nosplit
func Execute(address string, statement string, params cm.List[rdbmstypes.ParameterValue]) (result cm.ErrResult[struct{}, MysqlError]) {
	address0, address1 := cm.LowerString(address)
	statement0, statement1 := cm.LowerString(statement)
	params0, params1 := cm.LowerList(params)
	wasmimport_Execute((*uint8)(address0), (uint32)(address1), (*uint8)(statement0), (uint32)(statement1), (*rdbmstypes.ParameterValue)(params0), (uint32)(params1), &result)
	return
}

//go:wasmimport fermyon:spin/mysql execute
//go:noescape
func wasmimport_Execute(address0 *uint8, address1 uint32, statement0 *uint8, statement1 uint32, params0 *rdbmstypes.ParameterValue, params1 uint32, result *cm.ErrResult[struct{}, MysqlError])
