// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: entities/variant.proto

package pb

import (
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

type Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code            string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name            string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Barcode         string  `protobuf:"bytes,4,opt,name=barcode,proto3" json:"barcode,omitempty"`
	DecisionNumber  string  `protobuf:"bytes,5,opt,name=decision_number,json=decisionNumber,proto3" json:"decision_number,omitempty"`
	RegisterNumber  string  `protobuf:"bytes,6,opt,name=register_number,json=registerNumber,proto3" json:"register_number,omitempty"`
	Longevity       string  `protobuf:"bytes,7,opt,name=longevity,proto3" json:"longevity,omitempty"`
	Vat             float32 `protobuf:"fixed32,8,opt,name=vat,proto3" json:"vat,omitempty"`
	Product         int32   `protobuf:"varint,9,opt,name=product,proto3" json:"product,omitempty"`
	Media           string  `protobuf:"bytes,10,opt,name=media,proto3" json:"media,omitempty"`
	QuantityInStock *int32  `protobuf:"varint,11,opt,name=quantity_in_stock,json=quantityInStock,proto3,oneof" json:"quantity_in_stock,omitempty"`
	Units           []*Unit `protobuf:"bytes,12,rep,name=units,proto3" json:"units,omitempty"`
	PriceSell       float32 `protobuf:"fixed32,13,opt,name=price_sell,json=priceSell,proto3" json:"price_sell,omitempty"`
	PriceImport     float32 `protobuf:"fixed32,14,opt,name=price_import,json=priceImport,proto3" json:"price_import,omitempty"`
}

func (x *Variant) Reset() {
	*x = Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_variant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_entities_variant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_entities_variant_proto_rawDescGZIP(), []int{0}
}

func (x *Variant) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Variant) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Variant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Variant) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *Variant) GetDecisionNumber() string {
	if x != nil {
		return x.DecisionNumber
	}
	return ""
}

func (x *Variant) GetRegisterNumber() string {
	if x != nil {
		return x.RegisterNumber
	}
	return ""
}

func (x *Variant) GetLongevity() string {
	if x != nil {
		return x.Longevity
	}
	return ""
}

func (x *Variant) GetVat() float32 {
	if x != nil {
		return x.Vat
	}
	return 0
}

func (x *Variant) GetProduct() int32 {
	if x != nil {
		return x.Product
	}
	return 0
}

func (x *Variant) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

func (x *Variant) GetQuantityInStock() int32 {
	if x != nil && x.QuantityInStock != nil {
		return *x.QuantityInStock
	}
	return 0
}

func (x *Variant) GetUnits() []*Unit {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *Variant) GetPriceSell() float32 {
	if x != nil {
		return x.PriceSell
	}
	return 0
}

func (x *Variant) GetPriceImport() float32 {
	if x != nil {
		return x.PriceImport
	}
	return 0
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value       int32   `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	SellPrice   float32 `protobuf:"fixed32,4,opt,name=sell_price,json=sellPrice,proto3" json:"sell_price,omitempty"`
	ImportPrice float32 `protobuf:"fixed32,5,opt,name=import_price,json=importPrice,proto3" json:"import_price,omitempty"`
	Weight      float32 `protobuf:"fixed32,6,opt,name=weight,proto3" json:"weight,omitempty"`
	WeightUnit  string  `protobuf:"bytes,7,opt,name=weight_unit,json=weightUnit,proto3" json:"weight_unit,omitempty"`
	Default     bool    `protobuf:"varint,8,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_variant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_entities_variant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_entities_variant_proto_rawDescGZIP(), []int{1}
}

func (x *Unit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Unit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Unit) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Unit) GetSellPrice() float32 {
	if x != nil {
		return x.SellPrice
	}
	return 0
}

func (x *Unit) GetImportPrice() float32 {
	if x != nil {
		return x.ImportPrice
	}
	return 0
}

func (x *Unit) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Unit) GetWeightUnit() string {
	if x != nil {
		return x.WeightUnit
	}
	return ""
}

func (x *Unit) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

var File_entities_variant_proto protoreflect.FileDescriptor

var file_entities_variant_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x16, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03, 0x0a,
	0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x76, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x76, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x76, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x2f, 0x0a, 0x11,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x42,
	0x14, 0x0a, 0x12, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x5f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0xd5, 0x01, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x65,
	0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e,
	0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61,
	0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_variant_proto_rawDescOnce sync.Once
	file_entities_variant_proto_rawDescData = file_entities_variant_proto_rawDesc
)

func file_entities_variant_proto_rawDescGZIP() []byte {
	file_entities_variant_proto_rawDescOnce.Do(func() {
		file_entities_variant_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_variant_proto_rawDescData)
	})
	return file_entities_variant_proto_rawDescData
}

var file_entities_variant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entities_variant_proto_goTypes = []interface{}{
	(*Variant)(nil), // 0: pb.Variant
	(*Unit)(nil),    // 1: pb.Unit
}
var file_entities_variant_proto_depIdxs = []int32{
	1, // 0: pb.Variant.units:type_name -> pb.Unit
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entities_variant_proto_init() }
func file_entities_variant_proto_init() {
	if File_entities_variant_proto != nil {
		return
	}
	file_entities_address_proto_init()
	file_entities_company_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_variant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant); i {
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
		file_entities_variant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
	file_entities_variant_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entities_variant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_variant_proto_goTypes,
		DependencyIndexes: file_entities_variant_proto_depIdxs,
		MessageInfos:      file_entities_variant_proto_msgTypes,
	}.Build()
	File_entities_variant_proto = out.File
	file_entities_variant_proto_rawDesc = nil
	file_entities_variant_proto_goTypes = nil
	file_entities_variant_proto_depIdxs = nil
}
