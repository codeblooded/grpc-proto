// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/testing/worker_service.proto

package testing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("grpc/testing/worker_service.proto", fileDescriptor_1ea9e34632114fc1) }

var fileDescriptor_1ea9e34632114fc1 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x14, 0x45, 0x59, 0x0b, 0x61, 0x83, 0x36, 0xa9, 0x34, 0x9d, 0x9d, 0x8d, 0x89, 0xa8, 0x20, 0x96,
	0xee, 0xe0, 0x07, 0xb8, 0x0b, 0x0a, 0x36, 0xe2, 0x24, 0x8f, 0x18, 0x1c, 0xf3, 0xc6, 0x97, 0x17,
	0xfd, 0x14, 0x7f, 0x57, 0x26, 0x99, 0x62, 0x76, 0x9c, 0x2e, 0xb9, 0xf7, 0xe4, 0x40, 0xae, 0x38,
	0xf3, 0xd4, 0x5b, 0xc3, 0x90, 0x38, 0x44, 0x6f, 0x7e, 0x90, 0x3e, 0x80, 0x5e, 0x13, 0xd0, 0x77,
	0xb0, 0xa0, 0x7b, 0x42, 0x46, 0x79, 0x34, 0x20, 0x7a, 0x44, 0x94, 0xda, 0x7b, 0x60, 0x31, 0x32,
	0x61, 0x57, 0xc9, 0xab, 0xdf, 0x03, 0x71, 0xfc, 0x5c, 0x14, 0xbb, 0x6a, 0x90, 0x0f, 0x62, 0xbd,
	0xcd, 0x71, 0xb8, 0x01, 0xc9, 0x13, 0x3d, 0x35, 0xe9, 0x9a, 0xde, 0x93, 0x4f, 0x4a, 0x2d, 0x35,
	0x3b, 0x7e, 0xe3, 0x9c, 0xce, 0x57, 0x97, 0xab, 0x51, 0xd3, 0x74, 0x01, 0x22, 0xcf, 0x35, 0x35,
	0x5d, 0xd2, 0xd4, 0x66, 0xa2, 0xd9, 0x88, 0x75, 0x83, 0x04, 0x0d, 0xe6, 0xc8, 0xf2, 0x74, 0x06,
	0x23, 0xc1, 0x16, 0xbe, 0x32, 0x24, 0xfe, 0xe7, 0x29, 0x55, 0xea, 0x31, 0x26, 0x90, 0x37, 0x42,
	0x3c, 0xe6, 0xc0, 0xf5, 0x9b, 0x52, 0xee, 0x93, 0x4f, 0x18, 0x9c, 0x5a, 0xc8, 0x36, 0x77, 0x2f,
	0xb7, 0x3e, 0xf0, 0x7b, 0x6e, 0xb5, 0xc5, 0x4f, 0x63, 0xd1, 0x41, 0xdb, 0x21, 0x3a, 0x70, 0x66,
	0x60, 0x2f, 0xca, 0x7c, 0xc6, 0x43, 0x1c, 0x0f, 0x93, 0x89, 0xdb, 0xc3, 0x92, 0x5d, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x03, 0x14, 0x72, 0x70, 0xaa, 0x01, 0x00, 0x00,
}
