package messages

import (
	"encoding/xml"
	"log"

	"github.com/kapytein/syncml/pkg/syncml"
)

type Message struct {
	Name              string
	RequestParameters SyncMLRequestParameters
}

type SyncMLRequestParameters struct {
	MsgID     int
	DeviceID  string
	SessionID int
	SourceURI string
}

func NewMessage(name string, r SyncMLRequestParameters) Message {
	return Message{name, r}
}

func (m Message) GetCommands() syncml.SyncMLCommands {
	commands := syncml.SyncMLCommands{}

	switch m.Name {
	case "basic_info":
		commands = m.getBasicInfoSyncMLRequest()
	}

	return commands
}

func (m Message) BuildSyncMLResponseMessage(commands []syncml.SyncMLCommands) ([]byte, error) {
	var getCommands []syncml.Get
	var deleteCommands []syncml.Delete
	var addCommands []syncml.Add
	var replaceCommands []syncml.Replace
	var execCommands []syncml.Exec

	for _, cmd := range commands {
		getCommands = append(getCommands, cmd.Get...)
		deleteCommands = append(deleteCommands, cmd.Delete...)
		addCommands = append(addCommands, cmd.Add...)
		replaceCommands = append(replaceCommands, cmd.Replace...)
		execCommands = append(execCommands, cmd.Exec...)
	}

	source := syncml.Source{LocURI: m.RequestParameters.SourceURI}
	target := syncml.Target{LocURI: m.RequestParameters.DeviceID}

	status := syncml.Command{CmdID: 1, MsgRef: m.RequestParameters.MsgID, CmdRef: "0", Cmd: "SyncHdr", Data: "200"}
	messageHeader := syncml.SyncHdr{VerDTD: "1.2", VerProto: "DM/1.2", SessionID: m.RequestParameters.SessionID, MsgID: m.RequestParameters.MsgID, Target: target, Source: source}
	messageBody := syncml.SyncBody{Status: []syncml.Command{status}, Get: getCommands, Add: addCommands, Delete: deleteCommands, Replace: replaceCommands, Exec: execCommands}
	message := syncml.SyncML{SyncHdr: messageHeader, SyncBody: messageBody}

	xmlResponse, err := xml.Marshal(message)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return xmlResponse, nil
}
