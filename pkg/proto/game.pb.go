// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package proto

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

type GameObject struct {
	ID                   string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Tags                 []*TagObject  `protobuf:"bytes,2,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Description          string        `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Name                 string        `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Title                string        `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	Developer            *LinkObject   `protobuf:"bytes,6,opt,name=Developer,proto3" json:"Developer,omitempty"`
	ReleaseDate          string        `protobuf:"bytes,7,opt,name=ReleaseDate,proto3" json:"ReleaseDate,omitempty"`
	GenreMain            *TagObject    `protobuf:"bytes,8,opt,name=GenreMain,proto3" json:"GenreMain,omitempty"`
	Genres               []*TagObject  `protobuf:"bytes,9,rep,name=Genres,proto3" json:"Genres,omitempty"`
	Platforms            *Platforms    `protobuf:"bytes,10,opt,name=Platforms,proto3" json:"Platforms,omitempty"`
	Requirements         *Requirements `protobuf:"bytes,11,opt,name=Requirements,proto3" json:"Requirements,omitempty"`
	Languages            *Languages    `protobuf:"bytes,12,opt,name=Languages,proto3" json:"Languages,omitempty"`
	DisplayRemainingTime bool          `protobuf:"varint,13,opt,name=DisplayRemainingTime,proto3" json:"DisplayRemainingTime,omitempty"`
	FeaturesControl      string        `protobuf:"bytes,14,opt,name=FeaturesControl,proto3" json:"FeaturesControl,omitempty"`
	Features             []string      `protobuf:"bytes,15,rep,name=Features,proto3" json:"Features,omitempty"`
	Media                *Media        `protobuf:"bytes,16,opt,name=Media,proto3" json:"Media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GameObject) Reset()         { *m = GameObject{} }
func (m *GameObject) String() string { return proto.CompactTextString(m) }
func (*GameObject) ProtoMessage()    {}
func (*GameObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *GameObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameObject.Unmarshal(m, b)
}
func (m *GameObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameObject.Marshal(b, m, deterministic)
}
func (m *GameObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameObject.Merge(m, src)
}
func (m *GameObject) XXX_Size() int {
	return xxx_messageInfo_GameObject.Size(m)
}
func (m *GameObject) XXX_DiscardUnknown() {
	xxx_messageInfo_GameObject.DiscardUnknown(m)
}

var xxx_messageInfo_GameObject proto.InternalMessageInfo

func (m *GameObject) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GameObject) GetTags() []*TagObject {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *GameObject) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GameObject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameObject) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GameObject) GetDeveloper() *LinkObject {
	if m != nil {
		return m.Developer
	}
	return nil
}

func (m *GameObject) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

func (m *GameObject) GetGenreMain() *TagObject {
	if m != nil {
		return m.GenreMain
	}
	return nil
}

func (m *GameObject) GetGenres() []*TagObject {
	if m != nil {
		return m.Genres
	}
	return nil
}

func (m *GameObject) GetPlatforms() *Platforms {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *GameObject) GetRequirements() *Requirements {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *GameObject) GetLanguages() *Languages {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *GameObject) GetDisplayRemainingTime() bool {
	if m != nil {
		return m.DisplayRemainingTime
	}
	return false
}

func (m *GameObject) GetFeaturesControl() string {
	if m != nil {
		return m.FeaturesControl
	}
	return ""
}

func (m *GameObject) GetFeatures() []string {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *GameObject) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

type Media struct {
	CoverImage           *LocalizedString        `protobuf:"bytes,1,opt,name=CoverImage,proto3" json:"CoverImage,omitempty"`
	CoverVideo           *LocalizedString        `protobuf:"bytes,2,opt,name=CoverVideo,proto3" json:"CoverVideo,omitempty"`
	Trailers             []*LocalizedStringArray `protobuf:"bytes,3,rep,name=Trailers,proto3" json:"Trailers,omitempty"`
	Screenshots          []*LocalizedStringArray `protobuf:"bytes,4,rep,name=Screenshots,proto3" json:"Screenshots,omitempty"`
	Special              *LocalizedString        `protobuf:"bytes,5,opt,name=Special,proto3" json:"Special,omitempty"`
	Friends              *LocalizedString        `protobuf:"bytes,6,opt,name=Friends,proto3" json:"Friends,omitempty"`
	CapsuleGeneric       *LocalizedString        `protobuf:"bytes,7,opt,name=CapsuleGeneric,proto3" json:"CapsuleGeneric,omitempty"`
	CapsuleSmall         *LocalizedString        `protobuf:"bytes,8,opt,name=CapsuleSmall,proto3" json:"CapsuleSmall,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}
func (*Media) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *Media) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Media.Unmarshal(m, b)
}
func (m *Media) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Media.Marshal(b, m, deterministic)
}
func (m *Media) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Media.Merge(m, src)
}
func (m *Media) XXX_Size() int {
	return xxx_messageInfo_Media.Size(m)
}
func (m *Media) XXX_DiscardUnknown() {
	xxx_messageInfo_Media.DiscardUnknown(m)
}

