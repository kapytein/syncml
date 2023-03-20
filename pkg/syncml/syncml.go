package syncml

import "encoding/xml"

type SyncML struct {
	XMLName  xml.Name `xml:"SYNCML:SYNCML1.2 SyncML"`
	SyncHdr  SyncHdr  `xml:"SyncHdr"`
	SyncBody SyncBody `xml:"SyncBody"`
}

type SyncHdr struct {
	VerDTD    string `xml:"VerDTD"`
	VerProto  string `xml:"VerProto"`
	SessionID int    `xml:"SessionID"`
	MsgID     int    `xml:"MsgID"`
	Target    Target `xml:"Target"`
	Source    Source `xml:"Source"`
}
type Target struct {
	LocURI string `xml:"LocURI"`
}

type Source struct {
	LocURI string `xml:"LocURI"`
}

type SyncBody struct {
	Status  []Command `xml:"Status,omitempty"`
	Get     []Get     `xml:"Get,omitempty"`
	Add     []Add     `xml:"Add,omitempty"`
	Delete  []Delete  `xml:"Delete,omitempty"`
	Replace []Replace `xml:"Replace,omitempty"`
	Exec    []Exec    `xml:"Exec,omitempty"`
	Results []Results `xml:"Results,omitempty"`
	Final   xml.Name  `xml:"Final"`
}

type Get struct {
	CmdID string `xml:"CmdID"`
	Item  Item   `xml:"Item"`
}

type Add struct {
	CmdID string `xml:"CmdID"`
	Item  Item   `xml:"Item"`
}

type Delete struct {
	CmdID string `xml:"CmdID"`
	Item  Item   `xml:"Item"`
}

type Exec struct {
	CmdID string `xml:"CmdID"`
	Item  Item   `xml:"Item"`
}

type Replace struct {
	CmdID string `xml:"CmdID"`
	Item  Item   `xml:"Item"`
}

type Results struct {
	CmdID  int  `xml:"CmdID,omitempty"`
	MsgRef int  `xml:"MsgRef,omitempty"`
	CmdRef int  `xml:"CmdRef,omitempty"`
	Item   Item `xml:"Item"`
}

type Item struct {
	Target Target  `xml:"Target,omitempty"`
	Source *Source `xml:"Source,omitempty"`
	Data   string  `xml:"Data,omitempty"`
	Meta   *Meta   `xml:"Meta,omitempty"` // workaround for omitempty - see https://pkg.go.dev/encoding/xml#Marshal
}

type Meta struct {
	Format     string `xml:"syncml:metinf Format,omitempty"`
	NextNonce  string `xml:"syncml:metinf NextNonce,omitempty"`
	MaxMsgSize string `xml:"syncml:metinf MaxMsgSize,omitempty"`
	Type       string `xml:"syncml:metinf Type,omitempty"`
}

type Command struct {
	CmdID  int    `xml:"CmdID"`
	MsgRef int    `xml:"MsgRef"`
	CmdRef string `xml:"CmdRef"`
	Cmd    string `xml:"Cmd"`
	Data   string `xml:"Data"`
}

// Only used for passing around SyncML Commands, not for rendering response
type SyncMLCommands struct {
	Get     []Get
	Add     []Add
	Delete  []Delete
	Replace []Replace
	Exec    []Exec
}
