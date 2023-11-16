package stemma

import "encoding/xml"

type Group struct {
	Key            string           `xml:"Key,attr"`
	Title          *string          `xml:"Title,omitempty"`
	Type           *string          `xml:"Type,omitempty"`
	SubType        *string          `xml:"SubType,omitempty"`
	NameVariants   []Names          `xml:"Names,omitempty"`
	ParentGroupLnk []ParentGroupLnk `xml:"ParentGroupLnk,omitempty"`
	Creation       *LifeEvent       `xml:"Creation,omitempty"`
	Demise         *LifeEvent       `xml:"Demise,omitempty"`
	Eventlets      []Eventlet       `xml:"Eventlet,omitempty"`
	SourceLnk      []SourceLnk      `xml:"SourceLnk,omitempty"`
	RelatedTo      []RelatedTo      `xml:"RelatedTo,omitempty"`
	Operations     []Operation      `xml:"Operation,omitempty"`
	ExternalID     *string          `xml:"ExternalID,omitempty"`
	TextSeg        []TextSegment    `xml:"TextSegment,omitempty"`
}

type RangeFrom struct {
	AfterEvent *string `xml:"AfterEvent,attr,omitempty"`
	FromEvent  *string `xml:"FromEvent,attr,omitempty"`
	After      *string `xml:"After,attr,omitempty"`
	From       *string `xml:"From,attr,omitempty"`
}

type RangeTo struct {
	BeforeEvent *string `xml:"BeforeEvent,attr,omitempty"`
	UntilEvent  *string `xml:"UntilEvent,attr,omitempty"`
	Before      *string `xml:"Before,attr,omitempty"`
	Until       *string `xml:"Until,attr,omitempty"`
}

type Creation struct {
	EventLnk []EventLnk    `xml:"EventLnk,omitempty"`
	Eventlet []Eventlet    `xml:"Eventlet,omitempty"`
	JoinFrom []JoinFrom    `xml:"JoinFrom,omitempty"`
	TextSeg  []TextSegment `xml:"TextSegment,omitempty"`
}

type Demise struct {
	EventLnk []EventLnk    `xml:"EventLnk,omitempty"`
	Eventlet []Eventlet    `xml:"Eventlet,omitempty"`
	SplitTo  []SplitTo     `xml:"SplitTo,omitempty"`
	TextSeg  []TextSegment `xml:"TextSegment,omitempty"`
}

type JoinFrom struct {
	Key string `xml:"Key,attr"`
}
type SplitTo struct {
	Key string `xml:"Key,attr"`
}

type Animal struct {
	Key             string           `xml:"Key,attr"`
	Title           *string          `xml:"Title,omitempty"`
	Type            *string          `xml:"Type,omitempty"`
	SubType         *string          `xml:"SubType,omitempty"`
	Sex             *Sex             `xml:"Sex,omitempty"`
	NameVariants    []Names          `xml:"Names,omitempty"`
	FatherAnimalLnk *ParentAnimalLnk `xml:"FatherAnimalLnk,omitempty"`
	MotherAnimalLnk *ParentAnimalLnk `xml:"MotherAnimalLnk,omitempty"`
	Birth           *LifeEvent       `xml:"Birth,omitempty"`
	Death           *LifeEvent       `xml:"Death,omitempty"`
	MemberOf        []MemberOf       `xml:"MemberOf,omitempty"`
	Eventlets       []Eventlet       `xml:"Eventlet,omitempty"`
	SourceLnk       []SourceLnk      `xml:"SourceLnk,omitempty"`
	ExternalID      *string          `xml:"ExternalID,omitempty"`
	TextSeg         []TextSegment    `xml:"TextSegment,omitempty"`
}

type AnimalName struct {
	DataAttribute string `xml:",attr,omitempty"`
	AnimalName    string `xml:",chardata"`
}

type ParentAnimalLnk struct {
	Key           string        `xml:"Key,attr"`
	DataAttribute string        `xml:",attr,omitempty"`
	TextSeg       []TextSegment `xml:"TextSegment,omitempty"`
}

