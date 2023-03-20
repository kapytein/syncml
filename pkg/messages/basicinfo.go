package messages

import (
	"github.com/kapytein/syncml/pkg/syncml"
)

// SyncML request message to retrieve basic device information
func (m Message) getBasicInfoSyncMLRequest() syncml.SyncMLCommands {
	osEdition := syncml.Get{CmdID: "2", Item: syncml.Item{Target: syncml.Target{LocURI: "./DevDetail/Ext/Microsoft/OSPlatform"}}}
	macAddressWLAN := syncml.Get{CmdID: "3", Item: syncml.Item{Target: syncml.Target{LocURI: "./DevDetail/Ext/WLANMACAddress"}}}
	deviceType := syncml.Get{CmdID: "4", Item: syncml.Item{Target: syncml.Target{LocURI: "./DevDetail/DevTyp"}}}
	hostName := syncml.Get{CmdID: "5", Item: syncml.Item{Target: syncml.Target{LocURI: "./DevDetail/Ext/Microsoft/DeviceName"}}}
	arch := syncml.Get{CmdID: "6", Item: syncml.Item{Target: syncml.Target{LocURI: "./DevDetail/Ext/Microsoft/ProcessorArchitecture"}}}
	serialNumber := syncml.Get{CmdID: "6", Item: syncml.Item{Target: syncml.Target{LocURI: "./DevDetail/Ext/Microsoft/SMBIOSSerialNumber"}}}

	return syncml.SyncMLCommands{Get: []syncml.Get{osEdition, macAddressWLAN, deviceType, hostName, arch, serialNumber}}
}
