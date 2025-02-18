// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

package outgoinghandler

import (
	"github.com/rajatjindal/wasip2-string/internal/wasi/http/v0.2.0/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

func lower_OptionRequestOptions(v cm.Option[types.RequestOptions]) (f0 uint32, f1 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := cm.Reinterpret[uint32](*some)
		f1 = (uint32)(v1)
	}
	return
}