var xxx_messageInfo_Media proto.InternalMessageInfo

func (m *Media) GetCoverImage() *LocalizedString {
	if m != nil {
		return m.CoverImage
	}
	return nil
}

func (m *Media) GetCoverVideo() *LocalizedString {
	if m != nil {
		return m.CoverVideo
	}
	return nil
}

func (m *Media) GetTrailers() []*LocalizedStringArray {
	if m != nil {
		return m.Trailers
	}
	return nil
}

func (m *Media) GetScreenshots() []*LocalizedStringArray {
	if m != nil {
		return m.Screenshots
	}
	return nil
}

func (m *Media) GetSpecial() *LocalizedString {
	if m != nil {
		return m.Special
	}
	return nil
}

func (m *Media) GetFriends() *LocalizedString {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *Media) GetCapsuleGeneric() *LocalizedString {
	if m != nil {
		return m.CapsuleGeneric
	}
	return nil
}

func (m *Media) GetCapsuleSmall() *LocalizedString {
	if m != nil {
		return m.CapsuleSmall
	}
	return nil
}

type LocalizedStringArray struct {
	EN                   []*Language `protobuf:"bytes,1,rep,name=EN,proto3" json:"EN,omitempty"`
	RU                   []*Language `protobuf:"bytes,2,rep,name=RU,proto3" json:"RU,omitempty"`
	FR                   []*Language `protobuf:"bytes,3,rep,name=FR,proto3" json:"FR,omitempty"`
	ES                   []*Language `protobuf:"bytes,4,rep,name=ES,proto3" json:"ES,omitempty"`
	DE                   []*Language `protobuf:"bytes,5,rep,name=DE,proto3" json:"DE,omitempty"`
	IT                   []*Language `protobuf:"bytes,6,rep,name=IT,proto3" json:"IT,omitempty"`
	PT                   []*Language `protobuf:"bytes,7,rep,name=PT,proto3" json:"PT,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LocalizedStringArray) Reset()         { *m = LocalizedStringArray{} }
func (m *LocalizedStringArray) String() string { return proto.CompactTextString(m) }
func (*LocalizedStringArray) ProtoMessage()    {}
func (*LocalizedStringArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *LocalizedStringArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalizedStringArray.Unmarshal(m, b)
}
func (m *LocalizedStringArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalizedStringArray.Marshal(b, m, deterministic)
}
func (m *LocalizedStringArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalizedStringArray.Merge(m, src)
}
func (m *LocalizedStringArray) XXX_Size() int {
	return xxx_messageInfo_LocalizedStringArray.Size(m)
}
func (m *LocalizedStringArray) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalizedStringArray.DiscardUnknown(m)
}

var xxx_messageInfo_LocalizedStringArray proto.InternalMessageInfo

func (m *LocalizedStringArray) GetEN() []*Language {
	if m != nil {
		return m.EN
	}
	return nil
}

func (m *LocalizedStringArray) GetRU() []*Language {
	if m != nil {
		return m.RU
	}
	return nil
}

func (m *LocalizedStringArray) GetFR() []*Language {
	if m != nil {
		return m.FR
	}
	return nil
}

func (m *LocalizedStringArray) GetES() []*Language {
	if m != nil {
		return m.ES
	}
	return nil
}

func (m *LocalizedStringArray) GetDE() []*Language {
	if m != nil {
		return m.DE
	}
	return nil
}

func (m *LocalizedStringArray) GetIT() []*Language {
	if m != nil {
		return m.IT
	}
	return nil
}

func (m *LocalizedStringArray) GetPT() []*Language {
	if m != nil {
		return m.PT
	}
	return nil
}

type Languages struct {
	EN                   *Language `protobuf:"bytes,1,opt,name=EN,proto3" json:"EN,omitempty"`
	RU                   *Language `protobuf:"bytes,2,opt,name=RU,proto3" json:"RU,omitempty"`
	FR                   *Language `protobuf:"bytes,3,opt,name=FR,proto3" json:"FR,omitempty"`
	ES                   *Language `protobuf:"bytes,4,opt,name=ES,proto3" json:"ES,omitempty"`
	DE                   *Language `protobuf:"bytes,5,opt,name=DE,proto3" json:"DE,omitempty"`
	IT                   *Language `protobuf:"bytes,6,opt,name=IT,proto3" json:"IT,omitempty"`
	PT                   *Language `protobuf:"bytes,7,opt,name=PT,proto3" json:"PT,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Languages) Reset()         { *m = Languages{} }
func (m *Languages) String() string { return proto.CompactTextString(m) }
func (*Languages) ProtoMessage()    {}
func (*Languages) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *Languages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Languages.Unmarshal(m, b)
}
func (m *Languages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Languages.Marshal(b, m, deterministic)
}
func (m *Languages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Languages.Merge(m, src)
}
func (m *Languages) XXX_Size() int {
	return xxx_messageInfo_Languages.Size(m)
}
func (m *Languages) XXX_DiscardUnknown() {
	xxx_messageInfo_Languages.DiscardUnknown(m)
}

