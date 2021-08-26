// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/health/v3/hds.proto

package envoy_service_health_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Capability_Protocol int32

const (
	Capability_HTTP  Capability_Protocol = 0
	Capability_TCP   Capability_Protocol = 1
	Capability_REDIS Capability_Protocol = 2
)

var Capability_Protocol_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "REDIS",
}

var Capability_Protocol_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"REDIS": 2,
}

func (x Capability_Protocol) String() string {
	return proto.EnumName(Capability_Protocol_name, int32(x))
}

func (Capability_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{0, 0}
}

type Capability struct {
	HealthCheckProtocols []Capability_Protocol `protobuf:"varint,1,rep,packed,name=health_check_protocols,json=healthCheckProtocols,proto3,enum=envoy.service.health.v3.Capability_Protocol" json:"health_check_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Capability) Reset()         { *m = Capability{} }
func (m *Capability) String() string { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()    {}
func (*Capability) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{0}
}

func (m *Capability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capability.Unmarshal(m, b)
}
func (m *Capability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capability.Marshal(b, m, deterministic)
}
func (m *Capability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capability.Merge(m, src)
}
func (m *Capability) XXX_Size() int {
	return xxx_messageInfo_Capability.Size(m)
}
func (m *Capability) XXX_DiscardUnknown() {
	xxx_messageInfo_Capability.DiscardUnknown(m)
}

var xxx_messageInfo_Capability proto.InternalMessageInfo

func (m *Capability) GetHealthCheckProtocols() []Capability_Protocol {
	if m != nil {
		return m.HealthCheckProtocols
	}
	return nil
}

type HealthCheckRequest struct {
	Node                 *v3.Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Capability           *Capability `protobuf:"bytes,2,opt,name=capability,proto3" json:"capability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{1}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetNode() *v3.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *HealthCheckRequest) GetCapability() *Capability {
	if m != nil {
		return m.Capability
	}
	return nil
}

