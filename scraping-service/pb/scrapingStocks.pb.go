// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scrapingStocks.proto

package pb

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

type SummaryStatus int32

const (
	SummaryStatus_ACTIVE   SummaryStatus = 0
	SummaryStatus_INACTIVE SummaryStatus = 1
)

var SummaryStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
}

var SummaryStatus_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
}

func (x SummaryStatus) String() string {
	return proto.EnumName(SummaryStatus_name, int32(x))
}

func (SummaryStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_edfbfe4dc42693ea, []int{0}
}

type StockCode struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	National             bool     `protobuf:"varint,2,opt,name=national,proto3" json:"national,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockCode) Reset()         { *m = StockCode{} }
func (m *StockCode) String() string { return proto.CompactTextString(m) }
func (*StockCode) ProtoMessage()    {}
func (*StockCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfbfe4dc42693ea, []int{0}
}

func (m *StockCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockCode.Unmarshal(m, b)
}
func (m *StockCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockCode.Marshal(b, m, deterministic)
}
func (m *StockCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockCode.Merge(m, src)
}
func (m *StockCode) XXX_Size() int {
	return xxx_messageInfo_StockCode.Size(m)
}
func (m *StockCode) XXX_DiscardUnknown() {
	xxx_messageInfo_StockCode.DiscardUnknown(m)
}

var xxx_messageInfo_StockCode proto.InternalMessageInfo

func (m *StockCode) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *StockCode) GetNational() bool {
	if m != nil {
		return m.National
	}
	return false
}

type SummaryStock struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName          string             `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	CompanyCode          string             `protobuf:"bytes,3,opt,name=company_code,json=companyCode,proto3" json:"company_code,omitempty"`
	StockValue           *SumarryStockValue `protobuf:"bytes,4,opt,name=stock_value,json=stockValue,proto3" json:"stock_value,omitempty"`
	Summary              *Summary           `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Status               SummaryStatus      `protobuf:"varint,6,opt,name=status,proto3,enum=SummaryStatus" json:"status,omitempty"`
	BasicNosql           *BasicNoSQL        `protobuf:"bytes,7,opt,name=basic_nosql,json=basicNosql,proto3" json:"basic_nosql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SummaryStock) Reset()         { *m = SummaryStock{} }
func (m *SummaryStock) String() string { return proto.CompactTextString(m) }
func (*SummaryStock) ProtoMessage()    {}
func (*SummaryStock) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfbfe4dc42693ea, []int{1}
}

func (m *SummaryStock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryStock.Unmarshal(m, b)
}
func (m *SummaryStock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryStock.Marshal(b, m, deterministic)
}
func (m *SummaryStock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryStock.Merge(m, src)
}
func (m *SummaryStock) XXX_Size() int {
	return xxx_messageInfo_SummaryStock.Size(m)
}
func (m *SummaryStock) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryStock.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryStock proto.InternalMessageInfo

func (m *SummaryStock) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SummaryStock) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *SummaryStock) GetCompanyCode() string {
	if m != nil {
		return m.CompanyCode
	}
	return ""
}

func (m *SummaryStock) GetStockValue() *SumarryStockValue {
	if m != nil {
		return m.StockValue
	}
	return nil
}

func (m *SummaryStock) GetSummary() *Summary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *SummaryStock) GetStatus() SummaryStatus {
	if m != nil {
		return m.Status
	}
	return SummaryStatus_ACTIVE
}

func (m *SummaryStock) GetBasicNosql() *BasicNoSQL {
	if m != nil {
		return m.BasicNosql
	}
	return nil
}

