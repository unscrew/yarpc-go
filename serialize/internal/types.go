// Code generated by thriftrw v1.19.1. DO NOT EDIT.
// @generated

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package internal

import (
	bytes "bytes"
	base64 "encoding/base64"
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type RPC struct {
	SpanContext     []byte            `json:"spanContext,required"`
	CallerName      string            `json:"callerName,required"`
	ServiceName     string            `json:"serviceName,required"`
	Encoding        string            `json:"encoding,required"`
	Procedure       string            `json:"procedure,required"`
	Headers         map[string]string `json:"headers,omitempty"`
	ShardKey        *string           `json:"shardKey,omitempty"`
	RoutingKey      *string           `json:"routingKey,omitempty"`
	RoutingDelegate *string           `json:"routingDelegate,omitempty"`
	Body            []byte            `json:"body,omitempty"`
}

type _Map_String_String_MapItemList map[string]string

func (m _Map_String_String_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}

		vw, err := wire.NewValueString(v), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_String_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_String_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_String_MapItemList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Map_String_String_MapItemList) Close() {}

// ToWire translates a RPC struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *RPC) ToWire() (wire.Value, error) {
	var (
		fields [10]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.SpanContext == nil {
		return w, errors.New("field SpanContext of RPC is required")
	}
	w, err = wire.NewValueBinary(v.SpanContext), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueString(v.CallerName), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	w, err = wire.NewValueString(v.ServiceName), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++

	w, err = wire.NewValueString(v.Encoding), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 4, Value: w}
	i++

	w, err = wire.NewValueString(v.Procedure), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++
	if v.Headers != nil {
		w, err = wire.NewValueMap(_Map_String_String_MapItemList(v.Headers)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}
	if v.ShardKey != nil {
		w, err = wire.NewValueString(*(v.ShardKey)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 7, Value: w}
		i++
	}
	if v.RoutingKey != nil {
		w, err = wire.NewValueString(*(v.RoutingKey)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 8, Value: w}
		i++
	}
	if v.RoutingDelegate != nil {
		w, err = wire.NewValueString(*(v.RoutingDelegate)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 9, Value: w}
		i++
	}
	if v.Body != nil {
		w, err = wire.NewValueBinary(v.Body), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_String_String_Read(m wire.MapItemList) (map[string]string, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}

	if m.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make(map[string]string, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}

		v, err := x.Value.GetString(), error(nil)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a RPC struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a RPC struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v RPC
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *RPC) FromWire(w wire.Value) error {
	var err error

	spanContextIsSet := false
	callerNameIsSet := false
	serviceNameIsSet := false
	encodingIsSet := false
	procedureIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.SpanContext, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}
				spanContextIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.CallerName, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				callerNameIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				v.ServiceName, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				serviceNameIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TBinary {
				v.Encoding, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				encodingIsSet = true
			}
		case 5:
			if field.Value.Type() == wire.TBinary {
				v.Procedure, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				procedureIsSet = true
			}
		case 6:
			if field.Value.Type() == wire.TMap {
				v.Headers, err = _Map_String_String_Read(field.Value.GetMap())
				if err != nil {
					return err
				}

			}
		case 7:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.ShardKey = &x
				if err != nil {
					return err
				}

			}
		case 8:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.RoutingKey = &x
				if err != nil {
					return err
				}

			}
		case 9:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.RoutingDelegate = &x
				if err != nil {
					return err
				}

			}
		case 10:
			if field.Value.Type() == wire.TBinary {
				v.Body, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}

			}
		}
	}

	if !spanContextIsSet {
		return errors.New("field SpanContext of RPC is required")
	}

	if !callerNameIsSet {
		return errors.New("field CallerName of RPC is required")
	}

	if !serviceNameIsSet {
		return errors.New("field ServiceName of RPC is required")
	}

	if !encodingIsSet {
		return errors.New("field Encoding of RPC is required")
	}

	if !procedureIsSet {
		return errors.New("field Procedure of RPC is required")
	}

	return nil
}

// String returns a readable string representation of a RPC
// struct.
func (v *RPC) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [10]string
	i := 0
	fields[i] = fmt.Sprintf("SpanContext: %v", v.SpanContext)
	i++
	fields[i] = fmt.Sprintf("CallerName: %v", v.CallerName)
	i++
	fields[i] = fmt.Sprintf("ServiceName: %v", v.ServiceName)
	i++
	fields[i] = fmt.Sprintf("Encoding: %v", v.Encoding)
	i++
	fields[i] = fmt.Sprintf("Procedure: %v", v.Procedure)
	i++
	if v.Headers != nil {
		fields[i] = fmt.Sprintf("Headers: %v", v.Headers)
		i++
	}
	if v.ShardKey != nil {
		fields[i] = fmt.Sprintf("ShardKey: %v", *(v.ShardKey))
		i++
	}
	if v.RoutingKey != nil {
		fields[i] = fmt.Sprintf("RoutingKey: %v", *(v.RoutingKey))
		i++
	}
	if v.RoutingDelegate != nil {
		fields[i] = fmt.Sprintf("RoutingDelegate: %v", *(v.RoutingDelegate))
		i++
	}
	if v.Body != nil {
		fields[i] = fmt.Sprintf("Body: %v", v.Body)
		i++
	}

	return fmt.Sprintf("RPC{%v}", strings.Join(fields[:i], ", "))
}