type EndpointHealth struct {
	Endpoint             *v31.Endpoint   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	HealthStatus         v3.HealthStatus `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.config.core.v3.HealthStatus" json:"health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EndpointHealth) Reset()         { *m = EndpointHealth{} }
func (m *EndpointHealth) String() string { return proto.CompactTextString(m) }
func (*EndpointHealth) ProtoMessage()    {}
func (*EndpointHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{2}
}

func (m *EndpointHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealth.Unmarshal(m, b)
}
func (m *EndpointHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealth.Marshal(b, m, deterministic)
}
func (m *EndpointHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealth.Merge(m, src)
}
func (m *EndpointHealth) XXX_Size() int {
	return xxx_messageInfo_EndpointHealth.Size(m)
}
func (m *EndpointHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealth.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealth proto.InternalMessageInfo

func (m *EndpointHealth) GetEndpoint() *v31.Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *EndpointHealth) GetHealthStatus() v3.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return v3.HealthStatus_UNKNOWN
}

type EndpointHealthResponse struct {
	EndpointsHealth      []*EndpointHealth `protobuf:"bytes,1,rep,name=endpoints_health,json=endpointsHealth,proto3" json:"endpoints_health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EndpointHealthResponse) Reset()         { *m = EndpointHealthResponse{} }
func (m *EndpointHealthResponse) String() string { return proto.CompactTextString(m) }
func (*EndpointHealthResponse) ProtoMessage()    {}
func (*EndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{3}
}

func (m *EndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealthResponse.Unmarshal(m, b)
}
func (m *EndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealthResponse.Marshal(b, m, deterministic)
}
func (m *EndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealthResponse.Merge(m, src)
}
func (m *EndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_EndpointHealthResponse.Size(m)
}
func (m *EndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealthResponse proto.InternalMessageInfo

func (m *EndpointHealthResponse) GetEndpointsHealth() []*EndpointHealth {
	if m != nil {
		return m.EndpointsHealth
	}
	return nil
}

type HealthCheckRequestOrEndpointHealthResponse struct {
	// Types that are valid to be assigned to RequestType:
	//	*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest
	//	*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse
	RequestType          isHealthCheckRequestOrEndpointHealthResponse_RequestType `protobuf_oneof:"request_type"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *HealthCheckRequestOrEndpointHealthResponse) Reset() {
	*m = HealthCheckRequestOrEndpointHealthResponse{}
}
func (m *HealthCheckRequestOrEndpointHealthResponse) String() string {
	return proto.CompactTextString(m)
}
func (*HealthCheckRequestOrEndpointHealthResponse) ProtoMessage() {}
func (*HealthCheckRequestOrEndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{4}
}

func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Unmarshal(m, b)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Merge(m, src)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Size(m)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse proto.InternalMessageInfo

type isHealthCheckRequestOrEndpointHealthResponse_RequestType interface {
	isHealthCheckRequestOrEndpointHealthResponse_RequestType()
}

type HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest struct {
	HealthCheckRequest *HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof"`
}

type HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse struct {
	EndpointHealthResponse *EndpointHealthResponse `protobuf:"bytes,2,opt,name=endpoint_health_response,json=endpointHealthResponse,proto3,oneof"`
}

func (*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetRequestType() isHealthCheckRequestOrEndpointHealthResponse_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetHealthCheckRequest() *HealthCheckRequest {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest); ok {
		return x.HealthCheckRequest
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetEndpointHealthResponse() *EndpointHealthResponse {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse); ok {
		return x.EndpointHealthResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheckRequestOrEndpointHealthResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest)(nil),
		(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse)(nil),
	}
}

type LocalityEndpoints struct {
	Locality             *v3.Locality    `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	Endpoints            []*v31.Endpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LocalityEndpoints) Reset()         { *m = LocalityEndpoints{} }
func (m *LocalityEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityEndpoints) ProtoMessage()    {}
func (*LocalityEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{5}
}

func (m *LocalityEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityEndpoints.Unmarshal(m, b)
}
func (m *LocalityEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityEndpoints.Merge(m, src)
}
func (m *LocalityEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityEndpoints.Size(m)
}
func (m *LocalityEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityEndpoints proto.InternalMessageInfo

func (m *LocalityEndpoints) GetLocality() *v3.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityEndpoints) GetEndpoints() []*v31.Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ClusterHealthCheck struct {
	ClusterName          string               `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	HealthChecks         []*v3.HealthCheck    `protobuf:"bytes,2,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	LocalityEndpoints    []*LocalityEndpoints `protobuf:"bytes,3,rep,name=locality_endpoints,json=localityEndpoints,proto3" json:"locality_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterHealthCheck) Reset()         { *m = ClusterHealthCheck{} }
func (m *ClusterHealthCheck) String() string { return proto.CompactTextString(m) }
func (*ClusterHealthCheck) ProtoMessage()    {}
func (*ClusterHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{6}
}

func (m *ClusterHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterHealthCheck.Unmarshal(m, b)
}
func (m *ClusterHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterHealthCheck.Marshal(b, m, deterministic)
}
func (m *ClusterHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHealthCheck.Merge(m, src)
}
func (m *ClusterHealthCheck) XXX_Size() int {
	return xxx_messageInfo_ClusterHealthCheck.Size(m)
}
func (m *ClusterHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHealthCheck proto.InternalMessageInfo

func (m *ClusterHealthCheck) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterHealthCheck) GetHealthChecks() []*v3.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *ClusterHealthCheck) GetLocalityEndpoints() []*LocalityEndpoints {
	if m != nil {
		return m.LocalityEndpoints
	}
	return nil
}

type HealthCheckSpecifier struct {
	ClusterHealthChecks  []*ClusterHealthCheck `protobuf:"bytes,1,rep,name=cluster_health_checks,json=clusterHealthChecks,proto3" json:"cluster_health_checks,omitempty"`
	Interval             *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HealthCheckSpecifier) Reset()         { *m = HealthCheckSpecifier{} }
func (m *HealthCheckSpecifier) String() string { return proto.CompactTextString(m) }
func (*HealthCheckSpecifier) ProtoMessage()    {}
func (*HealthCheckSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c56caa659fccb8, []int{7}
}

func (m *HealthCheckSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckSpecifier.Unmarshal(m, b)
}
func (m *HealthCheckSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckSpecifier.Marshal(b, m, deterministic)
}
func (m *HealthCheckSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckSpecifier.Merge(m, src)
}
func (m *HealthCheckSpecifier) XXX_Size() int {
	return xxx_messageInfo_HealthCheckSpecifier.Size(m)
}
func (m *HealthCheckSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckSpecifier proto.InternalMessageInfo

func (m *HealthCheckSpecifier) GetClusterHealthChecks() []*ClusterHealthCheck {
	if m != nil {
		return m.ClusterHealthChecks
	}
	return nil
}

func (m *HealthCheckSpecifier) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.service.health.v3.Capability_Protocol", Capability_Protocol_name, Capability_Protocol_value)
	proto.RegisterType((*Capability)(nil), "envoy.service.health.v3.Capability")
	proto.RegisterType((*HealthCheckRequest)(nil), "envoy.service.health.v3.HealthCheckRequest")
	proto.RegisterType((*EndpointHealth)(nil), "envoy.service.health.v3.EndpointHealth")
	proto.RegisterType((*EndpointHealthResponse)(nil), "envoy.service.health.v3.EndpointHealthResponse")
	proto.RegisterType((*HealthCheckRequestOrEndpointHealthResponse)(nil), "envoy.service.health.v3.HealthCheckRequestOrEndpointHealthResponse")
	proto.RegisterType((*LocalityEndpoints)(nil), "envoy.service.health.v3.LocalityEndpoints")
	proto.RegisterType((*ClusterHealthCheck)(nil), "envoy.service.health.v3.ClusterHealthCheck")
	proto.RegisterType((*HealthCheckSpecifier)(nil), "envoy.service.health.v3.HealthCheckSpecifier")
}

func init() { proto.RegisterFile("envoy/service/health/v3/hds.proto", fileDescriptor_28c56caa659fccb8) }

var fileDescriptor_28c56caa659fccb8 = []byte{
	// 908 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xee, 0xb8, 0xcb, 0x92, 0xbe, 0x96, 0x92, 0x0e, 0xa5, 0x84, 0x68, 0x29, 0x5b, 0xa3, 0x55,
	0x4b, 0xdb, 0xb5, 0x77, 0x13, 0x40, 0xc8, 0x87, 0x15, 0x6a, 0xda, 0xdd, 0x20, 0xad, 0x96, 0x6a,
	0xd2, 0x0b, 0x27, 0xcb, 0xb1, 0xa7, 0x8d, 0xb5, 0xae, 0xc7, 0x78, 0x1c, 0x8b, 0xdc, 0x10, 0xa7,
	0xbd, 0x20, 0x21, 0x71, 0xe3, 0x07, 0xf0, 0x07, 0xb8, 0x71, 0xe1, 0x04, 0x5a, 0x89, 0x13, 0xe2,
	0xc2, 0x89, 0x13, 0xe2, 0x77, 0x20, 0x7b, 0x66, 0x1c, 0xa7, 0xae, 0xa3, 0xf4, 0xb4, 0xb7, 0xcc,
	0x7b, 0xdf, 0x7b, 0xef, 0xfb, 0xde, 0xbc, 0x37, 0x0e, 0xec, 0xd0, 0x30, 0x65, 0x13, 0x93, 0xd3,
	0x38, 0xf5, 0x5d, 0x6a, 0x8e, 0xa8, 0x13, 0x24, 0x23, 0x33, 0xed, 0x9a, 0x23, 0x8f, 0x1b, 0x51,
	0xcc, 0x12, 0x86, 0xdf, 0xc9, 0x21, 0x86, 0x84, 0x18, 0x02, 0x62, 0xa4, 0xdd, 0xf6, 0xfb, 0x22,
	0xd6, 0x65, 0xe1, 0xb9, 0x7f, 0x61, 0xba, 0x2c, 0xa6, 0x59, 0xe0, 0xd0, 0xe1, 0x54, 0x44, 0xb6,
	0x77, 0xaf, 0x05, 0x88, 0x04, 0xb6, 0x3b, 0xa2, 0xee, 0x73, 0x09, 0xec, 0xcc, 0x00, 0x69, 0xe8,
	0x45, 0xcc, 0x0f, 0x93, 0x0c, 0xac, 0x7e, 0xdb, 0x2e, 0xbb, 0x8c, 0x58, 0x48, 0xc3, 0x44, 0xd2,
	0x6a, 0xdf, 0xb9, 0x60, 0xec, 0x22, 0xa0, 0xa6, 0x13, 0xf9, 0xa6, 0x13, 0x86, 0x2c, 0x71, 0x12,
	0x9f, 0x85, 0xca, 0xbb, 0x2d, 0xbd, 0xf9, 0x69, 0x38, 0x3e, 0x37, 0xbd, 0x71, 0x9c, 0x03, 0xa4,
	0xff, 0xbd, 0xb1, 0x17, 0x39, 0xe5, 0x38, 0x93, 0x27, 0x4e, 0x32, 0x56, 0xe1, 0x3b, 0x15, 0x77,
	0x4a, 0x63, 0xee, 0xb3, 0xd0, 0x0f, 0x2f, 0x04, 0x44, 0x7f, 0x89, 0x00, 0x7a, 0x4e, 0xe4, 0x0c,
	0xfd, 0xc0, 0x4f, 0x26, 0x78, 0x08, 0x5b, 0x65, 0x61, 0x76, 0x0e, 0x72, 0x59, 0xc0, 0x5b, 0xe8,
	0xee, 0xf2, 0xde, 0x7a, 0xe7, 0xd0, 0xa8, 0x69, 0xa3, 0x31, 0x4d, 0x62, 0x9c, 0xca, 0x20, 0xb2,
	0x29, 0xdc, 0xbd, 0x2c, 0x95, 0x32, 0x72, 0x7d, 0x0f, 0x1a, 0xea, 0x80, 0x1b, 0x70, 0xab, 0x7f,
	0x76, 0x76, 0xda, 0x5c, 0xc2, 0xaf, 0xc3, 0xf2, 0x59, 0xef, 0xb4, 0x89, 0xf0, 0x0a, 0xbc, 0x46,
	0x4e, 0x8e, 0x3f, 0x1f, 0x34, 0x35, 0xeb, 0xf0, 0xc7, 0xdf, 0x5e, 0x6c, 0xef, 0xc2, 0xbd, 0xd9,
	0x9a, 0x9e, 0xcf, 0x5d, 0x96, 0xd2, 0x78, 0x62, 0xa4, 0x9d, 0x52, 0x59, 0xfd, 0x57, 0x04, 0xb8,
	0x3f, 0x2d, 0x48, 0xe8, 0x57, 0x63, 0xca, 0x13, 0x6c, 0xc0, 0xad, 0x90, 0x79, 0xb4, 0x85, 0xee,
	0xa2, 0xbd, 0xd5, 0x4e, 0x5b, 0x0a, 0x10, 0x97, 0x64, 0x64, 0xb7, 0x99, 0xb1, 0x7f, 0xc6, 0x3c,
	0x4a, 0x72, 0x1c, 0xee, 0x01, 0xb8, 0x45, 0xd2, 0x96, 0x96, 0x47, 0x7d, 0xb0, 0x80, 0x6c, 0x52,
	0x0a, 0xb3, 0x3e, 0xca, 0x98, 0x9b, 0x70, 0x7f, 0x0e, 0xf3, 0x2a, 0x55, 0xfd, 0x0f, 0x04, 0xeb,
	0x27, 0x72, 0x54, 0x84, 0x1b, 0x3f, 0x82, 0x86, 0x1a, 0x1e, 0xa9, 0x40, 0x9f, 0x55, 0xa0, 0xbc,
	0x19, 0x19, 0x15, 0x4b, 0x8a, 0x18, 0xfc, 0x04, 0xde, 0x90, 0x17, 0x2a, 0x26, 0x23, 0x17, 0xb4,
	0x7e, 0x35, 0x89, 0x6a, 0x83, 0x28, 0x3a, 0xc8, 0x91, 0x64, 0x6d, 0x54, 0x3a, 0x59, 0x0f, 0x32,
	0x45, 0x07, 0xf0, 0xe1, 0x1c, 0x45, 0xb3, 0xd4, 0xf5, 0x9f, 0x10, 0x6c, 0xcd, 0x9a, 0x08, 0xe5,
	0x11, 0x0b, 0x39, 0xc5, 0x04, 0x9a, 0x8a, 0x21, 0xb7, 0x45, 0x99, 0x7c, 0xc0, 0x56, 0x3b, 0xbb,
	0xb5, 0x9d, 0xbe, 0x92, 0xea, 0xcd, 0x22, 0x81, 0x30, 0x58, 0x9f, 0x66, 0x04, 0xbb, 0xf0, 0x70,
	0x61, 0x82, 0x8a, 0x8d, 0xfe, 0xb7, 0x06, 0xfb, 0xd5, 0xdb, 0xf8, 0x22, 0xae, 0x21, 0x6f, 0xc3,
	0xe6, 0xcc, 0x8e, 0xc4, 0x02, 0x2f, 0xaf, 0xe7, 0xa0, 0x56, 0x40, 0xb5, 0x44, 0x7f, 0x89, 0xe0,
	0x51, 0x75, 0x62, 0x9f, 0x43, 0xab, 0x78, 0x30, 0x64, 0xa5, 0x58, 0x16, 0x97, 0xf3, 0x68, 0x2e,
	0xda, 0x25, 0x19, 0xd6, 0x5f, 0x22, 0x5b, 0xf4, 0x5a, 0x8f, 0xf5, 0x34, 0x6b, 0xdb, 0x13, 0x38,
	0xb9, 0xd1, 0xa4, 0xd6, 0xf5, 0xe6, 0x68, 0x1d, 0xd6, 0x64, 0x3b, 0xec, 0x64, 0x12, 0x51, 0xfd,
	0x77, 0x04, 0x1b, 0x4f, 0x99, 0xeb, 0x64, 0x4b, 0xa1, 0x42, 0x38, 0xb6, 0xa0, 0x11, 0x48, 0xa3,
	0xec, 0xda, 0xf6, 0xf5, 0xf3, 0xa8, 0x42, 0x49, 0x81, 0xc7, 0x9f, 0xc1, 0x4a, 0x71, 0xf3, 0x2d,
	0x2d, 0x9f, 0x99, 0x45, 0x36, 0x62, 0x1a, 0x64, 0x75, 0x33, 0xc5, 0x06, 0x1c, 0xce, 0x51, 0x5c,
	0xa1, 0xac, 0x7f, 0xaf, 0x01, 0xee, 0x05, 0x63, 0x9e, 0xd0, 0xb8, 0xd4, 0x0e, 0xbc, 0x03, 0x6b,
	0xae, 0xb0, 0xda, 0xa1, 0x73, 0x29, 0x1e, 0x99, 0x15, 0xb2, 0x2a, 0x6d, 0xcf, 0x9c, 0x4b, 0x8a,
	0x1f, 0x17, 0x1b, 0x98, 0x8f, 0x8b, 0x22, 0xbd, 0x33, 0x6f, 0x03, 0x45, 0xaf, 0xd7, 0x4a, 0xb3,
	0xc1, 0xf1, 0x97, 0x80, 0x55, 0x13, 0xec, 0x69, 0x07, 0x96, 0xf3, 0x64, 0xfb, 0xb5, 0xf3, 0x50,
	0x51, 0x42, 0x36, 0x82, 0xab, 0xa6, 0x45, 0x5e, 0xab, 0xaa, 0x76, 0xfd, 0x3f, 0x04, 0x9b, 0xa5,
	0xf3, 0x20, 0xa2, 0xae, 0x7f, 0xee, 0xd3, 0x18, 0xdb, 0xf0, 0xb6, 0x6a, 0xca, 0xac, 0x72, 0xb1,
	0xe2, 0xf5, 0x1b, 0x52, 0x2d, 0x42, 0xde, 0x72, 0x2b, 0x36, 0x8e, 0x3f, 0x86, 0x86, 0x1f, 0x26,
	0x34, 0x4e, 0x9d, 0x40, 0x2e, 0xc4, 0xbb, 0x86, 0xf8, 0x52, 0x1a, 0xea, 0x4b, 0x69, 0x1c, 0xcb,
	0x2f, 0x25, 0x29, 0xa0, 0xd6, 0x27, 0x99, 0xcc, 0x87, 0x60, 0x2e, 0x36, 0xea, 0x85, 0x9e, 0xce,
	0x3f, 0x1a, 0x6c, 0x09, 0xc7, 0xb1, 0xc2, 0x0e, 0x44, 0x30, 0xfe, 0x0e, 0xc1, 0xc6, 0x20, 0x89,
	0xa9, 0x73, 0x59, 0x9e, 0x8a, 0xde, 0x0d, 0xde, 0x80, 0xba, 0x55, 0x6a, 0xdf, 0x5f, 0x24, 0x49,
	0x41, 0x52, 0x5f, 0xda, 0x43, 0x0f, 0x10, 0xfe, 0x19, 0x41, 0xf3, 0x31, 0x4d, 0xdc, 0xd1, 0xab,
	0xa6, 0x73, 0xf0, 0xed, 0x5f, 0xff, 0xfe, 0xa0, 0xdd, 0xd1, 0xdb, 0xd9, 0xdf, 0x9f, 0xa2, 0xc5,
	0x56, 0x79, 0x1e, 0x72, 0xc4, 0xb2, 0x85, 0xf6, 0x8f, 0x1e, 0xfd, 0xf2, 0xcd, 0xcb, 0x3f, 0x6f,
	0x6b, 0x4d, 0x0d, 0xee, 0xf9, 0x4c, 0xd4, 0x89, 0x62, 0xf6, 0xf5, 0xa4, 0xae, 0xe4, 0x51, 0xa3,
	0xef, 0xf1, 0xfc, 0x3f, 0xc4, 0x29, 0x7a, 0x81, 0xd0, 0xf0, 0x76, 0x7e, 0xeb, 0xdd, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x37, 0xb9, 0xfe, 0x07, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthDiscoveryServiceClient is the client API for HealthDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthDiscoveryServiceClient interface {
	StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error)
	FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error)
}

type healthDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewHealthDiscoveryServiceClient(cc *grpc.ClientConn) HealthDiscoveryServiceClient {
	return &healthDiscoveryServiceClient{cc}
}

func (c *healthDiscoveryServiceClient) StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HealthDiscoveryService_serviceDesc.Streams[0], "/envoy.service.health.v3.HealthDiscoveryService/StreamHealthCheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthDiscoveryServiceStreamHealthCheckClient{stream}
	return x, nil
}

