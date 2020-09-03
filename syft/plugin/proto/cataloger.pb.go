// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cataloger.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FileReference struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileReference) Reset()         { *m = FileReference{} }
func (m *FileReference) String() string { return proto.CompactTextString(m) }
func (*FileReference) ProtoMessage()    {}
func (*FileReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{0}
}

func (m *FileReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileReference.Unmarshal(m, b)
}
func (m *FileReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileReference.Marshal(b, m, deterministic)
}
func (m *FileReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileReference.Merge(m, src)
}
func (m *FileReference) XXX_Size() int {
	return xxx_messageInfo_FileReference.Size(m)
}
func (m *FileReference) XXX_DiscardUnknown() {
	xxx_messageInfo_FileReference.DiscardUnknown(m)
}

var xxx_messageInfo_FileReference proto.InternalMessageInfo

func (m *FileReference) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FileReference) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type FileReferenceContents struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Contents             string   `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileReferenceContents) Reset()         { *m = FileReferenceContents{} }
func (m *FileReferenceContents) String() string { return proto.CompactTextString(m) }
func (*FileReferenceContents) ProtoMessage()    {}
func (*FileReferenceContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{1}
}

func (m *FileReferenceContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileReferenceContents.Unmarshal(m, b)
}
func (m *FileReferenceContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileReferenceContents.Marshal(b, m, deterministic)
}
func (m *FileReferenceContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileReferenceContents.Merge(m, src)
}
func (m *FileReferenceContents) XXX_Size() int {
	return xxx_messageInfo_FileReferenceContents.Size(m)
}
func (m *FileReferenceContents) XXX_DiscardUnknown() {
	xxx_messageInfo_FileReferenceContents.DiscardUnknown(m)
}

var xxx_messageInfo_FileReferenceContents proto.InternalMessageInfo

func (m *FileReferenceContents) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FileReferenceContents) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileReferenceContents) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type FileResolverRequest struct {
	Paths                []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResolverRequest) Reset()         { *m = FileResolverRequest{} }
func (m *FileResolverRequest) String() string { return proto.CompactTextString(m) }
func (*FileResolverRequest) ProtoMessage()    {}
func (*FileResolverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{3}
}

func (m *FileResolverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResolverRequest.Unmarshal(m, b)
}
func (m *FileResolverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResolverRequest.Marshal(b, m, deterministic)
}
func (m *FileResolverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResolverRequest.Merge(m, src)
}
func (m *FileResolverRequest) XXX_Size() int {
	return xxx_messageInfo_FileResolverRequest.Size(m)
}
func (m *FileResolverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResolverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileResolverRequest proto.InternalMessageInfo

func (m *FileResolverRequest) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type FileResolverResponse struct {
	Files                []*FileReference `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FileResolverResponse) Reset()         { *m = FileResolverResponse{} }
func (m *FileResolverResponse) String() string { return proto.CompactTextString(m) }
func (*FileResolverResponse) ProtoMessage()    {}
func (*FileResolverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{4}
}

func (m *FileResolverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResolverResponse.Unmarshal(m, b)
}
func (m *FileResolverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResolverResponse.Marshal(b, m, deterministic)
}
func (m *FileResolverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResolverResponse.Merge(m, src)
}
func (m *FileResolverResponse) XXX_Size() int {
	return xxx_messageInfo_FileResolverResponse.Size(m)
}
func (m *FileResolverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResolverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileResolverResponse proto.InternalMessageInfo

func (m *FileResolverResponse) GetFiles() []*FileReference {
	if m != nil {
		return m.Files
	}
	return nil
}

type NameResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameResponse) Reset()         { *m = NameResponse{} }
func (m *NameResponse) String() string { return proto.CompactTextString(m) }
func (*NameResponse) ProtoMessage()    {}
func (*NameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{5}
}

