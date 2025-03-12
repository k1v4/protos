// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.2
// source: shoe/shoe.proto

package shoev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Shoe struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShoeId        int64                  `protobuf:"varint,1,opt,name=shoe_id,json=shoeId,proto3" json:"shoe_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl      string                 `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	UserId        int64                  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Shoe) Reset() {
	*x = Shoe{}
	mi := &file_shoe_shoe_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Shoe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shoe) ProtoMessage() {}

func (x *Shoe) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shoe.ProtoReflect.Descriptor instead.
func (*Shoe) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{0}
}

func (x *Shoe) GetShoeId() int64 {
	if x != nil {
		return x.ShoeId
	}
	return 0
}

func (x *Shoe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shoe) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Shoe) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddShoeRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageData []byte                 `protobuf:"bytes,4,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	// int64 user_id = 3;
	ImageName     string `protobuf:"bytes,5,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddShoeRequest) Reset() {
	*x = AddShoeRequest{}
	mi := &file_shoe_shoe_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddShoeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShoeRequest) ProtoMessage() {}

func (x *AddShoeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShoeRequest.ProtoReflect.Descriptor instead.
func (*AddShoeRequest) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{1}
}

func (x *AddShoeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddShoeRequest) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *AddShoeRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

type AddShoeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShoeId        int64                  `protobuf:"varint,1,opt,name=shoe_id,json=shoeId,proto3" json:"shoe_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddShoeResponse) Reset() {
	*x = AddShoeResponse{}
	mi := &file_shoe_shoe_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddShoeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShoeResponse) ProtoMessage() {}

func (x *AddShoeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShoeResponse.ProtoReflect.Descriptor instead.
func (*AddShoeResponse) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{2}
}

func (x *AddShoeResponse) GetShoeId() int64 {
	if x != nil {
		return x.ShoeId
	}
	return 0
}

type GetShoeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShoeId        int64                  `protobuf:"varint,1,opt,name=shoe_id,json=shoeId,proto3" json:"shoe_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShoeRequest) Reset() {
	*x = GetShoeRequest{}
	mi := &file_shoe_shoe_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShoeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShoeRequest) ProtoMessage() {}

func (x *GetShoeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShoeRequest.ProtoReflect.Descriptor instead.
func (*GetShoeRequest) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{3}
}

func (x *GetShoeRequest) GetShoeId() int64 {
	if x != nil {
		return x.ShoeId
	}
	return 0
}

type GetShoeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shoe          *Shoe                  `protobuf:"bytes,1,opt,name=shoe,proto3" json:"shoe,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShoeResponse) Reset() {
	*x = GetShoeResponse{}
	mi := &file_shoe_shoe_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShoeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShoeResponse) ProtoMessage() {}

func (x *GetShoeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShoeResponse.ProtoReflect.Descriptor instead.
func (*GetShoeResponse) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{4}
}

func (x *GetShoeResponse) GetShoe() *Shoe {
	if x != nil {
		return x.Shoe
	}
	return nil
}

type DeleteShoeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShoeId        int64                  `protobuf:"varint,1,opt,name=shoe_id,json=shoeId,proto3" json:"shoe_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteShoeRequest) Reset() {
	*x = DeleteShoeRequest{}
	mi := &file_shoe_shoe_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShoeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShoeRequest) ProtoMessage() {}

func (x *DeleteShoeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShoeRequest.ProtoReflect.Descriptor instead.
func (*DeleteShoeRequest) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteShoeRequest) GetShoeId() int64 {
	if x != nil {
		return x.ShoeId
	}
	return 0
}

type DeleteShoeResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	IsSuccessfully bool                   `protobuf:"varint,1,opt,name=isSuccessfully,proto3" json:"isSuccessfully,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeleteShoeResponse) Reset() {
	*x = DeleteShoeResponse{}
	mi := &file_shoe_shoe_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShoeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShoeResponse) ProtoMessage() {}

func (x *DeleteShoeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShoeResponse.ProtoReflect.Descriptor instead.
func (*DeleteShoeResponse) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteShoeResponse) GetIsSuccessfully() bool {
	if x != nil {
		return x.IsSuccessfully
	}
	return false
}

type UpdateShoeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShoeId        int64                  `protobuf:"varint,1,opt,name=shoe_id,json=shoeId,proto3" json:"shoe_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageData     []byte                 `protobuf:"bytes,5,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	NewFileName   string                 `protobuf:"bytes,6,opt,name=new_file_name,json=newFileName,proto3" json:"new_file_name,omitempty"` //int64 user_id = 4;
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateShoeRequest) Reset() {
	*x = UpdateShoeRequest{}
	mi := &file_shoe_shoe_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateShoeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShoeRequest) ProtoMessage() {}

func (x *UpdateShoeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShoeRequest.ProtoReflect.Descriptor instead.
func (*UpdateShoeRequest) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateShoeRequest) GetShoeId() int64 {
	if x != nil {
		return x.ShoeId
	}
	return 0
}

func (x *UpdateShoeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateShoeRequest) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *UpdateShoeRequest) GetNewFileName() string {
	if x != nil {
		return x.NewFileName
	}
	return ""
}

type UpdateShoeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shoe          *Shoe                  `protobuf:"bytes,1,opt,name=shoe,proto3" json:"shoe,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateShoeResponse) Reset() {
	*x = UpdateShoeResponse{}
	mi := &file_shoe_shoe_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateShoeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShoeResponse) ProtoMessage() {}

func (x *UpdateShoeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShoeResponse.ProtoReflect.Descriptor instead.
func (*UpdateShoeResponse) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateShoeResponse) GetShoe() *Shoe {
	if x != nil {
		return x.Shoe
	}
	return nil
}

type GetAllShoesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllShoesRequest) Reset() {
	*x = GetAllShoesRequest{}
	mi := &file_shoe_shoe_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllShoesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllShoesRequest) ProtoMessage() {}

func (x *GetAllShoesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllShoesRequest.ProtoReflect.Descriptor instead.
func (*GetAllShoesRequest) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{9}
}

type GetAllShoesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shoes         []*Shoe                `protobuf:"bytes,1,rep,name=shoes,proto3" json:"shoes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllShoesResponse) Reset() {
	*x = GetAllShoesResponse{}
	mi := &file_shoe_shoe_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllShoesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllShoesResponse) ProtoMessage() {}

func (x *GetAllShoesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoe_shoe_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllShoesResponse.ProtoReflect.Descriptor instead.
func (*GetAllShoesResponse) Descriptor() ([]byte, []int) {
	return file_shoe_shoe_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllShoesResponse) GetShoes() []*Shoe {
	if x != nil {
		return x.Shoes
	}
	return nil
}

var File_shoe_shoe_proto protoreflect.FileDescriptor

var file_shoe_shoe_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x68, 0x6f, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x68, 0x6f, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x68, 0x6f, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x62, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x65, 0x49, 0x64,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x65, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x04, 0x73, 0x68, 0x6f, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73,
	0x68, 0x6f, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x65, 0x22, 0x2c,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66,
	0x75, 0x6c, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d,
	0x6e, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x34, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x65,
	0x52, 0x04, 0x73, 0x68, 0x6f, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x68, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x05,
	0x73, 0x68, 0x6f, 0x65, 0x73, 0x32, 0xd7, 0x03, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f, 0x65,
	0x12, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x6f, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x6f, 0x65, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x60, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65, 0x12, 0x17,
	0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x6f, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x63, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65,
	0x12, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x1a, 0x17,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x2f, 0x7b, 0x73,
	0x68, 0x6f, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x6f, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x68, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x68, 0x6f, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x42,
	0x16, 0x5a, 0x14, 0x6b, 0x31, 0x76, 0x63, 0x68, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x2e, 0x76, 0x31,
	0x3b, 0x73, 0x68, 0x6f, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_shoe_shoe_proto_rawDescOnce sync.Once
	file_shoe_shoe_proto_rawDescData []byte
)

func file_shoe_shoe_proto_rawDescGZIP() []byte {
	file_shoe_shoe_proto_rawDescOnce.Do(func() {
		file_shoe_shoe_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shoe_shoe_proto_rawDesc), len(file_shoe_shoe_proto_rawDesc)))
	})
	return file_shoe_shoe_proto_rawDescData
}

var file_shoe_shoe_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_shoe_shoe_proto_goTypes = []any{
	(*Shoe)(nil),                // 0: shoe.Shoe
	(*AddShoeRequest)(nil),      // 1: shoe.AddShoeRequest
	(*AddShoeResponse)(nil),     // 2: shoe.AddShoeResponse
	(*GetShoeRequest)(nil),      // 3: shoe.GetShoeRequest
	(*GetShoeResponse)(nil),     // 4: shoe.GetShoeResponse
	(*DeleteShoeRequest)(nil),   // 5: shoe.DeleteShoeRequest
	(*DeleteShoeResponse)(nil),  // 6: shoe.DeleteShoeResponse
	(*UpdateShoeRequest)(nil),   // 7: shoe.UpdateShoeRequest
	(*UpdateShoeResponse)(nil),  // 8: shoe.UpdateShoeResponse
	(*GetAllShoesRequest)(nil),  // 9: shoe.GetAllShoesRequest
	(*GetAllShoesResponse)(nil), // 10: shoe.GetAllShoesResponse
}
var file_shoe_shoe_proto_depIdxs = []int32{
	0,  // 0: shoe.GetShoeResponse.shoe:type_name -> shoe.Shoe
	0,  // 1: shoe.UpdateShoeResponse.shoe:type_name -> shoe.Shoe
	0,  // 2: shoe.GetAllShoesResponse.shoes:type_name -> shoe.Shoe
	1,  // 3: shoe.ShoeService.AddShoe:input_type -> shoe.AddShoeRequest
	3,  // 4: shoe.ShoeService.GetShoe:input_type -> shoe.GetShoeRequest
	5,  // 5: shoe.ShoeService.DeleteShoe:input_type -> shoe.DeleteShoeRequest
	7,  // 6: shoe.ShoeService.UpdateShoe:input_type -> shoe.UpdateShoeRequest
	9,  // 7: shoe.ShoeService.GetShoes:input_type -> shoe.GetAllShoesRequest
	2,  // 8: shoe.ShoeService.AddShoe:output_type -> shoe.AddShoeResponse
	4,  // 9: shoe.ShoeService.GetShoe:output_type -> shoe.GetShoeResponse
	6,  // 10: shoe.ShoeService.DeleteShoe:output_type -> shoe.DeleteShoeResponse
	8,  // 11: shoe.ShoeService.UpdateShoe:output_type -> shoe.UpdateShoeResponse
	10, // 12: shoe.ShoeService.GetShoes:output_type -> shoe.GetAllShoesResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_shoe_shoe_proto_init() }
func file_shoe_shoe_proto_init() {
	if File_shoe_shoe_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shoe_shoe_proto_rawDesc), len(file_shoe_shoe_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shoe_shoe_proto_goTypes,
		DependencyIndexes: file_shoe_shoe_proto_depIdxs,
		MessageInfos:      file_shoe_shoe_proto_msgTypes,
	}.Build()
	File_shoe_shoe_proto = out.File
	file_shoe_shoe_proto_goTypes = nil
	file_shoe_shoe_proto_depIdxs = nil
}