func _Map_String_String_Equals(lhs, rhs map[string]string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !(lv == rv) {
			return false
		}
	}
	return true
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this RPC match the
// provided RPC.
//
// This function performs a deep comparison.
func (v *RPC) Equals(rhs *RPC) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !bytes.Equal(v.SpanContext, rhs.SpanContext) {
		return false
	}
	if !(v.CallerName == rhs.CallerName) {
		return false
	}
	if !(v.ServiceName == rhs.ServiceName) {
		return false
	}
	if !(v.Encoding == rhs.Encoding) {
		return false
	}
	if !(v.Procedure == rhs.Procedure) {
		return false
	}
	if !((v.Headers == nil && rhs.Headers == nil) || (v.Headers != nil && rhs.Headers != nil && _Map_String_String_Equals(v.Headers, rhs.Headers))) {
		return false
	}
	if !_String_EqualsPtr(v.ShardKey, rhs.ShardKey) {
		return false
	}
	if !_String_EqualsPtr(v.RoutingKey, rhs.RoutingKey) {
		return false
	}
	if !_String_EqualsPtr(v.RoutingDelegate, rhs.RoutingDelegate) {
		return false
	}
	if !((v.Body == nil && rhs.Body == nil) || (v.Body != nil && rhs.Body != nil && bytes.Equal(v.Body, rhs.Body))) {
		return false
	}

	return true
}

type _Map_String_String_Zapper map[string]string

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of _Map_String_String_Zapper.
func (m _Map_String_String_Zapper) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	for k, v := range m {
		enc.AddString((string)(k), v)
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of RPC.
func (v *RPC) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("spanContext", base64.StdEncoding.EncodeToString(v.SpanContext))
	enc.AddString("callerName", v.CallerName)
	enc.AddString("serviceName", v.ServiceName)
	enc.AddString("encoding", v.Encoding)
	enc.AddString("procedure", v.Procedure)
	if v.Headers != nil {
		err = multierr.Append(err, enc.AddObject("headers", (_Map_String_String_Zapper)(v.Headers)))
	}
	if v.ShardKey != nil {
		enc.AddString("shardKey", *v.ShardKey)
	}
	if v.RoutingKey != nil {
		enc.AddString("routingKey", *v.RoutingKey)
	}
	if v.RoutingDelegate != nil {
		enc.AddString("routingDelegate", *v.RoutingDelegate)
	}
	if v.Body != nil {
		enc.AddString("body", base64.StdEncoding.EncodeToString(v.Body))
	}
	return err
}

// GetSpanContext returns the value of SpanContext if it is set or its
// zero value if it is unset.
func (v *RPC) GetSpanContext() (o []byte) {
	if v != nil {
		o = v.SpanContext
	}
	return
}

// IsSetSpanContext returns true if SpanContext is not nil.
func (v *RPC) IsSetSpanContext() bool {
	return v != nil && v.SpanContext != nil
}

// GetCallerName returns the value of CallerName if it is set or its
// zero value if it is unset.
func (v *RPC) GetCallerName() (o string) {
	if v != nil {
		o = v.CallerName
	}
	return
}

// GetServiceName returns the value of ServiceName if it is set or its
// zero value if it is unset.
func (v *RPC) GetServiceName() (o string) {
	if v != nil {
		o = v.ServiceName
	}
	return
}

// GetEncoding returns the value of Encoding if it is set or its
// zero value if it is unset.
func (v *RPC) GetEncoding() (o string) {
	if v != nil {
		o = v.Encoding
	}
	return
}

// GetProcedure returns the value of Procedure if it is set or its
// zero value if it is unset.
func (v *RPC) GetProcedure() (o string) {
	if v != nil {
		o = v.Procedure
	}
	return
}

// GetHeaders returns the value of Headers if it is set or its
// zero value if it is unset.
func (v *RPC) GetHeaders() (o map[string]string) {
	if v != nil && v.Headers != nil {
		return v.Headers
	}

	return
}

// IsSetHeaders returns true if Headers is not nil.
func (v *RPC) IsSetHeaders() bool {
	return v != nil && v.Headers != nil
}

// GetShardKey returns the value of ShardKey if it is set or its
// zero value if it is unset.
func (v *RPC) GetShardKey() (o string) {
	if v != nil && v.ShardKey != nil {
		return *v.ShardKey
	}

	return
}

// IsSetShardKey returns true if ShardKey is not nil.
func (v *RPC) IsSetShardKey() bool {
	return v != nil && v.ShardKey != nil
}

// GetRoutingKey returns the value of RoutingKey if it is set or its
// zero value if it is unset.
func (v *RPC) GetRoutingKey() (o string) {
	if v != nil && v.RoutingKey != nil {
		return *v.RoutingKey
	}

	return
}

// IsSetRoutingKey returns true if RoutingKey is not nil.
func (v *RPC) IsSetRoutingKey() bool {
	return v != nil && v.RoutingKey != nil
}

// GetRoutingDelegate returns the value of RoutingDelegate if it is set or its
// zero value if it is unset.
func (v *RPC) GetRoutingDelegate() (o string) {
	if v != nil && v.RoutingDelegate != nil {
		return *v.RoutingDelegate
	}

	return
}

// IsSetRoutingDelegate returns true if RoutingDelegate is not nil.
func (v *RPC) IsSetRoutingDelegate() bool {
	return v != nil && v.RoutingDelegate != nil
}

// GetBody returns the value of Body if it is set or its
// zero value if it is unset.
func (v *RPC) GetBody() (o []byte) {
	if v != nil && v.Body != nil {
		return v.Body
	}

	return
}

// IsSetBody returns true if Body is not nil.
func (v *RPC) IsSetBody() bool {
	return v != nil && v.Body != nil
}
