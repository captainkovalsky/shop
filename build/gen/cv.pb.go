// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: cv.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Language_Level int32

const (
	Language_UPPER_INTERMEDIATE Language_Level = 0
	Language_ADVANCED           Language_Level = 1
	Language_FLUENT             Language_Level = 2
	Language_NATIVE             Language_Level = 3
)

// Enum value maps for Language_Level.
var (
	Language_Level_name = map[int32]string{
		0: "UPPER_INTERMEDIATE",
		1: "ADVANCED",
		2: "FLUENT",
		3: "NATIVE",
	}
	Language_Level_value = map[string]int32{
		"UPPER_INTERMEDIATE": 0,
		"ADVANCED":           1,
		"FLUENT":             2,
		"NATIVE":             3,
	}
)

func (x Language_Level) Enum() *Language_Level {
	p := new(Language_Level)
	*p = x
	return p
}

func (x Language_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_cv_proto_enumTypes[0].Descriptor()
}

func (Language_Level) Type() protoreflect.EnumType {
	return &file_cv_proto_enumTypes[0]
}

func (x Language_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language_Level.Descriptor instead.
func (Language_Level) EnumDescriptor() ([]byte, []int) {
	return file_cv_proto_rawDescGZIP(), []int{4, 0}
}

type Experience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate *timestamp.Timestamp `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Project   *Project             `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Company   *Company             `protobuf:"bytes,4,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *Experience) Reset() {
	*x = Experience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experience) ProtoMessage() {}

func (x *Experience) ProtoReflect() protoreflect.Message {
	mi := &file_cv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experience.ProtoReflect.Descriptor instead.
func (*Experience) Descriptor() ([]byte, []int) {
	return file_cv_proto_rawDescGZIP(), []int{0}
}

func (x *Experience) GetStartDate() *timestamp.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Experience) GetEndDate() *timestamp.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Experience) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *Experience) GetCompany() *Company {
	if x != nil {
		return x.Company
	}
	return nil
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	SiteUrl  string `protobuf:"bytes,3,opt,name=siteUrl,proto3" json:"siteUrl,omitempty"`
	Logo     string `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_cv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_cv_proto_rawDescGZIP(), []int{1}
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Company) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

func (x *Company) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Technology   []string `protobuf:"bytes,3,rep,name=technology,proto3" json:"technology,omitempty"`
	UseScrum     bool     `protobuf:"varint,4,opt,name=useScrum,proto3" json:"useScrum,omitempty"`
	BackendSize  int32    `protobuf:"varint,5,opt,name=backendSize,proto3" json:"backendSize,omitempty"`
	FrontendSize int32    `protobuf:"varint,6,opt,name=frontendSize,proto3" json:"frontendSize,omitempty"`
	QaSize       int32    `protobuf:"varint,7,opt,name=qaSize,proto3" json:"qaSize,omitempty"`
	SiteUrl      string   `protobuf:"bytes,8,opt,name=siteUrl,proto3" json:"siteUrl,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_cv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_cv_proto_rawDescGZIP(), []int{2}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetTechnology() []string {
	if x != nil {
		return x.Technology
	}
	return nil
}

func (x *Project) GetUseScrum() bool {
	if x != nil {
		return x.UseScrum
	}
	return false
}

func (x *Project) GetBackendSize() int32 {
	if x != nil {
		return x.BackendSize
	}
	return 0
}

func (x *Project) GetFrontendSize() int32 {
	if x != nil {
		return x.FrontendSize
	}
	return 0
}

func (x *Project) GetQaSize() int32 {
	if x != nil {
		return x.QaSize
	}
	return 0
}

func (x *Project) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

type CV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experience []*Experience `protobuf:"bytes,1,rep,name=Experience,proto3" json:"Experience,omitempty"`
	Firstname  string        `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Surname    string        `protobuf:"bytes,3,opt,name=Surname,proto3" json:"Surname,omitempty"`
	Language   []*Language   `protobuf:"bytes,4,rep,name=language,proto3" json:"language,omitempty"`
}

func (x *CV) Reset() {
	*x = CV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CV) ProtoMessage() {}

func (x *CV) ProtoReflect() protoreflect.Message {
	mi := &file_cv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CV.ProtoReflect.Descriptor instead.
func (*CV) Descriptor() ([]byte, []int) {
	return file_cv_proto_rawDescGZIP(), []int{3}
}

func (x *CV) GetExperience() []*Experience {
	if x != nil {
		return x.Experience
	}
	return nil
}

func (x *CV) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CV) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *CV) GetLanguage() []*Language {
	if x != nil {
		return x.Language
	}
	return nil
}

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string         `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Level Language_Level `protobuf:"varint,2,opt,name=level,proto3,enum=Language_Level" json:"level,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_cv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_cv_proto_rawDescGZIP(), []int{4}
}

func (x *Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Language) GetLevel() Language_Level {
	if x != nil {
		return x.Level
	}
	return Language_UPPER_INTERMEDIATE
}

var File_cv_proto protoreflect.FileDescriptor

var file_cv_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0a,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x22, 0x67, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x53, 0x63, 0x72, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x53, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x71, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x71, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72,
	0x6c, 0x22, 0x90, 0x01, 0x0a, 0x02, 0x43, 0x56, 0x12, 0x2b, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x45, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x4c, 0x55, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x03, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x6b, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x6b,
	0x79, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cv_proto_rawDescOnce sync.Once
	file_cv_proto_rawDescData = file_cv_proto_rawDesc
)

func file_cv_proto_rawDescGZIP() []byte {
	file_cv_proto_rawDescOnce.Do(func() {
		file_cv_proto_rawDescData = protoimpl.X.CompressGZIP(file_cv_proto_rawDescData)
	})
	return file_cv_proto_rawDescData
}

var file_cv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cv_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cv_proto_goTypes = []interface{}{
	(Language_Level)(0),         // 0: Language.Level
	(*Experience)(nil),          // 1: Experience
	(*Company)(nil),             // 2: Company
	(*Project)(nil),             // 3: Project
	(*CV)(nil),                  // 4: CV
	(*Language)(nil),            // 5: Language
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_cv_proto_depIdxs = []int32{
	6, // 0: Experience.startDate:type_name -> google.protobuf.Timestamp
	6, // 1: Experience.endDate:type_name -> google.protobuf.Timestamp
	3, // 2: Experience.project:type_name -> Project
	2, // 3: Experience.company:type_name -> Company
	1, // 4: CV.Experience:type_name -> Experience
	5, // 5: CV.language:type_name -> Language
	0, // 6: Language.level:type_name -> Language.Level
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cv_proto_init() }
func file_cv_proto_init() {
	if File_cv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experience); i {
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
		file_cv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_cv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_cv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CV); i {
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
		file_cv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cv_proto_goTypes,
		DependencyIndexes: file_cv_proto_depIdxs,
		EnumInfos:         file_cv_proto_enumTypes,
		MessageInfos:      file_cv_proto_msgTypes,
	}.Build()
	File_cv_proto = out.File
	file_cv_proto_rawDesc = nil
	file_cv_proto_goTypes = nil
	file_cv_proto_depIdxs = nil
}
