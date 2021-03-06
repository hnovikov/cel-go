// Copyright 2018 Google LLC
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

package pb

import (
	descpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	emptypb "github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/struct"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

var (
	// CheckedPrimitives map from proto field descriptor type to expr.Type.
	CheckedPrimitives = map[descpb.FieldDescriptorProto_Type]*expr.Type{
		descpb.FieldDescriptorProto_TYPE_BOOL:    checkedBool,
		descpb.FieldDescriptorProto_TYPE_BYTES:   checkedBytes,
		descpb.FieldDescriptorProto_TYPE_DOUBLE:  checkedDouble,
		descpb.FieldDescriptorProto_TYPE_FLOAT:   checkedDouble,
		descpb.FieldDescriptorProto_TYPE_INT32:   checkedInt,
		descpb.FieldDescriptorProto_TYPE_INT64:   checkedInt,
		descpb.FieldDescriptorProto_TYPE_SINT32:  checkedInt,
		descpb.FieldDescriptorProto_TYPE_SINT64:  checkedInt,
		descpb.FieldDescriptorProto_TYPE_UINT32:  checkedUint,
		descpb.FieldDescriptorProto_TYPE_UINT64:  checkedUint,
		descpb.FieldDescriptorProto_TYPE_FIXED32: checkedUint,
		descpb.FieldDescriptorProto_TYPE_FIXED64: checkedUint,
		descpb.FieldDescriptorProto_TYPE_STRING:  checkedString}

	// CheckedWellKnowns map from qualified proto type name to expr.Type for
	// well-known proto types.
	CheckedWellKnowns = map[string]*expr.Type{
		"google.protobuf.DoubleValue": checkedWrap(checkedDouble),
		"google.protobuf.FloatValue":  checkedWrap(checkedDouble),
		"google.protobuf.Int64Value":  checkedWrap(checkedInt),
		"google.protobuf.Int32Value":  checkedWrap(checkedInt),
		"google.protobuf.UInt64Value": checkedWrap(checkedUint),
		"google.protobuf.UInt32Value": checkedWrap(checkedUint),
		"google.protobuf.BoolValue":   checkedWrap(checkedBool),
		"google.protobuf.StringValue": checkedWrap(checkedString),
		"google.protobuf.BytesValue":  checkedWrap(checkedBytes),
		"google.protobuf.NullValue":   checkedNull,
		"google.protobuf.Timestamp":   checkedTimestamp,
		"google.protobuf.Duration":    checkedDuration,
		"google.protobuf.Struct":      checkedDyn,
		"google.protobuf.Value":       checkedDyn,
		"google.protobuf.ListValue":   checkedDyn,
		"google.protobuf.Any":         checkedAny}

	// common types
	checkedBool      = checkedPrimitive(expr.Type_BOOL)
	checkedBytes     = checkedPrimitive(expr.Type_BYTES)
	checkedDouble    = checkedPrimitive(expr.Type_DOUBLE)
	checkedDyn       = &expr.Type{TypeKind: &expr.Type_Dyn{Dyn: &emptypb.Empty{}}}
	checkedInt       = checkedPrimitive(expr.Type_INT64)
	checkedNull      = &expr.Type{TypeKind: &expr.Type_Null{Null: structpb.NullValue_NULL_VALUE}}
	checkedString    = checkedPrimitive(expr.Type_STRING)
	checkedUint      = checkedPrimitive(expr.Type_UINT64)
	checkedAny       = checkedWellKnown(expr.Type_ANY)
	checkedDuration  = checkedWellKnown(expr.Type_DURATION)
	checkedTimestamp = checkedWellKnown(expr.Type_TIMESTAMP)
)