type HealthDiscoveryService_StreamHealthCheckClient interface {
	Send(*HealthCheckRequestOrEndpointHealthResponse) error
	Recv() (*HealthCheckSpecifier, error)
	grpc.ClientStream
}

type healthDiscoveryServiceStreamHealthCheckClient struct {
	grpc.ClientStream
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Send(m *HealthCheckRequestOrEndpointHealthResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Recv() (*HealthCheckSpecifier, error) {
	m := new(HealthCheckSpecifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthDiscoveryServiceClient) FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error) {
	out := new(HealthCheckSpecifier)
	err := c.cc.Invoke(ctx, "/envoy.service.health.v3.HealthDiscoveryService/FetchHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthDiscoveryServiceServer is the server API for HealthDiscoveryService service.
type HealthDiscoveryServiceServer interface {
	StreamHealthCheck(HealthDiscoveryService_StreamHealthCheckServer) error
	FetchHealthCheck(context.Context, *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error)
}

// UnimplementedHealthDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthDiscoveryServiceServer struct {
}

func (*UnimplementedHealthDiscoveryServiceServer) StreamHealthCheck(srv HealthDiscoveryService_StreamHealthCheckServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHealthCheck not implemented")
}
func (*UnimplementedHealthDiscoveryServiceServer) FetchHealthCheck(ctx context.Context, req *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHealthCheck not implemented")
}

func RegisterHealthDiscoveryServiceServer(s *grpc.Server, srv HealthDiscoveryServiceServer) {
	s.RegisterService(&_HealthDiscoveryService_serviceDesc, srv)
}

func _HealthDiscoveryService_StreamHealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthDiscoveryServiceServer).StreamHealthCheck(&healthDiscoveryServiceStreamHealthCheckServer{stream})
}

type HealthDiscoveryService_StreamHealthCheckServer interface {
	Send(*HealthCheckSpecifier) error
	Recv() (*HealthCheckRequestOrEndpointHealthResponse, error)
	grpc.ServerStream
}

type healthDiscoveryServiceStreamHealthCheckServer struct {
	grpc.ServerStream
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Send(m *HealthCheckSpecifier) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Recv() (*HealthCheckRequestOrEndpointHealthResponse, error) {
	m := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HealthDiscoveryService_FetchHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.health.v3.HealthDiscoveryService/FetchHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, req.(*HealthCheckRequestOrEndpointHealthResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.health.v3.HealthDiscoveryService",
	HandlerType: (*HealthDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHealthCheck",
			Handler:    _HealthDiscoveryService_FetchHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHealthCheck",
			Handler:       _HealthDiscoveryService_StreamHealthCheck_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/health/v3/hds.proto",
}