func (m *NameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameResponse.Unmarshal(m, b)
}
func (m *NameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameResponse.Marshal(b, m, deterministic)
}
func (m *NameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameResponse.Merge(m, src)
}
func (m *NameResponse) XXX_Size() int {
	return xxx_messageInfo_NameResponse.Size(m)
}
func (m *NameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NameResponse proto.InternalMessageInfo

func (m *NameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SelectFilesRequest struct {
	FileResolverBrokerId uint32   `protobuf:"varint,1,opt,name=fileResolverBrokerId,proto3" json:"fileResolverBrokerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectFilesRequest) Reset()         { *m = SelectFilesRequest{} }
func (m *SelectFilesRequest) String() string { return proto.CompactTextString(m) }
func (*SelectFilesRequest) ProtoMessage()    {}
func (*SelectFilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{6}
}

func (m *SelectFilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectFilesRequest.Unmarshal(m, b)
}
func (m *SelectFilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectFilesRequest.Marshal(b, m, deterministic)
}
func (m *SelectFilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectFilesRequest.Merge(m, src)
}
func (m *SelectFilesRequest) XXX_Size() int {
	return xxx_messageInfo_SelectFilesRequest.Size(m)
}
func (m *SelectFilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectFilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelectFilesRequest proto.InternalMessageInfo

func (m *SelectFilesRequest) GetFileResolverBrokerId() uint32 {
	if m != nil {
		return m.FileResolverBrokerId
	}
	return 0
}

type SelectFilesResponse struct {
	Files                []*FileReference `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SelectFilesResponse) Reset()         { *m = SelectFilesResponse{} }
func (m *SelectFilesResponse) String() string { return proto.CompactTextString(m) }
func (*SelectFilesResponse) ProtoMessage()    {}
func (*SelectFilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{7}
}

func (m *SelectFilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectFilesResponse.Unmarshal(m, b)
}
func (m *SelectFilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectFilesResponse.Marshal(b, m, deterministic)
}
func (m *SelectFilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectFilesResponse.Merge(m, src)
}
func (m *SelectFilesResponse) XXX_Size() int {
	return xxx_messageInfo_SelectFilesResponse.Size(m)
}
func (m *SelectFilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectFilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SelectFilesResponse proto.InternalMessageInfo

func (m *SelectFilesResponse) GetFiles() []*FileReference {
	if m != nil {
		return m.Files
	}
	return nil
}

type Package struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	FoundBy              string            `protobuf:"bytes,3,opt,name=foundBy,proto3" json:"foundBy,omitempty"`
	Source               []*FileReference  `protobuf:"bytes,4,rep,name=source,proto3" json:"source,omitempty"`
	Licenses             []string          `protobuf:"bytes,5,rep,name=licenses,proto3" json:"licenses,omitempty"`
	Language             uint64            `protobuf:"varint,6,opt,name=language,proto3" json:"language,omitempty"`
	Type                 string            `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{8}
}

func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Package) GetFoundBy() string {
	if m != nil {
		return m.FoundBy
	}
	return ""
}

func (m *Package) GetSource() []*FileReference {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Package) GetLicenses() []string {
	if m != nil {
		return m.Licenses
	}
	return nil
}

func (m *Package) GetLanguage() uint64 {
	if m != nil {
		return m.Language
	}
	return 0
}

func (m *Package) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Package) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CatalogRequest struct {
	Contents             []*FileReferenceContents `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CatalogRequest) Reset()         { *m = CatalogRequest{} }
func (m *CatalogRequest) String() string { return proto.CompactTextString(m) }
func (*CatalogRequest) ProtoMessage()    {}
func (*CatalogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{9}
}

func (m *CatalogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogRequest.Unmarshal(m, b)
}
func (m *CatalogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogRequest.Marshal(b, m, deterministic)
}
func (m *CatalogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogRequest.Merge(m, src)
}
func (m *CatalogRequest) XXX_Size() int {
	return xxx_messageInfo_CatalogRequest.Size(m)
}
func (m *CatalogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogRequest proto.InternalMessageInfo

func (m *CatalogRequest) GetContents() []*FileReferenceContents {
	if m != nil {
		return m.Contents
	}
	return nil
}

type CatalogResponse struct {
	Package              []*Package `protobuf:"bytes,1,rep,name=package,proto3" json:"package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CatalogResponse) Reset()         { *m = CatalogResponse{} }
func (m *CatalogResponse) String() string { return proto.CompactTextString(m) }
func (*CatalogResponse) ProtoMessage()    {}
func (*CatalogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_346aaad5d7f2c092, []int{10}
}

func (m *CatalogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogResponse.Unmarshal(m, b)
}
func (m *CatalogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogResponse.Marshal(b, m, deterministic)
}
func (m *CatalogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogResponse.Merge(m, src)
}
func (m *CatalogResponse) XXX_Size() int {
	return xxx_messageInfo_CatalogResponse.Size(m)
}
func (m *CatalogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogResponse proto.InternalMessageInfo

func (m *CatalogResponse) GetPackage() []*Package {
	if m != nil {
		return m.Package
	}
	return nil
}

func init() {
	proto.RegisterType((*FileReference)(nil), "proto.FileReference")
	proto.RegisterType((*FileReferenceContents)(nil), "proto.FileReferenceContents")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*FileResolverRequest)(nil), "proto.FileResolverRequest")
	proto.RegisterType((*FileResolverResponse)(nil), "proto.FileResolverResponse")
	proto.RegisterType((*NameResponse)(nil), "proto.NameResponse")
	proto.RegisterType((*SelectFilesRequest)(nil), "proto.SelectFilesRequest")
	proto.RegisterType((*SelectFilesResponse)(nil), "proto.SelectFilesResponse")
	proto.RegisterType((*Package)(nil), "proto.Package")
	proto.RegisterMapType((map[string]string)(nil), "proto.Package.MetadataEntry")
	proto.RegisterType((*CatalogRequest)(nil), "proto.CatalogRequest")
	proto.RegisterType((*CatalogResponse)(nil), "proto.CatalogResponse")
}

func init() { proto.RegisterFile("cataloger.proto", fileDescriptor_346aaad5d7f2c092) }

var fileDescriptor_346aaad5d7f2c092 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x8e, 0xd3, 0x3c,
	0x14, 0x56, 0x7a, 0x4b, 0x7b, 0x7a, 0x99, 0x5f, 0x6e, 0xe7, 0x57, 0x08, 0x2c, 0xaa, 0xac, 0xca,
	0x45, 0x5d, 0x74, 0x36, 0x15, 0xb3, 0xa2, 0x65, 0x86, 0x8b, 0x04, 0x1a, 0x99, 0x05, 0x6b, 0x4f,
	0x7a, 0x5a, 0xa2, 0xa6, 0x71, 0x88, 0xdd, 0x4a, 0x79, 0x19, 0x1e, 0x83, 0x07, 0xe0, 0xc9, 0x50,
	0x6c, 0x27, 0xd4, 0x10, 0x21, 0x34, 0xab, 0xf8, 0xe4, 0xfb, 0xfc, 0x9d, 0xcf, 0xe7, 0x02, 0x17,
	0x21, 0x93, 0x2c, 0xe6, 0x3b, 0xcc, 0xe6, 0x69, 0xc6, 0x25, 0x27, 0x6d, 0xf5, 0x09, 0xae, 0x60,
	0x78, 0x1b, 0xc5, 0x48, 0x71, 0x8b, 0x19, 0x26, 0x21, 0x92, 0x11, 0x34, 0xa2, 0x8d, 0xe7, 0x4c,
	0x9d, 0x59, 0x93, 0x36, 0xa2, 0x0d, 0x21, 0xd0, 0x4a, 0x99, 0xfc, 0xe2, 0x35, 0xa6, 0xce, 0xac,
	0x47, 0xd5, 0x39, 0xf8, 0x0c, 0x97, 0xd6, 0xa5, 0x35, 0x4f, 0x24, 0x26, 0x52, 0xfc, 0xcb, 0x65,
	0xe2, 0x43, 0x37, 0x34, 0x7c, 0xaf, 0xa9, 0xfe, 0x57, 0x71, 0xe0, 0x42, 0xfb, 0xe6, 0x90, 0xca,
	0x3c, 0x78, 0x0e, 0x63, 0x9d, 0x41, 0xf0, 0xf8, 0x84, 0x19, 0xc5, 0xaf, 0x47, 0x14, 0x92, 0x4c,
	0xa0, 0x5d, 0x68, 0x08, 0xcf, 0x99, 0x36, 0x67, 0x3d, 0xaa, 0x83, 0x60, 0x05, 0x13, 0x9b, 0x2c,
	0x52, 0x9e, 0x08, 0x24, 0xcf, 0xa0, 0xbd, 0x8d, 0x62, 0xd4, 0xec, 0xfe, 0x62, 0xa2, 0x5f, 0x3e,
	0xb7, 0xac, 0x53, 0x4d, 0x09, 0x02, 0x18, 0x7c, 0x64, 0x07, 0xac, 0xee, 0x12, 0x68, 0x25, 0xec,
	0x80, 0xea, 0x2d, 0x3d, 0xaa, 0xce, 0xc1, 0x5b, 0x20, 0x9f, 0x30, 0xc6, 0x50, 0x16, 0x0a, 0xa2,
	0xf4, 0xb4, 0x80, 0xc9, 0xf6, 0x2c, 0xfb, 0x2a, 0xe3, 0x7b, 0xcc, 0xde, 0xe9, 0x2a, 0x0c, 0x69,
	0x2d, 0x16, 0xbc, 0x82, 0xb1, 0xa5, 0xf4, 0x00, 0xc3, 0x3f, 0x1a, 0xe0, 0xde, 0xb1, 0x70, 0xcf,
	0x76, 0xb5, 0x66, 0x89, 0x07, 0xee, 0x09, 0x33, 0x11, 0xf1, 0xc4, 0x54, 0xbf, 0x0c, 0x0b, 0x64,
	0xcb, 0x8f, 0xc9, 0x66, 0x95, 0x9b, 0xfa, 0x97, 0x21, 0x79, 0x01, 0x1d, 0xc1, 0x8f, 0x59, 0x88,
	0x5e, 0xeb, 0x2f, 0x06, 0x0c, 0xa7, 0x68, 0x64, 0x1c, 0x85, 0x98, 0x08, 0x14, 0x5e, 0x5b, 0xf5,
	0xa3, 0x8a, 0x15, 0xc6, 0x92, 0xdd, 0x91, 0xed, 0xd0, 0xeb, 0x4c, 0x9d, 0x59, 0x8b, 0x56, 0x71,
	0xe1, 0x56, 0xe6, 0x29, 0x7a, 0xae, 0x76, 0x5b, 0x9c, 0xc9, 0x12, 0xba, 0x07, 0x94, 0x6c, 0xc3,
	0x24, 0xf3, 0xba, 0x2a, 0xf7, 0x13, 0x93, 0xdb, 0xbc, 0x71, 0xfe, 0xc1, 0xc0, 0x37, 0x89, 0xcc,
	0x72, 0x5a, 0xb1, 0xfd, 0x6b, 0x18, 0x5a, 0x10, 0xf9, 0x0f, 0x9a, 0x7b, 0xcc, 0x4d, 0x2d, 0x8a,
	0x63, 0x31, 0x35, 0x27, 0x16, 0x1f, 0xd1, 0x14, 0x42, 0x07, 0x2f, 0x1b, 0x4b, 0x27, 0x78, 0x0f,
	0xa3, 0xb5, 0xde, 0x8b, 0xb2, 0x9b, 0xcb, 0xb3, 0xe9, 0x74, 0x2c, 0x23, 0xb5, 0x13, 0x7f, 0x36,
	0xbb, 0xd7, 0x70, 0x51, 0x69, 0x99, 0x7e, 0xce, 0xc0, 0x4d, 0xb5, 0x7d, 0xa3, 0x35, 0xb2, 0x1f,
	0x45, 0x4b, 0x78, 0xf1, 0xdd, 0x81, 0xde, 0xba, 0xdc, 0x50, 0xf2, 0x14, 0x5a, 0xc5, 0x30, 0x92,
	0x81, 0xa1, 0xab, 0x9d, 0xf0, 0xc7, 0x26, 0xb2, 0xe6, 0xf4, 0x35, 0xf4, 0xcf, 0x26, 0x89, 0x3c,
	0x32, 0x9c, 0x3f, 0xe7, 0xd4, 0xf7, 0xeb, 0x20, 0xa3, 0xb2, 0x04, 0xd7, 0x64, 0x27, 0x97, 0x86,
	0x66, 0xd7, 0xc5, 0xff, 0xff, 0xf7, 0xdf, 0xfa, 0xe6, 0xe2, 0x9b, 0x03, 0x83, 0xf3, 0xe5, 0x23,
	0xb7, 0xd0, 0x57, 0xda, 0xab, 0xfc, 0x4e, 0x6d, 0xbb, 0x55, 0x3d, 0x6b, 0x9b, 0xfd, 0xc7, 0xb5,
	0x98, 0xb1, 0xf4, 0x4b, 0xe7, 0x4d, 0xcc, 0xef, 0x1f, 0xac, 0x73, 0xdf, 0x51, 0xd8, 0xd5, 0xcf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xfe, 0x04, 0x3e, 0x01, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatalogerClient is the client API for Cataloger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogerClient interface {
	Name(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NameResponse, error)
	SelectFiles(ctx context.Context, in *SelectFilesRequest, opts ...grpc.CallOption) (*SelectFilesResponse, error)
	Catalog(ctx context.Context, in *CatalogRequest, opts ...grpc.CallOption) (*CatalogResponse, error)
}

type catalogerClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogerClient(cc grpc.ClientConnInterface) CatalogerClient {
	return &catalogerClient{cc}
}

func (c *catalogerClient) Name(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/proto.Cataloger/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogerClient) SelectFiles(ctx context.Context, in *SelectFilesRequest, opts ...grpc.CallOption) (*SelectFilesResponse, error) {
	out := new(SelectFilesResponse)
	err := c.cc.Invoke(ctx, "/proto.Cataloger/SelectFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogerClient) Catalog(ctx context.Context, in *CatalogRequest, opts ...grpc.CallOption) (*CatalogResponse, error) {
	out := new(CatalogResponse)
	err := c.cc.Invoke(ctx, "/proto.Cataloger/Catalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogerServer is the server API for Cataloger service.
type CatalogerServer interface {
	Name(context.Context, *Empty) (*NameResponse, error)
	SelectFiles(context.Context, *SelectFilesRequest) (*SelectFilesResponse, error)
	Catalog(context.Context, *CatalogRequest) (*CatalogResponse, error)
}

// UnimplementedCatalogerServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogerServer struct {
}

func (*UnimplementedCatalogerServer) Name(ctx context.Context, req *Empty) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (*UnimplementedCatalogerServer) SelectFiles(ctx context.Context, req *SelectFilesRequest) (*SelectFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectFiles not implemented")
}
func (*UnimplementedCatalogerServer) Catalog(ctx context.Context, req *CatalogRequest) (*CatalogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Catalog not implemented")
}

func RegisterCatalogerServer(s *grpc.Server, srv CatalogerServer) {
	s.RegisterService(&_Cataloger_serviceDesc, srv)
}

func _Cataloger_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogerServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cataloger/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogerServer).Name(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cataloger_SelectFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogerServer).SelectFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cataloger/SelectFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogerServer).SelectFiles(ctx, req.(*SelectFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cataloger_Catalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogerServer).Catalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cataloger/Catalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogerServer).Catalog(ctx, req.(*CatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cataloger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Cataloger",
	HandlerType: (*CatalogerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _Cataloger_Name_Handler,
		},
		{
			MethodName: "SelectFiles",
			Handler:    _Cataloger_SelectFiles_Handler,
		},
		{
			MethodName: "Catalog",
			Handler:    _Cataloger_Catalog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cataloger.proto",
}

// FileResolverClient is the client API for FileResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileResolverClient interface {
	FilesByPath(ctx context.Context, in *FileResolverRequest, opts ...grpc.CallOption) (*FileResolverResponse, error)
	FilesByGlob(ctx context.Context, in *FileResolverRequest, opts ...grpc.CallOption) (*FileResolverResponse, error)
}

type fileResolverClient struct {
	cc grpc.ClientConnInterface
}

func NewFileResolverClient(cc grpc.ClientConnInterface) FileResolverClient {
	return &fileResolverClient{cc}
}

func (c *fileResolverClient) FilesByPath(ctx context.Context, in *FileResolverRequest, opts ...grpc.CallOption) (*FileResolverResponse, error) {
	out := new(FileResolverResponse)
	err := c.cc.Invoke(ctx, "/proto.FileResolver/FilesByPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileResolverClient) FilesByGlob(ctx context.Context, in *FileResolverRequest, opts ...grpc.CallOption) (*FileResolverResponse, error) {
	out := new(FileResolverResponse)
	err := c.cc.Invoke(ctx, "/proto.FileResolver/FilesByGlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileResolverServer is the server API for FileResolver service.
type FileResolverServer interface {
	FilesByPath(context.Context, *FileResolverRequest) (*FileResolverResponse, error)
	FilesByGlob(context.Context, *FileResolverRequest) (*FileResolverResponse, error)
}

// UnimplementedFileResolverServer can be embedded to have forward compatible implementations.
type UnimplementedFileResolverServer struct {
}

func (*UnimplementedFileResolverServer) FilesByPath(ctx context.Context, req *FileResolverRequest) (*FileResolverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilesByPath not implemented")
}
func (*UnimplementedFileResolverServer) FilesByGlob(ctx context.Context, req *FileResolverRequest) (*FileResolverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilesByGlob not implemented")
}

func RegisterFileResolverServer(s *grpc.Server, srv FileResolverServer) {
	s.RegisterService(&_FileResolver_serviceDesc, srv)
}

func _FileResolver_FilesByPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileResolverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileResolverServer).FilesByPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileResolver/FilesByPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileResolverServer).FilesByPath(ctx, req.(*FileResolverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileResolver_FilesByGlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileResolverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileResolverServer).FilesByGlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileResolver/FilesByGlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileResolverServer).FilesByGlob(ctx, req.(*FileResolverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileResolver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FileResolver",
	HandlerType: (*FileResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FilesByPath",
			Handler:    _FileResolver_FilesByPath_Handler,
		},
		{
			MethodName: "FilesByGlob",
			Handler:    _FileResolver_FilesByGlob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cataloger.proto",
}
