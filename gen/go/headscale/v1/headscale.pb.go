// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: headscale/v1/headscale.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_headscale_v1_headscale_proto protoreflect.FileDescriptor

var file_headscale_v1_headscale_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb4, 0x19, 0x0a, 0x10, 0x48, 0x65,
	0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x12, 0x68, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22,
	0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x82, 0x01,
	0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x2f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x7b, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x6c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x12, 0x62, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65,
	0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x12, 0x87, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x2f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x12, 0x7a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x12, 0x7d, 0x0a,
	0x0f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x66, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12,
	0x1c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x74, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x6f, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x7b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x76, 0x0a, 0x0a, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x7b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22, 0x28, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x7b, 0x6e, 0x65,
	0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x62, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x6e, 0x0a, 0x08, 0x4d,
	0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x1b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x64, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x7c, 0x0a, 0x0b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x20, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x80, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22,
	0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f,
	0x7b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x7f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x7b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x75, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a,
	0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f,
	0x7b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x70, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x12, 0x77, 0x0a, 0x0c,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65,
	0x79, 0x12, 0x76, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x2a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79,
	0x2f, 0x7b, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x7d, 0x12, 0x58, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x43, 0x4c, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x6c, 0x12, 0x5b, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x12, 0x1b, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x41, 0x43, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x6c,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x75, 0x61, 0x6e, 0x66, 0x6f, 0x6e, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_headscale_v1_headscale_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),           // 0: headscale.v1.GetUserRequest
	(*CreateUserRequest)(nil),        // 1: headscale.v1.CreateUserRequest
	(*RenameUserRequest)(nil),        // 2: headscale.v1.RenameUserRequest
	(*DeleteUserRequest)(nil),        // 3: headscale.v1.DeleteUserRequest
	(*ListUsersRequest)(nil),         // 4: headscale.v1.ListUsersRequest
	(*CreatePreAuthKeyRequest)(nil),  // 5: headscale.v1.CreatePreAuthKeyRequest
	(*ExpirePreAuthKeyRequest)(nil),  // 6: headscale.v1.ExpirePreAuthKeyRequest
	(*ListPreAuthKeysRequest)(nil),   // 7: headscale.v1.ListPreAuthKeysRequest
	(*DebugCreateNodeRequest)(nil),   // 8: headscale.v1.DebugCreateNodeRequest
	(*GetNodeRequest)(nil),           // 9: headscale.v1.GetNodeRequest
	(*SetTagsRequest)(nil),           // 10: headscale.v1.SetTagsRequest
	(*RegisterNodeRequest)(nil),      // 11: headscale.v1.RegisterNodeRequest
	(*DeleteNodeRequest)(nil),        // 12: headscale.v1.DeleteNodeRequest
	(*ExpireNodeRequest)(nil),        // 13: headscale.v1.ExpireNodeRequest
	(*RenameNodeRequest)(nil),        // 14: headscale.v1.RenameNodeRequest
	(*ListNodesRequest)(nil),         // 15: headscale.v1.ListNodesRequest
	(*MoveNodeRequest)(nil),          // 16: headscale.v1.MoveNodeRequest
	(*GetRoutesRequest)(nil),         // 17: headscale.v1.GetRoutesRequest
	(*EnableRouteRequest)(nil),       // 18: headscale.v1.EnableRouteRequest
	(*DisableRouteRequest)(nil),      // 19: headscale.v1.DisableRouteRequest
	(*GetNodeRoutesRequest)(nil),     // 20: headscale.v1.GetNodeRoutesRequest
	(*DeleteRouteRequest)(nil),       // 21: headscale.v1.DeleteRouteRequest
	(*CreateApiKeyRequest)(nil),      // 22: headscale.v1.CreateApiKeyRequest
	(*ExpireApiKeyRequest)(nil),      // 23: headscale.v1.ExpireApiKeyRequest
	(*ListApiKeysRequest)(nil),       // 24: headscale.v1.ListApiKeysRequest
	(*DeleteApiKeyRequest)(nil),      // 25: headscale.v1.DeleteApiKeyRequest
	(*GetACLRequest)(nil),            // 26: headscale.v1.GetACLRequest
	(*SetACLRequest)(nil),            // 27: headscale.v1.SetACLRequest
	(*GetUserResponse)(nil),          // 28: headscale.v1.GetUserResponse
	(*CreateUserResponse)(nil),       // 29: headscale.v1.CreateUserResponse
	(*RenameUserResponse)(nil),       // 30: headscale.v1.RenameUserResponse
	(*DeleteUserResponse)(nil),       // 31: headscale.v1.DeleteUserResponse
	(*ListUsersResponse)(nil),        // 32: headscale.v1.ListUsersResponse
	(*CreatePreAuthKeyResponse)(nil), // 33: headscale.v1.CreatePreAuthKeyResponse
	(*ExpirePreAuthKeyResponse)(nil), // 34: headscale.v1.ExpirePreAuthKeyResponse
	(*ListPreAuthKeysResponse)(nil),  // 35: headscale.v1.ListPreAuthKeysResponse
	(*DebugCreateNodeResponse)(nil),  // 36: headscale.v1.DebugCreateNodeResponse
	(*GetNodeResponse)(nil),          // 37: headscale.v1.GetNodeResponse
	(*SetTagsResponse)(nil),          // 38: headscale.v1.SetTagsResponse
	(*RegisterNodeResponse)(nil),     // 39: headscale.v1.RegisterNodeResponse
	(*DeleteNodeResponse)(nil),       // 40: headscale.v1.DeleteNodeResponse
	(*ExpireNodeResponse)(nil),       // 41: headscale.v1.ExpireNodeResponse
	(*RenameNodeResponse)(nil),       // 42: headscale.v1.RenameNodeResponse
	(*ListNodesResponse)(nil),        // 43: headscale.v1.ListNodesResponse
	(*MoveNodeResponse)(nil),         // 44: headscale.v1.MoveNodeResponse
	(*GetRoutesResponse)(nil),        // 45: headscale.v1.GetRoutesResponse
	(*EnableRouteResponse)(nil),      // 46: headscale.v1.EnableRouteResponse
	(*DisableRouteResponse)(nil),     // 47: headscale.v1.DisableRouteResponse
	(*GetNodeRoutesResponse)(nil),    // 48: headscale.v1.GetNodeRoutesResponse
	(*DeleteRouteResponse)(nil),      // 49: headscale.v1.DeleteRouteResponse
	(*CreateApiKeyResponse)(nil),     // 50: headscale.v1.CreateApiKeyResponse
	(*ExpireApiKeyResponse)(nil),     // 51: headscale.v1.ExpireApiKeyResponse
	(*ListApiKeysResponse)(nil),      // 52: headscale.v1.ListApiKeysResponse
	(*DeleteApiKeyResponse)(nil),     // 53: headscale.v1.DeleteApiKeyResponse
	(*GetACLResponse)(nil),           // 54: headscale.v1.GetACLResponse
	(*SetACLResponse)(nil),           // 55: headscale.v1.SetACLResponse
}
var file_headscale_v1_headscale_proto_depIdxs = []int32{
	0,  // 0: headscale.v1.HeadscaleService.GetUser:input_type -> headscale.v1.GetUserRequest
	1,  // 1: headscale.v1.HeadscaleService.CreateUser:input_type -> headscale.v1.CreateUserRequest
	2,  // 2: headscale.v1.HeadscaleService.RenameUser:input_type -> headscale.v1.RenameUserRequest
	3,  // 3: headscale.v1.HeadscaleService.DeleteUser:input_type -> headscale.v1.DeleteUserRequest
	4,  // 4: headscale.v1.HeadscaleService.ListUsers:input_type -> headscale.v1.ListUsersRequest
	5,  // 5: headscale.v1.HeadscaleService.CreatePreAuthKey:input_type -> headscale.v1.CreatePreAuthKeyRequest
	6,  // 6: headscale.v1.HeadscaleService.ExpirePreAuthKey:input_type -> headscale.v1.ExpirePreAuthKeyRequest
	7,  // 7: headscale.v1.HeadscaleService.ListPreAuthKeys:input_type -> headscale.v1.ListPreAuthKeysRequest
	8,  // 8: headscale.v1.HeadscaleService.DebugCreateNode:input_type -> headscale.v1.DebugCreateNodeRequest
	9,  // 9: headscale.v1.HeadscaleService.GetNode:input_type -> headscale.v1.GetNodeRequest
	10, // 10: headscale.v1.HeadscaleService.SetTags:input_type -> headscale.v1.SetTagsRequest
	11, // 11: headscale.v1.HeadscaleService.RegisterNode:input_type -> headscale.v1.RegisterNodeRequest
	12, // 12: headscale.v1.HeadscaleService.DeleteNode:input_type -> headscale.v1.DeleteNodeRequest
	13, // 13: headscale.v1.HeadscaleService.ExpireNode:input_type -> headscale.v1.ExpireNodeRequest
	14, // 14: headscale.v1.HeadscaleService.RenameNode:input_type -> headscale.v1.RenameNodeRequest
	15, // 15: headscale.v1.HeadscaleService.ListNodes:input_type -> headscale.v1.ListNodesRequest
	16, // 16: headscale.v1.HeadscaleService.MoveNode:input_type -> headscale.v1.MoveNodeRequest
	17, // 17: headscale.v1.HeadscaleService.GetRoutes:input_type -> headscale.v1.GetRoutesRequest
	18, // 18: headscale.v1.HeadscaleService.EnableRoute:input_type -> headscale.v1.EnableRouteRequest
	19, // 19: headscale.v1.HeadscaleService.DisableRoute:input_type -> headscale.v1.DisableRouteRequest
	20, // 20: headscale.v1.HeadscaleService.GetNodeRoutes:input_type -> headscale.v1.GetNodeRoutesRequest
	21, // 21: headscale.v1.HeadscaleService.DeleteRoute:input_type -> headscale.v1.DeleteRouteRequest
	22, // 22: headscale.v1.HeadscaleService.CreateApiKey:input_type -> headscale.v1.CreateApiKeyRequest
	23, // 23: headscale.v1.HeadscaleService.ExpireApiKey:input_type -> headscale.v1.ExpireApiKeyRequest
	24, // 24: headscale.v1.HeadscaleService.ListApiKeys:input_type -> headscale.v1.ListApiKeysRequest
	25, // 25: headscale.v1.HeadscaleService.DeleteApiKey:input_type -> headscale.v1.DeleteApiKeyRequest
	26, // 26: headscale.v1.HeadscaleService.GetACL:input_type -> headscale.v1.GetACLRequest
	27, // 27: headscale.v1.HeadscaleService.SetACL:input_type -> headscale.v1.SetACLRequest
	28, // 28: headscale.v1.HeadscaleService.GetUser:output_type -> headscale.v1.GetUserResponse
	29, // 29: headscale.v1.HeadscaleService.CreateUser:output_type -> headscale.v1.CreateUserResponse
	30, // 30: headscale.v1.HeadscaleService.RenameUser:output_type -> headscale.v1.RenameUserResponse
	31, // 31: headscale.v1.HeadscaleService.DeleteUser:output_type -> headscale.v1.DeleteUserResponse
	32, // 32: headscale.v1.HeadscaleService.ListUsers:output_type -> headscale.v1.ListUsersResponse
	33, // 33: headscale.v1.HeadscaleService.CreatePreAuthKey:output_type -> headscale.v1.CreatePreAuthKeyResponse
	34, // 34: headscale.v1.HeadscaleService.ExpirePreAuthKey:output_type -> headscale.v1.ExpirePreAuthKeyResponse
	35, // 35: headscale.v1.HeadscaleService.ListPreAuthKeys:output_type -> headscale.v1.ListPreAuthKeysResponse
	36, // 36: headscale.v1.HeadscaleService.DebugCreateNode:output_type -> headscale.v1.DebugCreateNodeResponse
	37, // 37: headscale.v1.HeadscaleService.GetNode:output_type -> headscale.v1.GetNodeResponse
	38, // 38: headscale.v1.HeadscaleService.SetTags:output_type -> headscale.v1.SetTagsResponse
	39, // 39: headscale.v1.HeadscaleService.RegisterNode:output_type -> headscale.v1.RegisterNodeResponse
	40, // 40: headscale.v1.HeadscaleService.DeleteNode:output_type -> headscale.v1.DeleteNodeResponse
	41, // 41: headscale.v1.HeadscaleService.ExpireNode:output_type -> headscale.v1.ExpireNodeResponse
	42, // 42: headscale.v1.HeadscaleService.RenameNode:output_type -> headscale.v1.RenameNodeResponse
	43, // 43: headscale.v1.HeadscaleService.ListNodes:output_type -> headscale.v1.ListNodesResponse
	44, // 44: headscale.v1.HeadscaleService.MoveNode:output_type -> headscale.v1.MoveNodeResponse
	45, // 45: headscale.v1.HeadscaleService.GetRoutes:output_type -> headscale.v1.GetRoutesResponse
	46, // 46: headscale.v1.HeadscaleService.EnableRoute:output_type -> headscale.v1.EnableRouteResponse
	47, // 47: headscale.v1.HeadscaleService.DisableRoute:output_type -> headscale.v1.DisableRouteResponse
	48, // 48: headscale.v1.HeadscaleService.GetNodeRoutes:output_type -> headscale.v1.GetNodeRoutesResponse
	49, // 49: headscale.v1.HeadscaleService.DeleteRoute:output_type -> headscale.v1.DeleteRouteResponse
	50, // 50: headscale.v1.HeadscaleService.CreateApiKey:output_type -> headscale.v1.CreateApiKeyResponse
	51, // 51: headscale.v1.HeadscaleService.ExpireApiKey:output_type -> headscale.v1.ExpireApiKeyResponse
	52, // 52: headscale.v1.HeadscaleService.ListApiKeys:output_type -> headscale.v1.ListApiKeysResponse
	53, // 53: headscale.v1.HeadscaleService.DeleteApiKey:output_type -> headscale.v1.DeleteApiKeyResponse
	54, // 54: headscale.v1.HeadscaleService.GetACL:output_type -> headscale.v1.GetACLResponse
	55, // 55: headscale.v1.HeadscaleService.SetACL:output_type -> headscale.v1.SetACLResponse
	28, // [28:56] is the sub-list for method output_type
	0,  // [0:28] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_headscale_v1_headscale_proto_init() }
func file_headscale_v1_headscale_proto_init() {
	if File_headscale_v1_headscale_proto != nil {
		return
	}
	file_headscale_v1_user_proto_init()
	file_headscale_v1_preauthkey_proto_init()
	file_headscale_v1_node_proto_init()
	file_headscale_v1_routes_proto_init()
	file_headscale_v1_apikey_proto_init()
	file_headscale_v1_acl_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_headscale_v1_headscale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_headscale_v1_headscale_proto_goTypes,
		DependencyIndexes: file_headscale_v1_headscale_proto_depIdxs,
	}.Build()
	File_headscale_v1_headscale_proto = out.File
	file_headscale_v1_headscale_proto_rawDesc = nil
	file_headscale_v1_headscale_proto_goTypes = nil
	file_headscale_v1_headscale_proto_depIdxs = nil
}
