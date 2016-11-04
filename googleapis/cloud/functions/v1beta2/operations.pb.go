// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/cloud/functions/v1beta2/operations.proto
// DO NOT EDIT!

package google_cloud_functions_v1beta2 // import "google.golang.org/genproto/googleapis/cloud/functions/v1beta2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A type of an operation.
type OperationType int32

const (
	// Unknown operation type.
	OperationType_OPERATION_UNSPECIFIED OperationType = 0
	// Triggered by CreateFunction call
	OperationType_CREATE_FUNCTION OperationType = 1
	// Triggered by UpdateFunction call
	OperationType_UPDATE_FUNCTION OperationType = 2
	// Triggered by DeleteFunction call.
	OperationType_DELETE_FUNCTION OperationType = 3
)

var OperationType_name = map[int32]string{
	0: "OPERATION_UNSPECIFIED",
	1: "CREATE_FUNCTION",
	2: "UPDATE_FUNCTION",
	3: "DELETE_FUNCTION",
}
var OperationType_value = map[string]int32{
	"OPERATION_UNSPECIFIED": 0,
	"CREATE_FUNCTION":       1,
	"UPDATE_FUNCTION":       2,
	"DELETE_FUNCTION":       3,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Metadata describing an [Operation][google.longrunning.Operation]
type OperationMetadataV1Beta2 struct {
	// Target of the operation - for example
	// projects/project-1/locations/region-1/functions/function-1
	Target string `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	// Type of operation.
	Type OperationType `protobuf:"varint,2,opt,name=type,enum=google.cloud.functions.v1beta2.OperationType" json:"type,omitempty"`
	// The original request that started the operation.
	Request *google_protobuf1.Any `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
}

func (m *OperationMetadataV1Beta2) Reset()                    { *m = OperationMetadataV1Beta2{} }
func (m *OperationMetadataV1Beta2) String() string            { return proto.CompactTextString(m) }
func (*OperationMetadataV1Beta2) ProtoMessage()               {}
func (*OperationMetadataV1Beta2) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *OperationMetadataV1Beta2) GetRequest() *google_protobuf1.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationMetadataV1Beta2)(nil), "google.cloud.functions.v1beta2.OperationMetadataV1Beta2")
	proto.RegisterEnum("google.cloud.functions.v1beta2.OperationType", OperationType_name, OperationType_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/cloud/functions/v1beta2/operations.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x8f, 0x5f, 0x4f, 0xf2, 0x30,
	0x18, 0xc5, 0xdf, 0xc2, 0x1b, 0x8c, 0x35, 0x2a, 0x99, 0x7f, 0x32, 0xb9, 0x30, 0x84, 0x2b, 0x62,
	0x62, 0x1b, 0xf0, 0x0b, 0x38, 0xa0, 0x24, 0x4b, 0x74, 0x2c, 0x13, 0xbc, 0x25, 0xdd, 0x28, 0xb5,
	0x09, 0xb4, 0x75, 0xeb, 0x48, 0xf6, 0x81, 0xfc, 0x9e, 0xa6, 0x83, 0x11, 0xb8, 0x51, 0x2f, 0x7a,
	0xd1, 0xf3, 0x3c, 0xe7, 0xf7, 0x9c, 0x03, 0x03, 0xae, 0x14, 0x5f, 0x31, 0xc4, 0xd5, 0x8a, 0x4a,
	0x8e, 0x54, 0xca, 0x31, 0x67, 0x52, 0xa7, 0xca, 0x28, 0xbc, 0x1d, 0x51, 0x2d, 0x32, 0x9c, 0xac,
	0x54, 0xbe, 0xc0, 0xcb, 0x5c, 0x26, 0x46, 0x28, 0x99, 0xe1, 0x4d, 0x2f, 0x66, 0x86, 0xf6, 0xb1,
	0xd2, 0x2c, 0xa5, 0xa5, 0x84, 0x4a, 0x8f, 0x73, 0xbf, 0xe3, 0x95, 0x06, 0xb4, 0x37, 0xa0, 0x9d,
	0xa1, 0xe5, 0xff, 0xed, 0x1e, 0xd5, 0x02, 0x67, 0x2c, 0xdd, 0x88, 0x84, 0x25, 0x4a, 0x2e, 0x05,
	0xc7, 0x54, 0x4a, 0x65, 0x0e, 0x4f, 0xb5, 0x30, 0x17, 0xe6, 0x23, 0x8f, 0x51, 0xa2, 0xd6, 0x78,
	0x8b, 0xc3, 0xe5, 0x20, 0xce, 0x97, 0x58, 0x9b, 0x42, 0xb3, 0x0c, 0x53, 0x59, 0xd8, 0xb7, 0x35,
	0x74, 0xbe, 0x00, 0x74, 0x27, 0x55, 0xe0, 0x57, 0x66, 0xe8, 0x82, 0x1a, 0xfa, 0xde, 0x1b, 0xd8,
	0x60, 0xce, 0x2d, 0x6c, 0x18, 0x9a, 0x72, 0x66, 0x5c, 0xd0, 0x06, 0xdd, 0xd3, 0x68, 0xf7, 0x73,
	0x3c, 0xf8, 0xdf, 0xb2, 0xdc, 0x5a, 0x1b, 0x74, 0x2f, 0xfa, 0x8f, 0xe8, 0xe7, 0x7e, 0x68, 0xcf,
	0x9f, 0x16, 0x9a, 0x45, 0xa5, 0xd5, 0x41, 0xf0, 0x24, 0x65, 0x9f, 0x39, 0xcb, 0x8c, 0x5b, 0x6f,
	0x83, 0xee, 0x59, 0xff, 0xba, 0xa2, 0x54, 0x79, 0x91, 0x27, 0x8b, 0xa8, 0x5a, 0x7a, 0x10, 0xf0,
	0xfc, 0x08, 0xe3, 0xdc, 0xc1, 0x9b, 0x49, 0x48, 0x22, 0x6f, 0xea, 0x4f, 0x82, 0xf9, 0x2c, 0x78,
	0x0b, 0xc9, 0xd0, 0x1f, 0xfb, 0x64, 0xd4, 0xfc, 0xe7, 0x5c, 0xc1, 0xcb, 0x61, 0x44, 0xbc, 0x29,
	0x99, 0x8f, 0x67, 0xc1, 0xd0, 0x2e, 0x34, 0x81, 0x15, 0x67, 0xe1, 0xe8, 0x48, 0xac, 0x59, 0x71,
	0x44, 0x5e, 0xc8, 0xa1, 0x58, 0x1f, 0x3c, 0xc3, 0x4e, 0xa2, 0xd6, 0xbf, 0x94, 0x1a, 0xb8, 0xe3,
	0x4a, 0xda, 0xe7, 0xca, 0x42, 0x1b, 0x3d, 0x04, 0x71, 0xa3, 0xec, 0xf0, 0xf4, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xd6, 0x4b, 0x8f, 0xf4, 0x49, 0x02, 0x00, 0x00,
}