type SumarryStockValue struct {
	Price                float64  `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	RangeDay             float64  `protobuf:"fixed64,2,opt,name=range_day,json=rangeDay,proto3" json:"range_day,omitempty"`
	PercentRange         float64  `protobuf:"fixed64,3,opt,name=percent_range,json=percentRange,proto3" json:"percent_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumarryStockValue) Reset()         { *m = SumarryStockValue{} }
func (m *SumarryStockValue) String() string { return proto.CompactTextString(m) }
func (*SumarryStockValue) ProtoMessage()    {}
func (*SumarryStockValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfbfe4dc42693ea, []int{2}
}

func (m *SumarryStockValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumarryStockValue.Unmarshal(m, b)
}
func (m *SumarryStockValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumarryStockValue.Marshal(b, m, deterministic)
}
func (m *SumarryStockValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumarryStockValue.Merge(m, src)
}
func (m *SumarryStockValue) XXX_Size() int {
	return xxx_messageInfo_SumarryStockValue.Size(m)
}
func (m *SumarryStockValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SumarryStockValue.DiscardUnknown(m)
}

var xxx_messageInfo_SumarryStockValue proto.InternalMessageInfo

func (m *SumarryStockValue) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SumarryStockValue) GetRangeDay() float64 {
	if m != nil {
		return m.RangeDay
	}
	return 0
}

func (m *SumarryStockValue) GetPercentRange() float64 {
	if m != nil {
		return m.PercentRange
	}
	return 0
}

type Summary struct {
	PreviousClose        float64     `protobuf:"fixed64,1,opt,name=previous_close,json=previousClose,proto3" json:"previous_close,omitempty"`
	Open                 float64     `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	DayRange             *RangePrice `protobuf:"bytes,3,opt,name=day_range,json=dayRange,proto3" json:"day_range,omitempty"`
	YearRange            *RangePrice `protobuf:"bytes,4,opt,name=year_range,json=yearRange,proto3" json:"year_range,omitempty"`
	Volume               uint64      `protobuf:"varint,5,opt,name=volume,proto3" json:"volume,omitempty"`
	AvgVol               uint64      `protobuf:"varint,6,opt,name=avg_vol,json=avgVol,proto3" json:"avg_vol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfbfe4dc42693ea, []int{3}
}

func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (m *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(m, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetPreviousClose() float64 {
	if m != nil {
		return m.PreviousClose
	}
	return 0
}

func (m *Summary) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Summary) GetDayRange() *RangePrice {
	if m != nil {
		return m.DayRange
	}
	return nil
}

func (m *Summary) GetYearRange() *RangePrice {
	if m != nil {
		return m.YearRange
	}
	return nil
}

func (m *Summary) GetVolume() uint64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Summary) GetAvgVol() uint64 {
	if m != nil {
		return m.AvgVol
	}
	return 0
}

type RangePrice struct {
	Start                float64  `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RangePrice) Reset()         { *m = RangePrice{} }
func (m *RangePrice) String() string { return proto.CompactTextString(m) }
func (*RangePrice) ProtoMessage()    {}
func (*RangePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfbfe4dc42693ea, []int{4}
}

func (m *RangePrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RangePrice.Unmarshal(m, b)
}
func (m *RangePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RangePrice.Marshal(b, m, deterministic)
}
func (m *RangePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangePrice.Merge(m, src)
}
func (m *RangePrice) XXX_Size() int {
	return xxx_messageInfo_RangePrice.Size(m)
}
func (m *RangePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_RangePrice.DiscardUnknown(m)
}

var xxx_messageInfo_RangePrice proto.InternalMessageInfo

func (m *RangePrice) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *RangePrice) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

type BasicNoSQL struct {
	CreatedAt            int64    `protobuf:"varint,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasicNoSQL) Reset()         { *m = BasicNoSQL{} }
func (m *BasicNoSQL) String() string { return proto.CompactTextString(m) }
func (*BasicNoSQL) ProtoMessage()    {}
func (*BasicNoSQL) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfbfe4dc42693ea, []int{5}
}

func (m *BasicNoSQL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasicNoSQL.Unmarshal(m, b)
}
func (m *BasicNoSQL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasicNoSQL.Marshal(b, m, deterministic)
}
func (m *BasicNoSQL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicNoSQL.Merge(m, src)
}
func (m *BasicNoSQL) XXX_Size() int {
	return xxx_messageInfo_BasicNoSQL.Size(m)
}
func (m *BasicNoSQL) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicNoSQL.DiscardUnknown(m)
}

var xxx_messageInfo_BasicNoSQL proto.InternalMessageInfo