type MemberOf struct {
	Key       string        `xml:"Key,attr"`
	RangeFrom *RangeFrom    `xml:"RangeFrom,attr,omitempty"`
	RangeTo   *RangeTo      `xml:"RangeTo,attr,omitempty"`
	TextSeg   []TextSegment `xml:"TextSegment,omitempty"`
}

type Citation struct {
	Key               string              `xml:"Key,attr"`
	Abstract          *bool               `xml:"Abstract,attr,omitempty"`
	Title             *string             `xml:"Title,omitempty"`
	URI               *string             `xml:"URI,omitempty"`
	Params            []Param             `xml:"Params>Param,omitempty"`
	DisplayFormat     []DisplayFormat     `xml:"DisplayFormat,omitempty"`
	ParentCitationLnk []ParentCitationLnk `xml:"ParentCitationLnk,omitempty"`
	BaseCitationLnk   *BaseCitationLnk    `xml:"BaseCitationLnk,omitempty"`
	TextSeg           []TextSegment       `xml:"TextSegment,omitempty"`
}

type DisplayFormat struct {
	Mode    string        `xml:"Mode,attr,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type ParentCitationLnk struct {
	Key               string              `xml:"Key,attr"`
	Type              *string             `xml:"Type,attr,omitempty"`
	ParamValues       []ParamValue        `xml:"Param,omitempty"`
	ParentCitationLnk []ParentCitationLnk `xml:"ParentCitationLnk,omitempty"`
	TextSeg           []TextSegment       `xml:"TextSegment,omitempty"`
}

type BaseCitationLnk struct {
	Key     string        `xml:"Key,attr"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type Param struct {
	Name     string  `xml:"Name,attr"`
	Type     *string `xml:"Type,attr,omitempty"`
	SemType  *string `xml:"SemType,attr,omitempty"`
	ItemList *bool   `xml:"ItemList,attr,omitempty"`
	Optional *bool   `xml:"Optional,attr,omitempty"`
	WhereIn  *bool   `xml:"WhereIn,attr,omitempty"`
	Value    string  `xml:",chardata"`
}

type ParamValue struct {
	Param string  `xml:"Name,attr"`
	Key   *string `xml:"Key,attr,omitempty"`
	Subst *string `xml:"Subst,attr,omitempty"`
	Value string  `xml:",chardata"`
	Items []Item  `xml:"Item,omitempty"`
}

type Item struct {
	Key   *string `xml:"Key,attr,omitempty"`
	Value string  `xml:",chardata"`
}

type Constraints struct {
	Constraints []Constraint `xml:"Constraint"`
}

type Datasets struct {
	XMLName xml.Name  `xml:"Datasets"`
	Header  Header    `xml:"Header"`
	Dataset []Dataset `xml:"Dataset"`
}

type Header struct {
	Created   string        `xml:"Created"`
	Product   Product       `xml:"Product"`
	Languages []string      `xml:"Language,omitempty"`
	Locales   []string      `xml:"Locale,omitempty"`
	TextSeg   []TextSegment `xml:"TextSegment,omitempty"`
}

type Product struct {
	Id      string `xml:"Id"`
	Name    string `xml:"Name"`
	Version string `xml:"Version"`
}

type Dataset struct {
	Name string `xml:"Name,attr"`
	//ExtendedProperties []ExtendedProperties `xml:"EXTENDED_PROPERTIES,omitempty"`
	//Imports            []Import             `xml:"IMPORTS,omitempty"`
	DatasetBody []DatasetBody `xml:",any"`
}

type Content struct {
	Created      string        `xml:"Created"`
	Author       string        `xml:"Author"`
	Version      string        `xml:"Version"`
	Copyright    *string       `xml:"Copyright,omitempty"`
	LastModified *LastModified `xml:"LastModified,omitempty"`
	Languages    []string      `xml:"Language,omitempty"`
	Locales      []string      `xml:"Locale,omitempty"`
	Counters     []Counter     `xml:"Counters>Counter,omitempty"`
	TextSeg      []TextSegment `xml:"TextSegment,omitempty"`
}

type LastModified struct {
	DateTime   string `xml:"iso-datetime"`
	ModifiedBy string `xml:"ModifiedBy"`
}

type Counter struct {
	Tag   *string `xml:"Tag,attr,omitempty"`
	Value int     `xml:"integer"`
}

type DatasetBody struct {
	Person        []Person    `xml:"Person,omitempty"`
	Animal        []Animal    `xml:"Animal,omitempty"`
	Place         []Place     `xml:"Place,omitempty"`
	Group         []Group     `xml:"Group,omitempty"`
	Event         []Event     `xml:"Event,omitempty"`
	Citation      []Citation  `xml:"Citation,omitempty"`
	Resource      []Resource  `xml:"Resource,omitempty"`
	Source        []Source    `xml:"Source,omitempty"`
	Matrix        []Matrix    `xml:"Matrix,omitempty"`
	NarrativeText []Narrative `xml:"Narrative,omitempty"`
}

type Event struct {
	Key            string           `xml:"Key,attr"`
	Abstract       *bool            `xml:"Abstract,attr,omitempty"`
	Title          *string          `xml:"Title,omitempty"`
	Type           *string          `xml:"Type,omitempty"`
	SubType        *string          `xml:"SubType,omitempty"`
	PlaceLnk       *PlaceLnk        `xml:"PlaceLnk,omitempty"`
	When           When             `xml:"When"`
	Until          *When            `xml:"Until,omitempty"`
	ParentEventLnk []ParentEventLnk `xml:"ParentEventLnk,omitempty"`
	BaseEventLnk   *BaseEventLnk    `xml:"BaseEventLnk,omitempty"`
	SourceLnks     []SourceLnk      `xml:"SourceLnk,omitempty"`
	ExternalID     *string          `xml:"ExternalID,omitempty"`
	TextSeg        []TextSegment    `xml:"TextSegment,omitempty"`
}

type ParentEventLnk struct {
	Key     string        `xml:"Key,attr"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type BaseEventLnk struct {
	Key     string        `xml:"Key,attr"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type EvSourceLnk struct {
	Key             string          `xml:"Key,attr"`
	PersonLnks      []PersonLnk     `xml:"PersonLnk,omitempty"`
	AnimalLnks      []AnimalLnk     `xml:"AnimalLnk,omitempty"`
	PlaceLnks       []PlaceLnk      `xml:"PlaceLnk,omitempty"`
	GroupLnks       []GroupLnk      `xml:"GroupLnk,omitempty"`
	EventProperties []EventProperty `xml:"EventProperty,omitempty"`
	TextSeg         []TextSegment   `xml:"TextSegment,omitempty"`
}

type EventProperty struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type Eventlet struct {
	Title      *string       `xml:"Title,omitempty"`
	Type       *string       `xml:"Type,omitempty"`
	SubType    *string       `xml:"SubType,omitempty"`
	PlaceLnk   *PlaceLnk     `xml:"PlaceLnk,omitempty"`
	When       When          `xml:"When"`
	Until      *When         `xml:"Until,omitempty"`
	SourceLnks []EvSourceLnk `xml:"EvSourceLnk,omitempty"`
	TextSeg    []TextSegment `xml:"TextSegment,omitempty"`
}

type EventConstraints struct {
	Constraints []Constraint `xml:"Constraint"`
}

type Constraint struct {
	AfterEvent  *string       `xml:"AfterEvent,attr,omitempty"`
	BeforeEvent *string       `xml:"BeforeEvent,attr,omitempty"`
	FromEvent   *string       `xml:"FromEvent,attr,omitempty"`
	UntilEvent  *string       `xml:"UntilEvent,attr,omitempty"`
	AtEvent     *string       `xml:"AtEvent,attr,omitempty"`
	TextSeg     []TextSegment `xml:"TextSegment,omitempty"`
}

type DateEntity struct {
	Values  []DateValue   `xml:"Value"`
	Ranges  []DateRange   `xml:"Range"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type DateValue struct {
	Calendar *string `xml:"Calendar,attr,omitempty"`
	Calc     *bool   `xml:"Calc,attr,omitempty"`
	Margin   *string `xml:"Margin,attr,omitempty"`
	Value    string  `xml:",chardata"`
}

type DateRange struct {
	Calendar *string `xml:"Calendar,attr,omitempty"`
	Calc     *bool   `xml:"Calc,attr,omitempty"`
	Min      string  `xml:"Min"`
	Max      string  `xml:"Max"`
}

type When struct {
	Value            string        `xml:"Value,attr,omitempty"`
	DataAttribute    string        `xml:",attr,omitempty"`
	DateEntity       []DateEntity  `xml:"DateEntity,omitempty"`
	EventConstraints []Constraint  `xml:"Constraint,omitempty"`
	TextSeg          []TextSegment `xml:"TextSegment,omitempty"`
}

type PropertyDef struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type GroupName struct {
	DataAttribute string `xml:",attr"`
	GroupName     string `xml:",chardata"`
}

type ParentGroupLnk struct {
	RangeFrom RangeFrom     `xml:"RangeFrom,attr,omitempty"`
	RangeTo   RangeTo       `xml:"RangeTo,attr,omitempty"`
	Key       string        `xml:"Key,attr"`
	TextSeg   []TextSegment `xml:"TextSegment,omitempty"`
}

type GroupLifeEvent struct {
	EventLnk  []EventLnk    `xml:"EventLnk,omitempty"`
	Eventlets []Eventlet    `xml:"Eventnet,omitempty"`
	JoinFrom  []JoinFrom    `xml:"JoinFrom,omitempty"`
	SplitTo   []SplitTo     `xml:"SplitTo,omitempty"`
	TextSeg   []TextSegment `xml:"TextSegment,omitempty"`
}

type RelatedTo struct {
	GroupLnk []GroupLnk    `xml:"GroupLnk,omitempty"`
	TextSeg  []TextSegment `xml:"TextSegment,omitempty"`
}

type Operation struct {
	Name string `xml:"Name,attr"`
	Key  string `xml:"Key,attr"`
}

type GroupProperty struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type PersonLnk struct {
	Key              *string          `xml:"Key,attr,omitempty"`
	PersonProperties []PersonProperty `xml:"PersonProperty,omitempty"`
	TextSeg          []TextSegment    `xml:"TextSegment,omitempty"`
}

type AnimalLnk struct {
	Key              *string          `xml:"Key,attr,omitempty"`
	AnimalProperties []AnimalProperty `xml:"AnimalProperty,omitempty"`
	TextSeg          []TextSegment    `xml:"TextSegment,omitempty"`
}

type PlaceLnk struct {
	Key             *string         `xml:"Key,attr,omitempty"`
	PlaceProperties []PlaceProperty `xml:"PlaceProperty,omitempty"`
	TextSeg         []TextSegment   `xml:"TextSegment,omitempty"`
}

type GroupLnk struct {
	Key             *string         `xml:"Key,attr,omitempty"`
	GroupProperties []GroupProperty `xml:"GroupProperty,omitempty"`
	TextSeg         []TextSegment   `xml:"TextSegment,omitempty"`
}

type EventLnk struct {
	Key     string        `xml:"Key,attr"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type ResourceLnk struct {
	Key         string        `xml:"Key,attr"`
	ParamValues []ParamValue  `xml:"ParamValue,omitempty"`
	TextSeg     []TextSegment `xml:"TextSegment,omitempty"`
}

type CitationLnk struct {
	Key                string              `xml:"Key,attr"`
	ParamValues        []ParamValue        `xml:"ParamValue,omitempty"`
	ParentCitationLnks []ParentCitationLnk `xml:"ParentCitationLnk,omitempty"`
	TextSeg            []TextSegment       `xml:"TextSegment,omitempty"`
}

type SourceLnk struct {
	Key              string           `xml:"Key,attr"`
	AnimalProperties []AnimalProperty `xml:"AnimalProperty,omitempty"`
	PlaceProperties  []PlaceProperty  `xml:"PlaceProperty,omitempty"`
	GroupProperties  []GroupProperty  `xml:"GroupProperty,omitempty"`
	TextSeg          []TextSegment    `xml:"TextSegment,omitempty"`
}

type PersonProperty struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type AnimalProperty struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type PlaceProperty struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type Matrix struct {
	Key         string         `xml:"Key,attr"`
	Title       *string        `xml:"Title,omitempty"`
	Frame       Frame          `xml:"Frame,omitempty"`
	ProtoPerson *[]ProtoPerson `xml:"ProtoPerson,omitempty"`
	ProtoAnimal *[]ProtoAnimal `xml:"ProtoAnimal,omitempty"`
	ProtoPlace  *[]ProtoPlace  `xml:"ProtoPlace,omitempty"`
	ProtoGroup  *[]ProtoGroup  `xml:"ProtoGroup,omitempty"`
	ProtoEvent  *[]ProtoEvent  `xml:"ProtoEvent,omitempty"`
	Commentary  []Commentary   `xml:"Commentary,omitempty"`
	Dates       []ProtoDate    `xml:"ProtoDate,omitempty"`
	TextSeg     []TextSegment  `xml:"TextSegment,omitempty"`
}

type Frame struct {
	SourceLnk []SourceLnk   `xml:"SourceLnk,omitempty"`
	TextSeg   []TextSegment `xml:"TextSegment,omitempty"`
}

type ProtoAnimal struct {
	DetKey  string        `xml:"DetKey,attr,omitempty"`
	Key     *string       `xml:"Key,attr,omitempty"`
	Title   *string       `xml:"Title,omitempty"`
	Links   []Link        `xml:"Link,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type ProtoPerson struct {
	DetKey  string        `xml:"DetKey,attr,omitempty"`
	Key     *string       `xml:"Key,attr,omitempty"`
	Title   *string       `xml:"Title,omitempty"`
	Links   []Link        `xml:"Link,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type ProtoPlace struct {
	DetKey  string        `xml:"DetKey,attr,omitempty"`
	Key     *string       `xml:"Key,attr,omitempty"`
	Title   *string       `xml:"Title,omitempty"`
	Links   []Link        `xml:"Link,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type ProtoGroup struct {
	DetKey  string        `xml:"DetKey,attr,omitempty"`
	Key     *string       `xml:"Key,attr,omitempty"`
	Title   *string       `xml:"Title,omitempty"`
	Links   []Link        `xml:"Link,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type ProtoEvent struct {
	DetKey  string        `xml:"DetKey,attr,omitempty"`
	Key     *string       `xml:"Key,attr,omitempty"`
	Title   *string       `xml:"Title,omitempty"`
	Links   []Link        `xml:"Link,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type Commentary struct {
	DetKey  string        `xml:"DetKey,attr"`
	Title   *string       `xml:"Title,omitempty"`
	Links   []Link        `xml:"Link,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type ProtoDate struct {
	DetKey  string        `xml:"DetKey,attr"`
	Title   *string       `xml:"Title,omitempty"`
	Links   []Link        `xml:"Link,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type Link struct {
	Type    *string       `xml:"Type,attr,omitempty"`
	DetLnk  *string       `xml:"DetLnk,attr,omitempty"`
	Value   *string       `xml:"Value,attr,omitempty"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type Narrative struct {
	Key     string   `xml:"Key,attr"`
	Title   *string  `xml:"Title,omitempty"`
	TextSeg []string `xml:",chardata"`
}

type Contact struct {
	Key            string          `xml:"Key,attr"`
	Names          []Name          `xml:"Name,omitempty"`
	ContactDetails *ContactDetails `xml:"ContactDetails,omitempty"`
	TextSeg        []TextSegment   `xml:"TextSegment,omitempty"`
}

type Name struct {
	Mode string `xml:"Mode,attr,omitempty"`
	Text string `xml:",chardata"`
}

type ContactDetails struct {
	Address   *Address      `xml:"Address,omitempty"`
	Phone     []Number      `xml:"Phone>Number,omitempty"`
	Emails    []Email       `xml:"Emails>Email,omitempty"`
	Web       []URL         `xml:"Web>URL,omitempty"`
	Messaging []Message     `xml:"Messaging>Message,omitempty"`
	TextSeg   []TextSegment `xml:"TextSegment,omitempty"`
}

type TextSegment struct {
	Key           string  `xml:"Key,attr,omitempty"`
	TextType      *string `xml:"TextType,attr,omitempty"`
	DataAttribute *string `xml:",attr,omitempty"` // Assuming DATA_ATTRIBUTE is a placeholder for actual attribute name.
	Title         *string `xml:"Title,omitempty"`
	NarrativeText string  `xml:",chardata"` // The narrative text that appears as inner XML text content.
}

type FromText struct {
	Key string `xml:"Key,attr"`
}

type Address struct {
	AddressLine1    *string `xml:"AddressLine1,omitempty"`
	AddressLine2    *string `xml:"AddressLine2,omitempty"`
	AddressLine3    *string `xml:"AddressLine3,omitempty"`
	TownOrCity      *string `xml:"TownOrCity,omitempty"`
	StateOrProvince *string `xml:"StateOrProvince,omitempty"`
	PostalCode      *string `xml:"PostalCode,omitempty"`
	CountryCode     *string `xml:"CountryCode,omitempty"`
}

type Number struct {
	Tag  string `xml:"Tag,attr,omitempty"`
	Text string `xml:",chardata"`
}

type Email struct {
	Tag  string `xml:"Tag,attr,omitempty"`
	Text string `xml:",chardata"`
}

type URL struct {
	Tag  string `xml:"Tag,attr,omitempty"`
	Text string `xml:",chardata"`
}

type Message struct {
	Tag  string `xml:"Tag,attr,omitempty"`
	Text string `xml:",chardata"`
}

type Names struct {
	Sequences []SequenceVariant `xml:"Sequences"`
}

type SequenceVariant struct {
	RangeFrom     RangeFrom      `xml:"RangeFrom,attr,omitempty"`
	RangeTo       RangeTo        `xml:"RangeTo,attr,omitempty"`
	Type          string         `xml:"Type,attr,omitempty"`
	Culture       string         `xml:"Culture,attr,omitempty"`
	DataAttribute string         `xml:",attr,omitempty"`
	Canonical     []Canonical    `xml:"Canonical"`
	Sequences     []NameSequence `xml:"Sequence"`
}

type Canonical struct {
	Mode   string `xml:"Mode,attr,omitempty"`
	SortAs string `xml:"SortAs,attr,omitempty"`
	Text   string `xml:",chardata"`
}

type NameSequence struct {
	NameAttribute string        `xml:",attr,omitempty"`
	Tokens        []TokenBlock  `xml:"Tokens"`
	TextSeg       []TextSegment `xml:"TextSegment,omitempty"`
}

type TokenBlock struct {
	Optional bool    `xml:"Optional,attr,omitempty"`
	Initial  bool    `xml:"Initial,attr,omitempty"`
	Tokens   []Token `xml:"Token"`
}

type Token struct {
	Text string `xml:",chardata"`
}

type Person struct {
	Key             string           `xml:"Key,attr"`
	Title           *string          `xml:"Title,omitempty"`
	Sex             *Sex             `xml:"Sex,omitempty"`
	NameVariants    []Names          `xml:"Names,omitempty"`
	PersonalNames   []PersonalName   `xml:"PersonalName,omitempty"`
	FatherPersonLnk *ParentPersonLnk `xml:"FatherPersonLnk,omitempty"`
	MotherPersonLnk *ParentPersonLnk `xml:"MotherPersonLnk,omitempty"`
	Birth           *LifeEvent       `xml:"Birth,omitempty"`
	Death           *LifeEvent       `xml:"Death,omitempty"`
	MemberOf        []MemberOf       `xml:"MemberOf,omitempty"`
	Eventlets       []Eventlet       `xml:"Eventlet,omitempty"`
	SourceLnk       []SourceLnk      `xml:"SourceLnk,omitempty"`
	ContactDetails  *ContactDetails  `xml:"ContactDetails,omitempty"`
	ExternalID      *string          `xml:"ExternalID,omitempty"`
	TextSeg         []TextSegment    `xml:"TextSegment,omitempty"`
}

type Sex struct {
	DataAttribute string `xml:",attr,omitempty"`
	Value         string `xml:",chardata"`
}

type PersonalName struct {
	DataAttribute string `xml:",attr,omitempty"`
	Name          string `xml:",chardata"`
}

type ParentPersonLnk struct {
	Key           string        `xml:"Key,attr"`
	DataAttribute string        `xml:",attr,omitempty"`
	TextSeg       []TextSegment `xml:"TextSegment,omitempty"`
}

type LifeEvent struct {
	EventLnk []EventLnk    `xml:"EventLnk,omitempty"`
	Eventlet []Eventlet    `xml:"Eventlet,omitempty"`
	TextSeg  []TextSegment `xml:"TextSegment,omitempty"`
}

type Place struct {
	Key            string           `xml:"Key,attr"`
	Title          *string          `xml:"Title,omitempty"`
	Type           *string          `xml:"Type,omitempty"`
	Category       *string          `xml:"Category,omitempty"`
	NameVariants   []Names          `xml:"Names,omitempty"`
	PlaceNames     []PlaceName      `xml:"PlaceName,omitempty"`
	Locations      []Location       `xml:"Location,omitempty"`
	ParentPlaceLnk []ParentPlaceLnk `xml:"ParentPlaceLnk,omitempty"`
	Creation       *LifeEvent       `xml:"Creation,omitempty"`
	Demise         *LifeEvent       `xml:"Demise,omitempty"`
	Eventlets      []Eventlet       `xml:"Eventlet,omitempty"`
	SourceLnk      []SourceLnk      `xml:"SourceLnk,omitempty"`
	RelatedTo      []RelatedTo      `xml:"RelatedTo,omitempty"`
	ExternalID     *string          `xml:"ExternalID,omitempty"`
	TextSeg        []TextSegment    `xml:"TextSegment,omitempty"`
}

type PlaceName struct {
	DataAttribute string `xml:",attr,omitempty"`
	PlaceName     string `xml:",chardata"`
}

type Location struct {
	Open        *bool         `xml:"Open,attr,omitempty"`
	RangeFrom   *RangeFrom    `xml:"RangeFrom,attr,omitempty"`
	RangeTo     *RangeTo      `xml:"RangeTo,attr,omitempty"`
	Coordinates string        `xml:",chardata"`
	TextSeg     []TextSegment `xml:"TextSegment,omitempty"`
}

type ParentPlaceLnk struct {
	RangeFrom RangeFrom     `xml:"RangeFrom,attr,omitempty"`
	RangeTo   RangeTo       `xml:"RangeTo,attr,omitempty"`
	Key       string        `xml:"Key,attr"`
	TextSeg   []TextSegment `xml:"TextSegment,omitempty"`
}

type Resource struct {
	Key             string           `xml:"Key,attr"`
	Abstract        *bool            `xml:"Abstract,attr,omitempty"`
	Title           *string          `xml:"Title,omitempty"`
	URL             *ResourceURL     `xml:"URL,omitempty"`
	Type            *ResourceType    `xml:"Type,omitempty"`
	DataControl     *DataControl     `xml:"DataControl,omitempty"`
	Params          *Params          `xml:"Params,omitempty"`
	BaseResourceLnk *BaseResourceLnk `xml:"BaseResourceLnk,omitempty"`
	TextSeg         []TextSegment    `xml:"TextSegment,omitempty"`
}

type ResourceURL struct {
	ContentType string `xml:"ContentType,attr,omitempty"`
	URL         string `xml:",chardata"`
}

type ResourceType struct {
	Artefact *bool  `xml:"Artefact,attr,omitempty"`
	Type     string `xml:",chardata"`
}

type DataControl struct {
	Sensitivity *string       `xml:"Sensitivity,omitempty"`
	Copyright   *string       `xml:"Copyright,omitempty"`
	Permission  *string       `xml:"Permission,omitempty"`
	Prohibition *string       `xml:"Prohibition,omitempty"`
	Attribution *string       `xml:"Attribution,omitempty"`
	TextSeg     []TextSegment `xml:"TextSegment,omitempty"`
}

type Params struct {
	ParamDef   []ParamDef   `xml:"PARAM_DEF,omitempty"`
	ParamValue []ParamValue `xml:"PARAM_VALUE,omitempty"`
}

type ParamDef struct {
	Name         string  `xml:"Name,attr"`
	Type         *string `xml:"Type,attr,omitempty"`
	SemType      *string `xml:"SemType,attr,omitempty"`
	ItemList     *bool   `xml:"ItemList,attr,omitempty"`
	Optional     *bool   `xml:"Optional,attr,omitempty"`
	WhereIn      *bool   `xml:"WhereIn,attr,omitempty"`
	DefaultValue string  `xml:",chardata"`
}

type BaseResourceLnk struct {
	Key     string        `xml:"Key,attr"`
	TextSeg []TextSegment `xml:"TextSegment,omitempty"`
}

type Source struct {
	Key        string         `xml:"Key,attr"`
	Title      *string        `xml:"Title,omitempty"`
	Frame      *Frame         `xml:"Frame,omitempty"`
	ProtoSubj  []ProtoSubject `xml:"ProtoSub,omitempty"`
	Commentary []Commentary   `xml:"Commentary,omitempty"`
	Dates      []ProtoDate    `xml:"ProtoDate,omitempty"`
	SourceLets []SourceLet    `xml:"Sourcelet,omitempty"`
	TextSeg    []TextSegment  `xml:"TextSegment,omitempty"`
}

type Where struct {
	DetLnk string `xml:"DetLnk,attr"`
}

type SourceLet struct {
	Key        *string        `xml:"Key,attr,omitempty"`
	Title      *string        `xml:"Title,omitempty"`
	Frame      *Frame         `xml:"Frame,omitempty"`
	ProtoSubj  []ProtoSubject `xml:"ProtoSubject,omitempty"`
	Commentary []Commentary   `xml:"Commentary,omitempty"`
	Dates      []ProtoDate    `xml:"ProtoDate,omitempty"`
	TextSeg    []TextSegment  `xml:"TextSegment,omitempty"`
}

// ProtoSubject can be an interface that ProtoPerson, ProtoAnimal, ProtoPlace, ProtoGroup, ProtoEvent implement
type ProtoSubject interface {
	// Common methods that the prototypes would share (if any)
	IsProtoSubject() bool
}

// Note: Similar struct definitions would exist for ProtoAnimal, ProtoPlace, ProtoGroup, ProtoEvent

func (p ProtoPerson) IsProtoSubject() bool {
	return true
}

func (p ProtoAnimal) IsProtoSubject() bool {
	return true
}

func (p ProtoPlace) IsProtoSubject() bool {
	return true
}

func (p ProtoEvent) IsProtoSubject() bool {
	return true
}

func (p ProtoGroup) IsProtoSubject() bool {
	return true
}