var xxx_messageInfo_Languages proto.InternalMessageInfo

func (m *Languages) GetEN() *Language {
	if m != nil {
		return m.EN
	}
	return nil
}

func (m *Languages) GetRU() *Language {
	if m != nil {
		return m.RU
	}
	return nil
}

func (m *Languages) GetFR() *Language {
	if m != nil {
		return m.FR
	}
	return nil
}

func (m *Languages) GetES() *Language {
	if m != nil {
		return m.ES
	}
	return nil
}

func (m *Languages) GetDE() *Language {
	if m != nil {
		return m.DE
	}
	return nil
}

func (m *Languages) GetIT() *Language {
	if m != nil {
		return m.IT
	}
	return nil
}

func (m *Languages) GetPT() *Language {
	if m != nil {
		return m.PT
	}
	return nil
}

type Language struct {
	Voice                bool     `protobuf:"varint,1,opt,name=Voice,proto3" json:"Voice,omitempty"`
	Interface            bool     `protobuf:"varint,2,opt,name=Interface,proto3" json:"Interface,omitempty"`
	Subtitles            bool     `protobuf:"varint,3,opt,name=Subtitles,proto3" json:"Subtitles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Language) Reset()         { *m = Language{} }
func (m *Language) String() string { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()    {}
func (*Language) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4}
}

func (m *Language) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Language.Unmarshal(m, b)
}
func (m *Language) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Language.Marshal(b, m, deterministic)
}
func (m *Language) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Language.Merge(m, src)
}
func (m *Language) XXX_Size() int {
	return xxx_messageInfo_Language.Size(m)
}
func (m *Language) XXX_DiscardUnknown() {
	xxx_messageInfo_Language.DiscardUnknown(m)
}

var xxx_messageInfo_Language proto.InternalMessageInfo

func (m *Language) GetVoice() bool {
	if m != nil {
		return m.Voice
	}
	return false
}

func (m *Language) GetInterface() bool {
	if m != nil {
		return m.Interface
	}
	return false
}

func (m *Language) GetSubtitles() bool {
	if m != nil {
		return m.Subtitles
	}
	return false
}

type Requirements struct {
	Windows              *PlatformRequirements `protobuf:"bytes,1,opt,name=Windows,proto3" json:"Windows,omitempty"`
	MacOs                *PlatformRequirements `protobuf:"bytes,2,opt,name=MacOs,proto3" json:"MacOs,omitempty"`
	Linux                *PlatformRequirements `protobuf:"bytes,3,opt,name=Linux,proto3" json:"Linux,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Requirements) Reset()         { *m = Requirements{} }
func (m *Requirements) String() string { return proto.CompactTextString(m) }
func (*Requirements) ProtoMessage()    {}
func (*Requirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{5}
}

func (m *Requirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Requirements.Unmarshal(m, b)
}
func (m *Requirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Requirements.Marshal(b, m, deterministic)
}
func (m *Requirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Requirements.Merge(m, src)
}
func (m *Requirements) XXX_Size() int {
	return xxx_messageInfo_Requirements.Size(m)
}
func (m *Requirements) XXX_DiscardUnknown() {
	xxx_messageInfo_Requirements.DiscardUnknown(m)
}

var xxx_messageInfo_Requirements proto.InternalMessageInfo

func (m *Requirements) GetWindows() *PlatformRequirements {
	if m != nil {
		return m.Windows
	}
	return nil
}

func (m *Requirements) GetMacOs() *PlatformRequirements {
	if m != nil {
		return m.MacOs
	}
	return nil
}

func (m *Requirements) GetLinux() *PlatformRequirements {
	if m != nil {
		return m.Linux
	}
	return nil
}

type PlatformRequirements struct {
	Minimal              *MachineRequirements `protobuf:"bytes,1,opt,name=Minimal,proto3" json:"Minimal,omitempty"`
	Recommended          *MachineRequirements `protobuf:"bytes,2,opt,name=Recommended,proto3" json:"Recommended,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PlatformRequirements) Reset()         { *m = PlatformRequirements{} }
func (m *PlatformRequirements) String() string { return proto.CompactTextString(m) }
func (*PlatformRequirements) ProtoMessage()    {}
func (*PlatformRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{6}
}

func (m *PlatformRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformRequirements.Unmarshal(m, b)
}
func (m *PlatformRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformRequirements.Marshal(b, m, deterministic)
}
func (m *PlatformRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformRequirements.Merge(m, src)
}
func (m *PlatformRequirements) XXX_Size() int {
	return xxx_messageInfo_PlatformRequirements.Size(m)
}
func (m *PlatformRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformRequirements proto.InternalMessageInfo

func (m *PlatformRequirements) GetMinimal() *MachineRequirements {
	if m != nil {
		return m.Minimal
	}
	return nil
}

func (m *PlatformRequirements) GetRecommended() *MachineRequirements {
	if m != nil {
		return m.Recommended
	}
	return nil
}

type MachineRequirements struct {
	System               string   `protobuf:"bytes,1,opt,name=System,proto3" json:"System,omitempty"`
	Processor            string   `protobuf:"bytes,2,opt,name=Processor,proto3" json:"Processor,omitempty"`
	Graphics             string   `protobuf:"bytes,3,opt,name=Graphics,proto3" json:"Graphics,omitempty"`
	Sound                string   `protobuf:"bytes,4,opt,name=Sound,proto3" json:"Sound,omitempty"`
	Ram                  int32    `protobuf:"varint,5,opt,name=Ram,proto3" json:"Ram,omitempty"`
	RamDimension         string   `protobuf:"bytes,6,opt,name=RamDimension,proto3" json:"RamDimension,omitempty"`
	Storage              int32    `protobuf:"varint,7,opt,name=Storage,proto3" json:"Storage,omitempty"`
	StorageDimension     string   `protobuf:"bytes,8,opt,name=StorageDimension,proto3" json:"StorageDimension,omitempty"`
	Other                string   `protobuf:"bytes,9,opt,name=Other,proto3" json:"Other,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineRequirements) Reset()         { *m = MachineRequirements{} }
func (m *MachineRequirements) String() string { return proto.CompactTextString(m) }
func (*MachineRequirements) ProtoMessage()    {}
func (*MachineRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{7}
}

func (m *MachineRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineRequirements.Unmarshal(m, b)
}
func (m *MachineRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineRequirements.Marshal(b, m, deterministic)
}
func (m *MachineRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineRequirements.Merge(m, src)
}
func (m *MachineRequirements) XXX_Size() int {
	return xxx_messageInfo_MachineRequirements.Size(m)
}
func (m *MachineRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_MachineRequirements proto.InternalMessageInfo

func (m *MachineRequirements) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *MachineRequirements) GetProcessor() string {
	if m != nil {
		return m.Processor
	}
	return ""
}

func (m *MachineRequirements) GetGraphics() string {
	if m != nil {
		return m.Graphics
	}
	return ""
}

func (m *MachineRequirements) GetSound() string {
	if m != nil {
		return m.Sound
	}
	return ""
}

func (m *MachineRequirements) GetRam() int32 {
	if m != nil {
		return m.Ram
	}
	return 0
}

func (m *MachineRequirements) GetRamDimension() string {
	if m != nil {
		return m.RamDimension
	}
	return ""
}

func (m *MachineRequirements) GetStorage() int32 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *MachineRequirements) GetStorageDimension() string {
	if m != nil {
		return m.StorageDimension
	}
	return ""
}

func (m *MachineRequirements) GetOther() string {
	if m != nil {
		return m.Other
	}
	return ""
}

type Platforms struct {
	Windows              bool     `protobuf:"varint,1,opt,name=Windows,proto3" json:"Windows,omitempty"`
	MacOs                bool     `protobuf:"varint,2,opt,name=MacOs,proto3" json:"MacOs,omitempty"`
	Linux                bool     `protobuf:"varint,3,opt,name=Linux,proto3" json:"Linux,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platforms) Reset()         { *m = Platforms{} }
func (m *Platforms) String() string { return proto.CompactTextString(m) }
func (*Platforms) ProtoMessage()    {}
func (*Platforms) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{8}
}

func (m *Platforms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platforms.Unmarshal(m, b)
}
func (m *Platforms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platforms.Marshal(b, m, deterministic)
}
func (m *Platforms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platforms.Merge(m, src)
}
func (m *Platforms) XXX_Size() int {
	return xxx_messageInfo_Platforms.Size(m)
}
func (m *Platforms) XXX_DiscardUnknown() {
	xxx_messageInfo_Platforms.DiscardUnknown(m)
}

var xxx_messageInfo_Platforms proto.InternalMessageInfo

func (m *Platforms) GetWindows() bool {
	if m != nil {
		return m.Windows
	}
	return false
}

func (m *Platforms) GetMacOs() bool {
	if m != nil {
		return m.MacOs
	}
	return false
}

func (m *Platforms) GetLinux() bool {
	if m != nil {
		return m.Linux
	}
	return false
}

type LinkObject struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkObject) Reset()         { *m = LinkObject{} }
func (m *LinkObject) String() string { return proto.CompactTextString(m) }
func (*LinkObject) ProtoMessage()    {}
func (*LinkObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{9}
}

func (m *LinkObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkObject.Unmarshal(m, b)
}
func (m *LinkObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkObject.Marshal(b, m, deterministic)
}
func (m *LinkObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkObject.Merge(m, src)
}
func (m *LinkObject) XXX_Size() int {
	return xxx_messageInfo_LinkObject.Size(m)
}
func (m *LinkObject) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkObject.DiscardUnknown(m)
}

var xxx_messageInfo_LinkObject proto.InternalMessageInfo

func (m *LinkObject) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *LinkObject) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type TagObject struct {
	ID                   string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 *LocalizedString `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TagObject) Reset()         { *m = TagObject{} }
func (m *TagObject) String() string { return proto.CompactTextString(m) }
func (*TagObject) ProtoMessage()    {}
func (*TagObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{10}
}

func (m *TagObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagObject.Unmarshal(m, b)
}
func (m *TagObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagObject.Marshal(b, m, deterministic)
}
func (m *TagObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagObject.Merge(m, src)
}
func (m *TagObject) XXX_Size() int {
	return xxx_messageInfo_TagObject.Size(m)
}
func (m *TagObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TagObject.DiscardUnknown(m)
}

var xxx_messageInfo_TagObject proto.InternalMessageInfo

func (m *TagObject) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TagObject) GetName() *LocalizedString {
	if m != nil {
		return m.Name
	}
	return nil
}

