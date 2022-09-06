// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: collection_agent.proto

package profiling

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CollectionAgentClient is the client API for CollectionAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionAgentClient interface {
	// Sends once initial information about the host
	SaveHostInfo(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*empty.Empty, error)
	// For a list of traces, increments their counts by provided values
	AddCountsForTraces(ctx context.Context, in *AddCountsForTracesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Relates a list of stack traces with their hashes
	SetFramesForTraces(ctx context.Context, in *SetFramesForTracesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Adds metadata associated with a frame
	AddFrameMetadata(ctx context.Context, in *AddFrameMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Adds metadata to a list of executable files
	AddExecutableMetadata(ctx context.Context, in *AddExecutableMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Sends list of host metrics
	AddMetrics(ctx context.Context, in *Metrics, opts ...grpc.CallOption) (*empty.Empty, error)
	// Reports host metadata
	ReportHostMetadata(ctx context.Context, in *HostMetadata, opts ...grpc.CallOption) (*empty.Empty, error)
	// Adds fallback symbols for a set of frames, which can be used when full symbolization isn't
	// possible.
	AddFallbackSymbols(ctx context.Context, in *AddFallbackSymbolsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Heartbeat message from HA to CA on which we apply hosts per project throttling.
	Heartbeat(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetSymbolsPackageUploadURL returns an URL for uploading symbols of a given file.
	// The returned URL is a pre-authenticated HTTP PUT. It can be used to perform an upload of a package that contains
	// the associated symbols.
	// If no symbols are needed (because an upload URL was recently requested, an upload completed, or because symbols
	// are already known), the returned URL will be empty.
	// PackageUploadComplete must be called after the upload is completed, otherwise GetSymbolsPackageUploadURL may
	// again return an upload URL after some period of time (signalling a host agent to upload again).
	GetSymbolsPackageUploadURL(ctx context.Context, in *GetSymbolsPackageUploadURLRequest, opts ...grpc.CallOption) (*GetSymbolsPackageUploadURLResponse, error)
	// PackageUploadComplete signals the successful upload of a package that was previously requested through
	// GetSymbolsPackageUploadURL.
	// Invoking this signals that GetSymbolsPackageUploadURL must not return further upload URLs for the file(s)
	// associated with the upload.
	// API design note: this API exists so that the host-agents will not retry uploading the same package at regular
	// intervals when the indexing of the package doesn't actually yield symbols (or if the indexing of the uploaded
	// package doesn't happen for some unexpected reason).
	PackageUploadComplete(ctx context.Context, in *PackageUploadCompleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type collectionAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionAgentClient(cc grpc.ClientConnInterface) CollectionAgentClient {
	return &collectionAgentClient{cc}
}

func (c *collectionAgentClient) SaveHostInfo(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/SaveHostInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) AddCountsForTraces(ctx context.Context, in *AddCountsForTracesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/AddCountsForTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) SetFramesForTraces(ctx context.Context, in *SetFramesForTracesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/SetFramesForTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) AddFrameMetadata(ctx context.Context, in *AddFrameMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/AddFrameMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) AddExecutableMetadata(ctx context.Context, in *AddExecutableMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/AddExecutableMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) AddMetrics(ctx context.Context, in *Metrics, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/AddMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) ReportHostMetadata(ctx context.Context, in *HostMetadata, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/ReportHostMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) AddFallbackSymbols(ctx context.Context, in *AddFallbackSymbolsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/AddFallbackSymbols", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) Heartbeat(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) GetSymbolsPackageUploadURL(ctx context.Context, in *GetSymbolsPackageUploadURLRequest, opts ...grpc.CallOption) (*GetSymbolsPackageUploadURLResponse, error) {
	out := new(GetSymbolsPackageUploadURLResponse)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/GetSymbolsPackageUploadURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionAgentClient) PackageUploadComplete(ctx context.Context, in *PackageUploadCompleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/collectionagent.CollectionAgent/PackageUploadComplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionAgentServer is the server API for CollectionAgent service.
// All implementations must embed UnimplementedCollectionAgentServer
// for forward compatibility
type CollectionAgentServer interface {
	// Sends once initial information about the host
	SaveHostInfo(context.Context, *HostInfo) (*empty.Empty, error)
	// For a list of traces, increments their counts by provided values
	AddCountsForTraces(context.Context, *AddCountsForTracesRequest) (*empty.Empty, error)
	// Relates a list of stack traces with their hashes
	SetFramesForTraces(context.Context, *SetFramesForTracesRequest) (*empty.Empty, error)
	// Adds metadata associated with a frame
	AddFrameMetadata(context.Context, *AddFrameMetadataRequest) (*empty.Empty, error)
	// Adds metadata to a list of executable files
	AddExecutableMetadata(context.Context, *AddExecutableMetadataRequest) (*empty.Empty, error)
	// Sends list of host metrics
	AddMetrics(context.Context, *Metrics) (*empty.Empty, error)
	// Reports host metadata
	ReportHostMetadata(context.Context, *HostMetadata) (*empty.Empty, error)
	// Adds fallback symbols for a set of frames, which can be used when full symbolization isn't
	// possible.
	AddFallbackSymbols(context.Context, *AddFallbackSymbolsRequest) (*empty.Empty, error)
	// Heartbeat message from HA to CA on which we apply hosts per project throttling.
	Heartbeat(context.Context, *empty.Empty) (*empty.Empty, error)
	// GetSymbolsPackageUploadURL returns an URL for uploading symbols of a given file.
	// The returned URL is a pre-authenticated HTTP PUT. It can be used to perform an upload of a package that contains
	// the associated symbols.
	// If no symbols are needed (because an upload URL was recently requested, an upload completed, or because symbols
	// are already known), the returned URL will be empty.
	// PackageUploadComplete must be called after the upload is completed, otherwise GetSymbolsPackageUploadURL may
	// again return an upload URL after some period of time (signalling a host agent to upload again).
	GetSymbolsPackageUploadURL(context.Context, *GetSymbolsPackageUploadURLRequest) (*GetSymbolsPackageUploadURLResponse, error)
	// PackageUploadComplete signals the successful upload of a package that was previously requested through
	// GetSymbolsPackageUploadURL.
	// Invoking this signals that GetSymbolsPackageUploadURL must not return further upload URLs for the file(s)
	// associated with the upload.
	// API design note: this API exists so that the host-agents will not retry uploading the same package at regular
	// intervals when the indexing of the package doesn't actually yield symbols (or if the indexing of the uploaded
	// package doesn't happen for some unexpected reason).
	PackageUploadComplete(context.Context, *PackageUploadCompleteRequest) (*empty.Empty, error)
	mustEmbedUnimplementedCollectionAgentServer()
}

// UnimplementedCollectionAgentServer must be embedded to have forward compatible implementations.
type UnimplementedCollectionAgentServer struct {
}

func (UnimplementedCollectionAgentServer) SaveHostInfo(context.Context, *HostInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveHostInfo not implemented")
}
func (UnimplementedCollectionAgentServer) AddCountsForTraces(context.Context, *AddCountsForTracesRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCountsForTraces not implemented")
}
func (UnimplementedCollectionAgentServer) SetFramesForTraces(context.Context, *SetFramesForTracesRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFramesForTraces not implemented")
}
func (UnimplementedCollectionAgentServer) AddFrameMetadata(context.Context, *AddFrameMetadataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFrameMetadata not implemented")
}
func (UnimplementedCollectionAgentServer) AddExecutableMetadata(context.Context, *AddExecutableMetadataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExecutableMetadata not implemented")
}
func (UnimplementedCollectionAgentServer) AddMetrics(context.Context, *Metrics) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetrics not implemented")
}
func (UnimplementedCollectionAgentServer) ReportHostMetadata(context.Context, *HostMetadata) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportHostMetadata not implemented")
}
func (UnimplementedCollectionAgentServer) AddFallbackSymbols(context.Context, *AddFallbackSymbolsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFallbackSymbols not implemented")
}
func (UnimplementedCollectionAgentServer) Heartbeat(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedCollectionAgentServer) GetSymbolsPackageUploadURL(context.Context, *GetSymbolsPackageUploadURLRequest) (*GetSymbolsPackageUploadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSymbolsPackageUploadURL not implemented")
}
func (UnimplementedCollectionAgentServer) PackageUploadComplete(context.Context, *PackageUploadCompleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PackageUploadComplete not implemented")
}
func (UnimplementedCollectionAgentServer) mustEmbedUnimplementedCollectionAgentServer() {}

// UnsafeCollectionAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionAgentServer will
// result in compilation errors.
type UnsafeCollectionAgentServer interface {
	mustEmbedUnimplementedCollectionAgentServer()
}

func RegisterCollectionAgentServer(s grpc.ServiceRegistrar, srv CollectionAgentServer) {
	s.RegisterService(&CollectionAgent_ServiceDesc, srv)
}

func _CollectionAgent_SaveHostInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).SaveHostInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/SaveHostInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).SaveHostInfo(ctx, req.(*HostInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_AddCountsForTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCountsForTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).AddCountsForTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/AddCountsForTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).AddCountsForTraces(ctx, req.(*AddCountsForTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_SetFramesForTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFramesForTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).SetFramesForTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/SetFramesForTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).SetFramesForTraces(ctx, req.(*SetFramesForTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_AddFrameMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFrameMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).AddFrameMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/AddFrameMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).AddFrameMetadata(ctx, req.(*AddFrameMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_AddExecutableMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExecutableMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).AddExecutableMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/AddExecutableMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).AddExecutableMetadata(ctx, req.(*AddExecutableMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_AddMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).AddMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/AddMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).AddMetrics(ctx, req.(*Metrics))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_ReportHostMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).ReportHostMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/ReportHostMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).ReportHostMetadata(ctx, req.(*HostMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_AddFallbackSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFallbackSymbolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).AddFallbackSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/AddFallbackSymbols",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).AddFallbackSymbols(ctx, req.(*AddFallbackSymbolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).Heartbeat(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_GetSymbolsPackageUploadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSymbolsPackageUploadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).GetSymbolsPackageUploadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/GetSymbolsPackageUploadURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).GetSymbolsPackageUploadURL(ctx, req.(*GetSymbolsPackageUploadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionAgent_PackageUploadComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageUploadCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionAgentServer).PackageUploadComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collectionagent.CollectionAgent/PackageUploadComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionAgentServer).PackageUploadComplete(ctx, req.(*PackageUploadCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectionAgent_ServiceDesc is the grpc.ServiceDesc for CollectionAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectionAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collectionagent.CollectionAgent",
	HandlerType: (*CollectionAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveHostInfo",
			Handler:    _CollectionAgent_SaveHostInfo_Handler,
		},
		{
			MethodName: "AddCountsForTraces",
			Handler:    _CollectionAgent_AddCountsForTraces_Handler,
		},
		{
			MethodName: "SetFramesForTraces",
			Handler:    _CollectionAgent_SetFramesForTraces_Handler,
		},
		{
			MethodName: "AddFrameMetadata",
			Handler:    _CollectionAgent_AddFrameMetadata_Handler,
		},
		{
			MethodName: "AddExecutableMetadata",
			Handler:    _CollectionAgent_AddExecutableMetadata_Handler,
		},
		{
			MethodName: "AddMetrics",
			Handler:    _CollectionAgent_AddMetrics_Handler,
		},
		{
			MethodName: "ReportHostMetadata",
			Handler:    _CollectionAgent_ReportHostMetadata_Handler,
		},
		{
			MethodName: "AddFallbackSymbols",
			Handler:    _CollectionAgent_AddFallbackSymbols_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _CollectionAgent_Heartbeat_Handler,
		},
		{
			MethodName: "GetSymbolsPackageUploadURL",
			Handler:    _CollectionAgent_GetSymbolsPackageUploadURL_Handler,
		},
		{
			MethodName: "PackageUploadComplete",
			Handler:    _CollectionAgent_PackageUploadComplete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collection_agent.proto",
}
