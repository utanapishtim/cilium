// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_endpoint "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/endpoint"
import envoy_type "github.com/cilium/cilium/pkg/envoy/envoy/type"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Each route from RDS will map to a single cluster or traffic split across
// clusters using weights expressed in the RDS WeightedCluster.
//
// With EDS, each cluster is treated independently from a LB perspective, with
// LB taking place between the Localities within a cluster and at a finer
// granularity between the hosts within a locality. For a given cluster, the
// effective weight of a host is its load_balancing_weight multiplied by the
// load_balancing_weight of its Locality.
type ClusterLoadAssignment struct {
	// Name of the cluster. This will be the :ref:`service_name
	// <envoy_api_field_Cluster.EdsClusterConfig.service_name>` value if specified
	// in the cluster :ref:`EdsClusterConfig
	// <envoy_api_msg_Cluster.EdsClusterConfig>`.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	// List of endpoints to load balance to.
	Endpoints []*envoy_api_v2_endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
	// Load balancing policy settings.
	Policy *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy" json:"policy,omitempty"`
}

func (m *ClusterLoadAssignment) Reset()                    { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string            { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()               {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*envoy_api_v2_endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Load balancing policy settings.
type ClusterLoadAssignment_Policy struct {
	// Action to trim the overall incoming traffic to protect the upstream
	// hosts. This action allows protection in case the hosts are unable to
	// recover from an outage, or unable to autoscale or unable to handle
	// incoming traffic volume for any reason.
	//
	// At the client each category is applied one after the other to generate
	// the 'actual' drop percentage on all outgoing traffic. For example:
	//
	// .. code-block:: json
	//
	//  { "drop_overloads": [
	//      { "category": "throttle", "drop_percentage": 60 }
	//      { "category": "lb", "drop_percentage": 50 }
	//  ]}
	//
	// The actual drop percentages applied to the traffic at the clients will be
	//    "throttle"_drop = 60%
	//    "lb"_drop = 20%  // 50% of the remaining 'actual' load, which is 40%.
	//    actual_outgoing_load = 20% // remaining after applying all categories.
	DropOverloads []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads" json:"drop_overloads,omitempty"`
}

func (m *ClusterLoadAssignment_Policy) Reset()                    { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string            { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()               {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	// Identifier for the policy specifying the drop.
	Category string `protobuf:"bytes,1,opt,name=category" json:"category,omitempty"`
	// Percentage of traffic that should be dropped for the category.
	DropPercentage *envoy_type.Percent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage" json:"drop_percentage,omitempty"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *envoy_type.Percent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EndpointDiscoveryService service

type EndpointDiscoveryServiceClient interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EndpointDiscoveryService service

type EndpointDiscoveryServiceServer interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x3d, 0x6f, 0x13, 0x41,
	0x10, 0xcd, 0x9e, 0x2d, 0x13, 0x6f, 0x4c, 0x82, 0x96, 0x8f, 0x9c, 0x4e, 0x56, 0x62, 0x59, 0x20,
	0x59, 0x11, 0xba, 0x43, 0xa6, 0x40, 0x8a, 0x68, 0x30, 0x81, 0x02, 0x59, 0xc1, 0xba, 0x34, 0x54,
	0x44, 0x9b, 0xbb, 0xd1, 0xb1, 0xd2, 0x79, 0x67, 0xd9, 0x5d, 0x9f, 0x74, 0x2d, 0x15, 0x3d, 0x7f,
	0x82, 0xdf, 0x40, 0x03, 0x25, 0x3d, 0x3d, 0x05, 0xd0, 0x20, 0xfe, 0x04, 0xf2, 0x7d, 0x61, 0xcb,
	0x20, 0x51, 0xa4, 0x9b, 0xdb, 0x37, 0xef, 0xe9, 0xbd, 0x99, 0x39, 0x7a, 0x0b, 0x64, 0x86, 0x79,
	0xc0, 0x95, 0x08, 0xb2, 0x71, 0x00, 0xb1, 0xf1, 0x95, 0x46, 0x8b, 0xac, 0x57, 0xbc, 0xfb, 0x5c,
	0x09, 0x3f, 0x1b, 0x7b, 0xfd, 0xb5, 0xae, 0x58, 0x98, 0x08, 0x33, 0xd0, 0x79, 0xd9, 0xeb, 0xdd,
	0x5e, 0xd7, 0x90, 0xb1, 0x42, 0x21, 0x6d, 0x53, 0x54, 0x5d, 0x6e, 0xd9, 0x65, 0x73, 0x05, 0x81,
	0x02, 0x1d, 0x41, 0x83, 0xf4, 0x13, 0xc4, 0x24, 0x85, 0x42, 0x80, 0x4b, 0x89, 0x96, 0x5b, 0x81,
	0xb2, 0x72, 0xe2, 0xed, 0x67, 0x3c, 0x15, 0x31, 0xb7, 0x10, 0xd4, 0x45, 0x05, 0xdc, 0x48, 0x30,
	0xc1, 0xa2, 0x0c, 0x96, 0x55, 0xf9, 0x3a, 0xfc, 0xd8, 0xa2, 0x37, 0x1f, 0xa7, 0x0b, 0x63, 0x41,
	0x4f, 0x91, 0xc7, 0x8f, 0x8c, 0x11, 0x89, 0x9c, 0x83, 0xb4, 0xec, 0x2e, 0xed, 0x45, 0x25, 0x70,
	0x2e, 0xf9, 0x1c, 0x5c, 0x32, 0x20, 0xa3, 0xee, 0xa4, 0xfb, 0xe1, 0xe7, 0xa7, 0x56, 0x5b, 0x3b,
	0x03, 0x12, 0xee, 0x54, 0xf0, 0x29, 0x9f, 0x03, 0x3b, 0xa5, 0xdd, 0x3a, 0x80, 0x71, 0x9d, 0x41,
	0x6b, 0xb4, 0x33, 0x3e, 0xf2, 0x57, 0x87, 0xe2, 0x37, 0xf9, 0xa6, 0x18, 0xf1, 0x54, 0xd8, 0x7c,
	0x7a, 0xf1, 0xa4, 0x66, 0x4c, 0xda, 0x9f, 0xbf, 0x1e, 0x6e, 0x85, 0x7f, 0x24, 0xd8, 0x84, 0x76,
	0x14, 0xa6, 0x22, 0xca, 0xdd, 0xf6, 0x80, 0x6c, 0x8a, 0xfd, 0xd5, 0xb2, 0x3f, 0x2b, 0x18, 0x61,
	0xc5, 0xf4, 0xbe, 0x13, 0xda, 0x29, 0x9f, 0xd8, 0x4b, 0xba, 0x1b, 0x6b, 0x54, 0xe7, 0xcb, 0x3d,
	0xa4, 0xc8, 0xe3, 0xda, 0xe3, 0x83, 0xff, 0x97, 0xf5, 0x4f, 0x34, 0xaa, 0xe7, 0x15, 0x3f, 0xbc,
	0x1a, 0xaf, 0x7c, 0x19, 0xcf, 0xd0, 0xde, 0x2a, 0xcc, 0xee, 0xd0, 0xed, 0x88, 0x5b, 0x48, 0x50,
	0xe7, 0x9b, 0x83, 0x6b, 0x20, 0xf6, 0x90, 0xee, 0x15, 0xb6, 0xaa, 0x05, 0xf3, 0x04, 0x5c, 0xa7,
	0x88, 0x7b, 0xbd, 0xf2, 0xb5, 0x5c, 0xbf, 0x3f, 0x2b, 0xd1, 0xb0, 0x88, 0x30, 0x6b, 0x5a, 0x9f,
	0xb5, 0xb7, 0xc9, 0x35, 0x67, 0xfc, 0x8b, 0x50, 0xb7, 0x1e, 0xe4, 0x49, 0x7d, 0x6a, 0x67, 0xa0,
	0x33, 0x11, 0x01, 0x7b, 0x41, 0xf7, 0xce, 0xac, 0x06, 0x3e, 0x6f, 0x46, 0xcd, 0x0e, 0xd6, 0x23,
	0x37, 0x94, 0x10, 0x5e, 0x2f, 0xc0, 0x58, 0xef, 0xf0, 0x9f, 0xb8, 0x51, 0x28, 0x0d, 0x0c, 0xb7,
	0x46, 0xe4, 0x1e, 0x61, 0x0b, 0xba, 0xfb, 0x14, 0x6c, 0xf4, 0xea, 0x12, 0x85, 0x87, 0x6f, 0xbe,
	0xfc, 0x78, 0xe7, 0xf4, 0x87, 0xfb, 0x6b, 0x7f, 0xcd, 0x71, 0x73, 0x14, 0xc7, 0xe4, 0x68, 0x72,
	0xe5, 0xfd, 0xb7, 0x03, 0xf2, 0x96, 0x90, 0x8b, 0x4e, 0x71, 0xbf, 0xf7, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0xed, 0x97, 0x97, 0x92, 0x03, 0x00, 0x00,
}