type LocalizedString struct {
	EN                   string   `protobuf:"bytes,1,opt,name=EN,proto3" json:"EN,omitempty"`
	RU                   string   `protobuf:"bytes,2,opt,name=RU,proto3" json:"RU,omitempty"`
	FR                   string   `protobuf:"bytes,3,opt,name=FR,proto3" json:"FR,omitempty"`
	ES                   string   `protobuf:"bytes,4,opt,name=ES,proto3" json:"ES,omitempty"`
	DE                   string   `protobuf:"bytes,5,opt,name=DE,proto3" json:"DE,omitempty"`
	IT                   string   `protobuf:"bytes,6,opt,name=IT,proto3" json:"IT,omitempty"`
	PT                   string   `protobuf:"bytes,7,opt,name=PT,proto3" json:"PT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalizedString) Reset()         { *m = LocalizedString{} }
func (m *LocalizedString) String() string { return proto.CompactTextString(m) }
func (*LocalizedString) ProtoMessage()    {}
func (*LocalizedString) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{11}
}

func (m *LocalizedString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalizedString.Unmarshal(m, b)
}
func (m *LocalizedString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalizedString.Marshal(b, m, deterministic)
}
func (m *LocalizedString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalizedString.Merge(m, src)
}
func (m *LocalizedString) XXX_Size() int {
	return xxx_messageInfo_LocalizedString.Size(m)
}
func (m *LocalizedString) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalizedString.DiscardUnknown(m)
}

var xxx_messageInfo_LocalizedString proto.InternalMessageInfo

func (m *LocalizedString) GetEN() string {
	if m != nil {
		return m.EN
	}
	return ""
}

func (m *LocalizedString) GetRU() string {
	if m != nil {
		return m.RU
	}
	return ""
}

func (m *LocalizedString) GetFR() string {
	if m != nil {
		return m.FR
	}
	return ""
}

func (m *LocalizedString) GetES() string {
	if m != nil {
		return m.ES
	}
	return ""
}

func (m *LocalizedString) GetDE() string {
	if m != nil {
		return m.DE
	}
	return ""
}

func (m *LocalizedString) GetIT() string {
	if m != nil {
		return m.IT
	}
	return ""
}

func (m *LocalizedString) GetPT() string {
	if m != nil {
		return m.PT
	}
	return ""
}

type GameDeleted struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameDeleted) Reset()         { *m = GameDeleted{} }
func (m *GameDeleted) String() string { return proto.CompactTextString(m) }
func (*GameDeleted) ProtoMessage()    {}
func (*GameDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{12}
}

func (m *GameDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameDeleted.Unmarshal(m, b)
}
func (m *GameDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameDeleted.Marshal(b, m, deterministic)
}
func (m *GameDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameDeleted.Merge(m, src)
}
func (m *GameDeleted) XXX_Size() int {
	return xxx_messageInfo_GameDeleted.Size(m)
}
func (m *GameDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_GameDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_GameDeleted proto.InternalMessageInfo

func (m *GameDeleted) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*GameObject)(nil), "proto.GameObject")
	proto.RegisterType((*Media)(nil), "proto.Media")
	proto.RegisterType((*LocalizedStringArray)(nil), "proto.LocalizedStringArray")
	proto.RegisterType((*Languages)(nil), "proto.Languages")
	proto.RegisterType((*Language)(nil), "proto.Language")
	proto.RegisterType((*Requirements)(nil), "proto.Requirements")
	proto.RegisterType((*PlatformRequirements)(nil), "proto.PlatformRequirements")
	proto.RegisterType((*MachineRequirements)(nil), "proto.MachineRequirements")
	proto.RegisterType((*Platforms)(nil), "proto.Platforms")
	proto.RegisterType((*LinkObject)(nil), "proto.LinkObject")
	proto.RegisterType((*TagObject)(nil), "proto.TagObject")
	proto.RegisterType((*LocalizedString)(nil), "proto.LocalizedString")
	proto.RegisterType((*GameDeleted)(nil), "proto.GameDeleted")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x4d, 0x8f, 0x23, 0x35,
	0x13, 0x56, 0x27, 0x93, 0xa4, 0xbb, 0x32, 0x6f, 0x66, 0x5e, 0x6f, 0xb4, 0xb2, 0x76, 0x41, 0x44,
	0x2d, 0x0e, 0xd1, 0x1e, 0x06, 0x18, 0x3e, 0x56, 0x42, 0x80, 0x84, 0x26, 0x33, 0x51, 0xa4, 0xf9,
	0xc2, 0xc9, 0x2e, 0x17, 0x2e, 0x9e, 0x4e, 0x6d, 0xc6, 0xd0, 0xed, 0x0e, 0x76, 0x67, 0x61, 0xb8,
	0x71, 0xe5, 0xc2, 0x0f, 0xe1, 0xc7, 0xf0, 0x5f, 0xf6, 0x8e, 0x84, 0xec, 0x76, 0xba, 0x3b, 0x5f,
	0x13, 0x4e, 0x49, 0x55, 0x3d, 0x8f, 0xab, 0xaa, 0x5d, 0x7e, 0x6c, 0x80, 0x19, 0x4f, 0xf0, 0x64,
	0xae, 0xd2, 0x2c, 0x25, 0x0d, 0xfb, 0x13, 0xbe, 0x3b, 0x00, 0x18, 0xf2, 0x04, 0x6f, 0xee, 0x7e,
	0xc4, 0x28, 0x23, 0x1d, 0xa8, 0x8d, 0x06, 0xd4, 0xeb, 0x79, 0xfd, 0x80, 0xd5, 0x46, 0x03, 0xf2,
	0x21, 0x1c, 0x4c, 0xf8, 0x4c, 0xd3, 0x5a, 0xaf, 0xde, 0x6f, 0x9f, 0x1e, 0xe7, 0xdc, 0x93, 0x09,
	0x9f, 0xe5, 0x78, 0x66, 0xa3, 0xa4, 0x07, 0xed, 0x01, 0xea, 0x48, 0x89, 0x79, 0x26, 0x52, 0x49,
	0xeb, 0x96, 0x5e, 0x75, 0x11, 0x02, 0x07, 0xd7, 0x3c, 0x41, 0x7a, 0x60, 0x43, 0xf6, 0x3f, 0xe9,
	0x42, 0x63, 0x22, 0xb2, 0x18, 0x69, 0xc3, 0x3a, 0x73, 0x83, 0x7c, 0x04, 0xc1, 0x00, 0xdf, 0x62,
	0x9c, 0xce, 0x51, 0xd1, 0x66, 0xcf, 0xeb, 0xb7, 0x4f, 0xff, 0xef, 0xd2, 0x5e, 0x0a, 0xf9, 0x93,
	0xcb, 0x5b, 0x62, 0x4c, 0x72, 0x86, 0x31, 0x72, 0x8d, 0x03, 0x9e, 0x21, 0x6d, 0xe5, 0xc9, 0x2b,
	0x2e, 0x72, 0x02, 0xc1, 0x10, 0xa5, 0xc2, 0x2b, 0x2e, 0x24, 0xf5, 0xed, 0x92, 0x9b, 0x9d, 0x94,
	0x10, 0xd2, 0x87, 0xa6, 0x35, 0x34, 0x0d, 0x76, 0xb4, 0xed, 0xe2, 0x66, 0xe5, 0xdb, 0x98, 0x67,
	0x6f, 0x52, 0x95, 0x68, 0x0a, 0x2b, 0x2b, 0x17, 0x7e, 0x56, 0x42, 0xc8, 0x4b, 0x38, 0x64, 0xf8,
	0xf3, 0x42, 0x28, 0x4c, 0x50, 0x66, 0x9a, 0xb6, 0x2d, 0xe5, 0x89, 0xa3, 0x54, 0x43, 0x6c, 0x05,
	0x68, 0x12, 0x5d, 0x72, 0x39, 0x5b, 0xf0, 0x19, 0x6a, 0x7a, 0xb8, 0x92, 0xa8, 0xf0, 0xb3, 0x12,
	0x42, 0x4e, 0xa1, 0x3b, 0x10, 0x7a, 0x1e, 0xf3, 0x07, 0x86, 0x09, 0x17, 0x52, 0xc8, 0xd9, 0x44,
	0x24, 0x48, 0xff, 0xd7, 0xf3, 0xfa, 0x3e, 0xdb, 0x1a, 0x23, 0x7d, 0x38, 0xba, 0x40, 0x9e, 0x2d,
	0x14, 0xea, 0xb3, 0x54, 0x66, 0x2a, 0x8d, 0x69, 0xc7, 0x7e, 0xcc, 0x75, 0x37, 0x79, 0x06, 0xfe,
	0xd2, 0x45, 0x8f, 0x7a, 0xf5, 0x7e, 0xc0, 0x0a, 0x9b, 0x84, 0xd0, 0xb8, 0xc2, 0xa9, 0xe0, 0xf4,
	0xd8, 0x56, 0x79, 0xe8, 0xaa, 0xb4, 0x3e, 0x96, 0x87, 0xc2, 0xbf, 0xeb, 0x0e, 0x44, 0xbe, 0x00,
	0x38, 0x4b, 0xdf, 0xa2, 0x1a, 0x25, 0x7c, 0x86, 0x76, 0xee, 0xda, 0xa7, 0x4f, 0x97, 0x8d, 0xa5,
	0x11, 0x8f, 0xc5, 0x6f, 0x38, 0x1d, 0x67, 0x4a, 0xc8, 0x19, 0xab, 0x20, 0x0b, 0xde, 0x6b, 0x31,
	0xc5, 0x94, 0xd6, 0xfe, 0x03, 0xcf, 0x22, 0xc9, 0x4b, 0xf0, 0x27, 0x8a, 0x8b, 0x18, 0x95, 0xa6,
	0x75, 0xbb, 0xb9, 0xcf, 0xb7, 0xb3, 0xbe, 0x55, 0x8a, 0x3f, 0xb0, 0x02, 0x4c, 0xbe, 0x86, 0xf6,
	0x38, 0x52, 0x88, 0x52, 0xdf, 0xa7, 0x99, 0xa6, 0x07, 0xfb, 0xb9, 0x55, 0x3c, 0xf9, 0x18, 0x5a,
	0xe3, 0x39, 0x46, 0x82, 0xc7, 0x76, 0xda, 0x77, 0x17, 0xbb, 0x84, 0x19, 0xc6, 0x85, 0x12, 0x28,
	0xa7, 0xda, 0x9d, 0x82, 0x9d, 0x0c, 0x07, 0x23, 0xdf, 0x40, 0xe7, 0x8c, 0xcf, 0xf5, 0x22, 0xc6,
	0x21, 0x4a, 0x54, 0x22, 0xb2, 0x67, 0x61, 0x37, 0x71, 0x0d, 0x4d, 0xbe, 0x84, 0x43, 0xe7, 0x19,
	0x27, 0x3c, 0x8e, 0xdd, 0x49, 0xd9, 0xc5, 0x5e, 0xc1, 0x86, 0xff, 0x78, 0xd0, 0xdd, 0xf6, 0x15,
	0xc8, 0x07, 0x50, 0x3b, 0xbf, 0xa6, 0x9e, 0xfd, 0x5c, 0x47, 0x6b, 0x13, 0xcb, 0x6a, 0xe7, 0xd7,
	0x06, 0xc0, 0x5e, 0x39, 0x7d, 0xd9, 0x04, 0xb0, 0x57, 0x06, 0x70, 0xc1, 0xdc, 0x66, 0x6d, 0x02,
	0x2e, 0x98, 0x4d, 0x31, 0x76, 0x3b, 0xb2, 0x25, 0xc5, 0xd8, 0x00, 0x06, 0xe7, 0xb4, 0xb1, 0x03,
	0x30, 0x38, 0x37, 0x80, 0xd1, 0x84, 0x36, 0x77, 0x00, 0x46, 0x13, 0x03, 0xb8, 0x9d, 0xd0, 0xd6,
	0x0e, 0xc0, 0xed, 0x24, 0x7c, 0xe7, 0x55, 0x0e, 0x68, 0xd1, 0xb4, 0xb7, 0xaf, 0x69, 0x6f, 0x5f,
	0xd3, 0xde, 0xbe, 0xa6, 0xbd, 0x7d, 0x4d, 0x7b, 0xfb, 0x9a, 0xf6, 0xf6, 0x35, 0xed, 0xed, 0x6a,
	0xfa, 0x07, 0xf0, 0x97, 0xb6, 0x11, 0xf3, 0xd7, 0xa9, 0x88, 0xf2, 0x33, 0xec, 0xb3, 0xdc, 0x20,
	0xef, 0x41, 0x30, 0x92, 0x19, 0xaa, 0x37, 0x3c, 0x42, 0xdb, 0xae, 0xcf, 0x4a, 0x87, 0x89, 0x8e,
	0x17, 0x77, 0x99, 0x91, 0x7d, 0x6d, 0x7b, 0xf5, 0x59, 0xe9, 0x08, 0xff, 0xf2, 0x56, 0xc5, 0x92,
	0x7c, 0x0e, 0xad, 0xef, 0x85, 0x9c, 0xa6, 0xbf, 0x68, 0xf7, 0x69, 0x9f, 0xaf, 0x49, 0xed, 0x8a,
	0x7e, 0x2e, 0xb1, 0xe4, 0x13, 0x68, 0x5c, 0xf1, 0xe8, 0x46, 0xbb, 0xcf, 0xfd, 0x28, 0x29, 0x47,
	0x1a, 0xca, 0xa5, 0x90, 0x8b, 0x5f, 0xdd, 0x06, 0x3c, 0x4e, 0xb1, 0xc8, 0xf0, 0x0f, 0x0f, 0xba,
	0xdb, 0xe2, 0xe4, 0x33, 0x68, 0x5d, 0x09, 0x29, 0x12, 0x1e, 0xbb, 0xaa, 0x9f, 0x2d, 0x15, 0x91,
	0x47, 0xf7, 0x42, 0xe2, 0x6a, 0xd1, 0x0e, 0x4a, 0xbe, 0x32, 0x97, 0x5a, 0x94, 0x26, 0x09, 0xca,
	0x29, 0x4e, 0x5d, 0xe9, 0x8f, 0x31, 0xab, 0xf0, 0xf0, 0xcf, 0x1a, 0x3c, 0xd9, 0x02, 0x22, 0x4f,
	0xa1, 0x39, 0x7e, 0xd0, 0x19, 0x26, 0xee, 0x86, 0x77, 0x96, 0xd9, 0x88, 0x5b, 0x95, 0x46, 0xa8,
	0x75, 0xaa, 0x6c, 0xae, 0x80, 0x95, 0x0e, 0xa3, 0xf6, 0x43, 0xc5, 0xe7, 0xf7, 0x22, 0xd2, 0xee,
	0x6a, 0x2f, 0x6c, 0xb3, 0xed, 0xe3, 0x74, 0x21, 0xa7, 0xee, 0x62, 0xcf, 0x0d, 0x72, 0x0c, 0x75,
	0xc6, 0x13, 0x3b, 0x7c, 0x0d, 0x66, 0xfe, 0x92, 0x10, 0x0e, 0x19, 0x4f, 0x06, 0x22, 0x41, 0xa9,
	0xcd, 0x13, 0xa1, 0x69, 0xe1, 0x2b, 0x3e, 0x42, 0xa1, 0x35, 0xce, 0x52, 0x65, 0x2e, 0x82, 0x96,
	0x65, 0x2e, 0x4d, 0xf2, 0x02, 0x8e, 0xdd, 0xdf, 0x72, 0x05, 0xdf, 0xae, 0xb0, 0xe1, 0x37, 0x15,
	0xdd, 0x64, 0xf7, 0xa8, 0x68, 0x90, 0x57, 0x64, 0x8d, 0xf0, 0xbb, 0xca, 0x45, 0x6d, 0x12, 0x55,
	0x07, 0xc9, 0x2f, 0x67, 0xa5, 0x5b, 0x9d, 0x15, 0x7f, 0x39, 0x0e, 0xdd, 0xea, 0x38, 0xf8, 0xcb,
	0x1d, 0x3f, 0x05, 0x28, 0x1f, 0x24, 0x1b, 0x0f, 0xa7, 0xe2, 0x71, 0x53, 0xab, 0x3c, 0x6e, 0xc2,
	0x21, 0x04, 0xc5, 0x23, 0x62, 0x83, 0xf2, 0xc2, 0xbd, 0x91, 0x1e, 0xbf, 0xcd, 0x2c, 0x26, 0xfc,
	0xdd, 0x83, 0xa3, 0xb5, 0x88, 0x59, 0xcf, 0xa9, 0x4e, 0x60, 0x45, 0xa6, 0x53, 0x88, 0x4c, 0x60,
	0x35, 0xa5, 0x53, 0x68, 0x4a, 0x60, 0x25, 0xa4, 0x53, 0x48, 0x48, 0x60, 0x15, 0xa3, 0x53, 0x28,
	0x46, 0x60, 0x05, 0xa2, 0x53, 0x08, 0x44, 0x60, 0xf5, 0xa0, 0x53, 0xe8, 0x41, 0x60, 0x8f, 0xff,
	0xfb, 0xd0, 0x36, 0x2f, 0xc7, 0x01, 0xc6, 0x98, 0xe1, 0x74, 0xbd, 0x9d, 0xbb, 0xa6, 0xad, 0xff,
	0xd3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x17, 0x78, 0x07, 0x75, 0x0a, 0x00, 0x00,
}