func (m *BasicNoSQL) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *BasicNoSQL) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("SummaryStatus", SummaryStatus_name, SummaryStatus_value)
	proto.RegisterType((*StockCode)(nil), "StockCode")
	proto.RegisterType((*SummaryStock)(nil), "SummaryStock")
	proto.RegisterType((*SumarryStockValue)(nil), "SumarryStockValue")
	proto.RegisterType((*Summary)(nil), "Summary")
	proto.RegisterType((*RangePrice)(nil), "RangePrice")
	proto.RegisterType((*BasicNoSQL)(nil), "BasicNoSQL")
}

func init() {
	proto.RegisterFile("scrapingStocks.proto", fileDescriptor_edfbfe4dc42693ea)
}

var fileDescriptor_edfbfe4dc42693ea = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0x5d, 0x97, 0x26, 0xb7, 0x1f, 0x0c, 0x6b, 0x82, 0x68, 0x08, 0xa9, 0x04, 0x01, 0x65,
	0x62, 0xa9, 0xb4, 0x21, 0x84, 0xe0, 0x69, 0x2b, 0x08, 0x0d, 0xa1, 0x09, 0x5c, 0xd4, 0x07, 0x5e,
	0x22, 0x27, 0xb1, 0x42, 0xb4, 0x24, 0xce, 0x6c, 0x27, 0x52, 0x7e, 0x07, 0x3f, 0x8b, 0x3f, 0x85,
	0xec, 0x38, 0x5d, 0xa7, 0xbd, 0xdd, 0x73, 0xce, 0xf5, 0xbd, 0xc7, 0xd7, 0xbe, 0x70, 0x28, 0x62,
	0x4e, 0xaa, 0xac, 0x4c, 0xd7, 0x92, 0xc5, 0xd7, 0x22, 0xa8, 0x38, 0x93, 0xcc, 0xff, 0x04, 0xae,
	0xc6, 0x2b, 0x96, 0x50, 0x84, 0x60, 0x18, 0xb3, 0x84, 0x7a, 0xd6, 0xdc, 0x5a, 0xb8, 0x58, 0xc7,
	0xe8, 0x08, 0x9c, 0x92, 0xc8, 0x8c, 0x95, 0x24, 0xf7, 0x06, 0x73, 0x6b, 0xe1, 0xe0, 0x2d, 0xf6,
	0xff, 0x0e, 0x60, 0xb2, 0xae, 0x8b, 0x82, 0xf0, 0x56, 0x17, 0x41, 0x33, 0x18, 0x64, 0x89, 0x39,
	0x3e, 0xc8, 0x12, 0xf4, 0x1c, 0x26, 0x31, 0x2b, 0x2a, 0x52, 0xb6, 0x61, 0x49, 0x0a, 0xaa, 0x0b,
	0xb8, 0x78, 0x6c, 0xb8, 0x2b, 0x52, 0xd0, 0xdd, 0x14, 0xdd, 0x7b, 0xef, 0x4e, 0x8a, 0xb6, 0x75,
	0x06, 0x63, 0xa1, 0xca, 0x87, 0x0d, 0xc9, 0x6b, 0xea, 0x0d, 0xe7, 0xd6, 0x62, 0x7c, 0x8a, 0x82,
	0x75, 0x5d, 0x10, 0x6e, 0x3a, 0x6f, 0x94, 0x82, 0x41, 0x6c, 0x63, 0xe4, 0xc3, 0x48, 0x74, 0xd6,
	0xbc, 0x7d, 0x7d, 0xc0, 0x09, 0x8c, 0x55, 0xdc, 0x0b, 0xe8, 0x15, 0xd8, 0x42, 0x12, 0x59, 0x0b,
	0xcf, 0x9e, 0x5b, 0x8b, 0xd9, 0xe9, 0x2c, 0xd8, 0xde, 0x46, 0xb1, 0xd8, 0xa8, 0xe8, 0x2d, 0x8c,
	0x23, 0x22, 0xb2, 0x38, 0x2c, 0x99, 0xb8, 0xc9, 0xbd, 0x91, 0xae, 0x37, 0x0e, 0x2e, 0x14, 0x77,
	0xc5, 0xd6, 0x3f, 0xbf, 0x63, 0x88, 0xba, 0x58, 0xdc, 0xe4, 0xfe, 0x35, 0x3c, 0xba, 0x67, 0x0d,
	0x1d, 0xc2, 0x7e, 0xc5, 0xb3, 0xb8, 0x9b, 0xad, 0x85, 0x3b, 0x80, 0x9e, 0x82, 0xcb, 0x49, 0x99,
	0xd2, 0x30, 0x21, 0xad, 0x1e, 0x8e, 0x85, 0x1d, 0x4d, 0x7c, 0x26, 0x2d, 0x7a, 0x01, 0xd3, 0x8a,
	0xf2, 0x98, 0x96, 0x32, 0xd4, 0x9c, 0x1e, 0x8d, 0x85, 0x27, 0x86, 0xc4, 0x8a, 0xf3, 0xff, 0x59,
	0x30, 0x32, 0xa6, 0xd1, 0x4b, 0x98, 0x55, 0x9c, 0x36, 0x19, 0xab, 0x45, 0x18, 0xe7, 0x4c, 0xf4,
	0xcd, 0xa6, 0x3d, 0xbb, 0x52, 0xa4, 0x7a, 0x65, 0x56, 0xd1, 0xd2, 0xf4, 0xd3, 0x31, 0x5a, 0x80,
	0x9b, 0x90, 0x76, 0xa7, 0x8f, 0xba, 0x9f, 0xee, 0xf0, 0x43, 0x19, 0xc5, 0x4e, 0x42, 0x5a, 0x0d,
	0xd1, 0x31, 0x40, 0x4b, 0x09, 0x37, 0xa9, 0xc3, 0xfb, 0xa9, 0xae, 0x92, 0xbb, 0xdc, 0xc7, 0x60,
	0x37, 0x2c, 0xaf, 0x0b, 0xaa, 0x9f, 0x60, 0x88, 0x0d, 0x42, 0x4f, 0x60, 0x44, 0x9a, 0x34, 0x6c,
	0x58, 0xae, 0x07, 0x3f, 0xc4, 0x36, 0x69, 0xd2, 0x0d, 0xcb, 0xfd, 0x77, 0x00, 0xb7, 0x95, 0xd4,
	0xcc, 0x84, 0x24, 0x5c, 0xf6, 0x33, 0xd3, 0x00, 0x1d, 0xc0, 0x1e, 0x2d, 0x13, 0xe3, 0x5e, 0x85,
	0xfe, 0x37, 0x80, 0xdb, 0xa7, 0x40, 0xcf, 0x00, 0x62, 0x4e, 0x89, 0xa4, 0x49, 0x48, 0xba, 0xa3,
	0x7b, 0xd8, 0x35, 0xcc, 0xb9, 0x54, 0x72, 0x5d, 0x25, 0xbd, 0x3c, 0xe8, 0x64, 0xc3, 0x9c, 0xcb,
	0xe3, 0x37, 0x30, 0xbd, 0xf3, 0x07, 0x10, 0x80, 0x7d, 0xbe, 0xfa, 0x75, 0xb9, 0xf9, 0x72, 0xf0,
	0x00, 0x4d, 0xc0, 0xb9, 0xbc, 0x32, 0xc8, 0x3a, 0xfd, 0x08, 0x0f, 0xd7, 0xfd, 0x4a, 0x51, 0xde,
	0x28, 0xc7, 0xaf, 0x01, 0xbe, 0x52, 0xb9, 0xea, 0xfe, 0x2e, 0x82, 0x60, 0xbb, 0x5a, 0x47, 0xd3,
	0x60, 0x77, 0x51, 0x2e, 0x3e, 0xfc, 0x7e, 0x9f, 0x66, 0xf2, 0x4f, 0x1d, 0x05, 0x31, 0x2b, 0x96,
	0x3c, 0x8b, 0x68, 0xc6, 0x99, 0x20, 0x59, 0xc1, 0xca, 0x65, 0xc1, 0xd4, 0x7a, 0x9d, 0xa4, 0x6c,
	0xd9, 0x6f, 0xec, 0x89, 0xe8, 0xea, 0x2f, 0xab, 0x28, 0xb2, 0xf5, 0xde, 0x9e, 0xfd, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x9f, 0x24, 0xca, 0x07, 0xcf, 0x03, 0x00, 0x00,
}
