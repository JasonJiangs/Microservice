// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: message.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Week int32

const (
	Week_Monday    Week = 0
	Week_Tuesday   Week = 1
	Week_Wednesday Week = 2
	Week_Thursday  Week = 3
	Week_Friday    Week = 4
	Week_Saturday  Week = 5
	Week_Sunday    Week = 6
)

// Enum value maps for Week.
var (
	Week_name = map[int32]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}
	Week_value = map[string]int32{
		"Monday":    0,
		"Tuesday":   1,
		"Wednesday": 2,
		"Thursday":  3,
		"Friday":    4,
		"Saturday":  5,
		"Sunday":    6,
	}
)

func (x Week) Enum() *Week {
	p := new(Week)
	*p = x
	return p
}

func (x Week) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Week) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Week) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Week) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Week.Descriptor instead.
func (Week) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Address string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Day     Week    `protobuf:"varint,4,opt,name=day,proto3,enum=pb.Week" json:"day,omitempty"`
	P       *People `protobuf:"bytes,5,opt,name=p,proto3" json:"p,omitempty"`
	Score   []int32 `protobuf:"varint,6,rep,packed,name=score,proto3" json:"score,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Student_Email
	Data isStudent_Data `protobuf_oneof:"data"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Student) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Student) GetDay() Week {
	if x != nil {
		return x.Day
	}
	return Week_Monday
}

func (x *Student) GetP() *People {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *Student) GetScore() []int32 {
	if x != nil {
		return x.Score
	}
	return nil
}

func (m *Student) GetData() isStudent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Student) GetEmail() string {
	if x, ok := x.GetData().(*Student_Email); ok {
		return x.Email
	}
	return ""
}

type isStudent_Data interface {
	isStudent_Data()
}

type Student_Email struct {
	Email string `protobuf:"bytes,7,opt,name=email,proto3,oneof"`
}

func (*Student_Email) isStudent_Data() {}

type People struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *People) Reset() {
	*x = People{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *People) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*People) ProtoMessage() {}

func (x *People) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use People.ProtoReflect.Descriptor instead.
func (*People) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *People) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *People) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xb5, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x01, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x52, 0x01, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x06, 0x50,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x2a, 0x62, 0x0a, 0x04, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x57, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x54,
	0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x72, 0x69,
	0x64, 0x61, 0x79, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61,
	0x79, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x10, 0x06, 0x32,
	0x2b, 0x0a, 0x04, 0x61, 0x6a, 0x61, 0x78, 0x12, 0x23, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x1a,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_proto_goTypes = []interface{}{
	(Week)(0),       // 0: pb.Week
	(*Student)(nil), // 1: pb.Student
	(*People)(nil),  // 2: pb.People
}
var file_message_proto_depIdxs = []int32{
	0, // 0: pb.Student.day:type_name -> pb.Week
	2, // 1: pb.Student.p:type_name -> pb.People
	2, // 2: pb.ajax.SayHello:input_type -> pb.People
	1, // 3: pb.ajax.SayHello:output_type -> pb.Student
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*People); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Student_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AjaxClient is the client API for Ajax service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AjaxClient interface {
	SayHello(ctx context.Context, in *People, opts ...grpc.CallOption) (*Student, error)
}

type ajaxClient struct {
	cc grpc.ClientConnInterface
}

func NewAjaxClient(cc grpc.ClientConnInterface) AjaxClient {
	return &ajaxClient{cc}
}

func (c *ajaxClient) SayHello(ctx context.Context, in *People, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/pb.ajax/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AjaxServer is the server API for Ajax service.
type AjaxServer interface {
	SayHello(context.Context, *People) (*Student, error)
}

// UnimplementedAjaxServer can be embedded to have forward compatible implementations.
type UnimplementedAjaxServer struct {
}

func (*UnimplementedAjaxServer) SayHello(context.Context, *People) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterAjaxServer(s *grpc.Server, srv AjaxServer) {
	s.RegisterService(&_Ajax_serviceDesc, srv)
}

func _Ajax_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(People)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AjaxServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ajax/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AjaxServer).SayHello(ctx, req.(*People))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ajax_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ajax",
	HandlerType: (*AjaxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Ajax_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
