package bw2bind


//This file is autogenerated from https://github.com/immesys/bw2_pid/blob/master/allocations.yaml
//Binary (0.0.0.0/4): Binary protocols
//This is a superclass for classes that are generally unreadable in their plain
//form and require translation.
const PONumBinary = 0
const PODFMaskBinary = `0.0.0.0/4`
const PODFBinary = `0.0.0.0`
const POMaskBinary = 4

//Text (64.0.0.0/4): Human readable text
//This is a superclass for classes that are moderately understandable if they
//are read directly in their binary form. Generally these are protocols that
//were designed specifically to be human readable.
const PONumText = 1073741824
const PODFMaskText = `64.0.0.0/4`
const PODFText = `64.0.0.0`
const POMaskText = 4

//Blob (1.0.0.0/8): Blob
//This is a class for schemas that do not use a public encoding format. In
//general it should be avoided. Schemas below this should include the key
//"readme" with a url to a description of the schema that is sufficiently
//detailed to allow for a developer to reverse engineer the protocol if
//required.
const PONumBlob = 16777216
const PODFMaskBlob = `1.0.0.0/8`
const PODFBlob = `1.0.0.0`
const POMaskBlob = 8

//MsgPack (2.0.0.0/8): MsgPack
//This class is for schemas that are represented in MsgPack
const PONumMsgPack = 33554432
const PODFMaskMsgPack = `2.0.0.0/8`
const PODFMsgPack = `2.0.0.0`
const POMaskMsgPack = 8

//CapnP (3.0.0.0/8): Captain Proto
//This class is for captain proto interfaces. Schemas below this should include
//the key "schema" with a url to their .capnp file
const PONumCapnP = 50331648
const PODFMaskCapnP = `3.0.0.0/8`
const PODFCapnP = `3.0.0.0`
const POMaskCapnP = 8

//JSON (65.0.0.0/8): JSON
//This class is for schemas that are represented in JSON
const PONumJSON = 1090519040
const PODFMaskJSON = `65.0.0.0/8`
const PODFJSON = `65.0.0.0`
const POMaskJSON = 8

//XML (66.0.0.0/8): XML
//This class is for schemas that are represented in XML
const PONumXML = 1107296256
const PODFMaskXML = `66.0.0.0/8`
const PODFXML = `66.0.0.0`
const POMaskXML = 8

//YAML (67.0.0.0/8): YAML
//This class is for schemas that are represented in YAML
const PONumYAML = 1124073472
const PODFMaskYAML = `67.0.0.0/8`
const PODFYAML = `67.0.0.0`
const POMaskYAML = 8

//BWRoutingObject (0.0.0.0/24): Bosswave Routing Object
//This class and schema block is reserved for bosswave routing objects
//represented using the full PID.
const PONumBWRoutingObject = 0
const PODFMaskBWRoutingObject = `0.0.0.0/24`
const PODFBWRoutingObject = `0.0.0.0`
const POMaskBWRoutingObject = 24

//LogDict (2.0.1.0/24): LogDict
//This class is for log messages encoded in msgpack
const PONumLogDict = 33554688
const PODFMaskLogDict = `2.0.1.0/24`
const PODFLogDict = `2.0.1.0`
const POMaskLogDict = 24

//TSTaggedMP (2.0.3.0/24): TSTaggedMP
//This superclass describes "ts"->int64 tagged msgpack objects. The timestamp
//is used for merging entries and determining which is later and should be the
//final value.
const PONumTSTaggedMP = 33555200
const PODFMaskTSTaggedMP = `2.0.3.0/24`
const PODFTSTaggedMP = `2.0.3.0`
const POMaskTSTaggedMP = 24

//HamiltonBase (2.0.4.0/24): Hamilton Messages
//This is the base class for messages used with the Hamilton motes. The only
//key guaranteed is "#" that contains a uint16 representation of the serial of
//the mote the message is destined for or originated from.
const PONumHamiltonBase = 33555456
const PODFMaskHamiltonBase = `2.0.4.0/24`
const PODFHamiltonBase = `2.0.4.0`
const POMaskHamiltonBase = 24

//BW2ChatMessages (2.0.7.0/24): BW2ChatMessages
//These are MsgPack dictionaries sent for the BW2Chat program
//(https://github.com/gtfierro/bw2chat)
const PONumBW2ChatMessages = 33556224
const PODFMaskBW2ChatMessages = `2.0.7.0/24`
const PODFBW2ChatMessages = `2.0.7.0`
const POMaskBW2ChatMessages = 24

//HamiltonTelemetry (2.0.4.64/26): Hamilton Telemetry
//This object contains a "#" field for the serial number, as well as possibly
//containing an "A" field with a list of X, Y, and Z accelerometer values. A
//"T" field containing the temperature as an integer in degrees C multiplied by
//10000, and an "L" field containing the illumination in Lux.
const PONumHamiltonTelemetry = 33555520
const PODFMaskHamiltonTelemetry = `2.0.4.64/26`
const PODFHamiltonTelemetry = `2.0.4.64`
const POMaskHamiltonTelemetry = 26

//ROAccessDChainHash (0.0.0.1/32): Access DChain hash
//An access dchain hash
const PONumROAccessDChainHash = 1
const PODFMaskROAccessDChainHash = `0.0.0.1/32`
const PODFROAccessDChainHash = `0.0.0.1`
const POMaskROAccessDChainHash = 32

//ROAccessDChain (0.0.0.2/32): Access DChain
//An access dchain
const PONumROAccessDChain = 2
const PODFMaskROAccessDChain = `0.0.0.2/32`
const PODFROAccessDChain = `0.0.0.2`
const POMaskROAccessDChain = 32

//ROPermissionDChainHash (0.0.0.17/32): Permission DChain hash
//A permission dchain hash
const PONumROPermissionDChainHash = 17
const PODFMaskROPermissionDChainHash = `0.0.0.17/32`
const PODFROPermissionDChainHash = `0.0.0.17`
const POMaskROPermissionDChainHash = 32

//ROPermissionDChain (0.0.0.18/32): Permission DChain
//A permission dchain
const PONumROPermissionDChain = 18
const PODFMaskROPermissionDChain = `0.0.0.18/32`
const PODFROPermissionDChain = `0.0.0.18`
const POMaskROPermissionDChain = 32

//ROAccessDOT (0.0.0.32/32): Access DOT
//An access DOT
const PONumROAccessDOT = 32
const PODFMaskROAccessDOT = `0.0.0.32/32`
const PODFROAccessDOT = `0.0.0.32`
const POMaskROAccessDOT = 32

//ROPermissionDOT (0.0.0.33/32): Permission DOT
//A permission DOT
const PONumROPermissionDOT = 33
const PODFMaskROPermissionDOT = `0.0.0.33/32`
const PODFROPermissionDOT = `0.0.0.33`
const POMaskROPermissionDOT = 32

//ROEntity (0.0.0.48/32): Entity
//An entity
const PONumROEntity = 48
const PODFMaskROEntity = `0.0.0.48/32`
const PODFROEntity = `0.0.0.48`
const POMaskROEntity = 32

//ROOriginVK (0.0.0.49/32): Origin verifying key
//The origin VK of a message that does not contain a PAC
const PONumROOriginVK = 49
const PODFMaskROOriginVK = `0.0.0.49/32`
const PODFROOriginVK = `0.0.0.49`
const POMaskROOriginVK = 32

//ROEntityWKey (0.0.0.50/32): Entity with signing key
//An entity with signing key
const PONumROEntityWKey = 50
const PODFMaskROEntityWKey = `0.0.0.50/32`
const PODFROEntityWKey = `0.0.0.50`
const POMaskROEntityWKey = 32

//RODRVK (0.0.0.51/32): Designated router verifying key
//a 32 byte designated router verifying key
const PONumRODRVK = 51
const PODFMaskRODRVK = `0.0.0.51/32`
const PODFRODRVK = `0.0.0.51`
const POMaskRODRVK = 32

//ROExpiry (0.0.0.64/32): Expiry
//Sets an expiry for the message
const PONumROExpiry = 64
const PODFMaskROExpiry = `0.0.0.64/32`
const PODFROExpiry = `0.0.0.64`
const POMaskROExpiry = 32

//RORevocation (0.0.0.80/32): Revocation
//A revocation for an Entity or a DOT
const PONumRORevocation = 80
const PODFMaskRORevocation = `0.0.0.80/32`
const PODFRORevocation = `0.0.0.80`
const POMaskRORevocation = 32

//BinaryActuation (1.0.1.0/32): Binary actuation
//This payload object is one byte long, 0x00 for off, 0x01 for on.
const PONumBinaryActuation = 16777472
const PODFMaskBinaryActuation = `1.0.1.0/32`
const PODFBinaryActuation = `1.0.1.0`
const POMaskBinaryActuation = 32

//BWMessage (1.0.1.1/32): Packed Bosswave Message
//This object contains an entire signed and encoded bosswave message
const PONumBWMessage = 16777473
const PODFMaskBWMessage = `1.0.1.1/32`
const PODFBWMessage = `1.0.1.1`
const POMaskBWMessage = 32

//Double (1.0.2.0/32): Double
//This payload is an 8 byte long IEEE 754 double floating point value encoded
//in little endian. This should only be used if the semantic meaning is obvious
//in the context, otherwise a PID with a more specific semantic meaning should
//be used.
const PONumDouble = 16777728
const PODFMaskDouble = `1.0.2.0/32`
const PODFDouble = `1.0.2.0`
const POMaskDouble = 32

//Wavelet (1.0.6.1/32): Wavelet binary
//This object contains a BOSSWAVE Wavelet
const PONumWavelet = 16778753
const PODFMaskWavelet = `1.0.6.1/32`
const PODFWavelet = `1.0.6.1`
const POMaskWavelet = 32

//SpawnpointLog (2.0.2.0/32): Spawnpoint stdout
//This contains stdout data from a spawnpoint container. It is a msgpacked
//dictionary that contains a "service" key, a "time" key (unix nano timestamp)
//and a "contents" key and a "spalias" key.
const PONumSpawnpointLog = 33554944
const PODFMaskSpawnpointLog = `2.0.2.0/32`
const PODFSpawnpointLog = `2.0.2.0`
const POMaskSpawnpointLog = 32

//SMetadata (2.0.3.1/32): Simple Metadata entry
//This contains a simple "val" string and "ts" int64 metadata entry. The key is
//determined by the URI. Other information MAY be present in the msgpacked
//object. The timestamp is used for merging metadata entries.
const PONumSMetadata = 33555201
const PODFMaskSMetadata = `2.0.3.1/32`
const PODFSMetadata = `2.0.3.1`
const POMaskSMetadata = 32

//HSBLightMessage (2.0.5.1/32): HSBLight Message
//This object may contain "hue", "saturation", "brightness" fields with a float
//from 0 to 1. It may also contain an "on" key with a boolean. Omitting fields
//leaves them at their previous state.
const PONumHSBLightMessage = 33555713
const PODFMaskHSBLightMessage = `2.0.5.1/32`
const PODFHSBLightMessage = `2.0.5.1`
const POMaskHSBLightMessage = 32

//InterfaceDescriptor (2.0.6.1/32): InterfaceDescriptor
//This object is used to describe an interface. It contains "uri",
//"iface","svc","namespace" "prefix" and "metadata" keys.
const PONumInterfaceDescriptor = 33555969
const PODFMaskInterfaceDescriptor = `2.0.6.1/32`
const PODFInterfaceDescriptor = `2.0.6.1`
const POMaskInterfaceDescriptor = 32

//BW2Chat_CreateRoomMessage (2.0.7.1/32): BW2Chat_CreateRoomMessage
//A dictionary with a single key "Name" indicating the room to be created. This
//will likely be deprecated.
const PONumBW2Chat_CreateRoomMessage = 33556225
const PODFMaskBW2Chat_CreateRoomMessage = `2.0.7.1/32`
const PODFBW2Chat_CreateRoomMessage = `2.0.7.1`
const POMaskBW2Chat_CreateRoomMessage = 32

//BW2Chat_ChatMessage (2.0.7.2/32): BW2Chat_ChatMessage
//A textual message to be sent to all members of a chatroom. This is a
//dictionary with three keys: 'Room', the name of the room to publish to (this
//is actually implicit in the publishing), 'From', the alias you are using for
//the chatroom, and 'Message', the actual string to be displayed to all users
//in the room.
const PONumBW2Chat_ChatMessage = 33556226
const PODFMaskBW2Chat_ChatMessage = `2.0.7.2/32`
const PODFBW2Chat_ChatMessage = `2.0.7.2`
const POMaskBW2Chat_ChatMessage = 32

//BW2Chat_JoinRoom (2.0.7.3/32): BW2Chat_JoinRoom
//Notify users in the chatroom that you have joined. Dictionary with a single
//key "Alias" that has a value of your nickname
const PONumBW2Chat_JoinRoom = 33556227
const PODFMaskBW2Chat_JoinRoom = `2.0.7.3/32`
const PODFBW2Chat_JoinRoom = `2.0.7.3`
const POMaskBW2Chat_JoinRoom = 32

//BW2Chat_LeaveRoom (2.0.7.4/32): BW2Chat_LeaveRoom
//Notify users in the chatroom that you have left. Dictionary with a single key
//"Alias" that has a value of your nickname
const PONumBW2Chat_LeaveRoom = 33556228
const PODFMaskBW2Chat_LeaveRoom = `2.0.7.4/32`
const PODFBW2Chat_LeaveRoom = `2.0.7.4`
const POMaskBW2Chat_LeaveRoom = 32

//String (64.0.1.0/32): String
//A plain string with no rigid semantic meaning. This can be thought of as a
//print statement. Anything that has semantic meaning like a process log should
//use a different schema.
const PONumString = 1073742080
const PODFMaskString = `64.0.1.0/32`
const PODFString = `64.0.1.0`
const POMaskString = 32

//FMDIntentString (64.0.1.1/32): FMD Intent String
//A plain string used as an intent for the follow-me display service.
const PONumFMDIntentString = 1073742081
const PODFMaskFMDIntentString = `64.0.1.1/32`
const PODFFMDIntentString = `64.0.1.1`
const POMaskFMDIntentString = 32

//AccountBalance (64.0.1.2/32): Account balance
//A comma seperated representation of an account and its balance as
//addr,decimal,human_readable. For example
//0x49b1d037c33fdaad75d2532cd373fb5db87cc94c,57203431159181996982272,57203.4311
//Ether  . Be careful in that the decimal representation will frequently be
//bigger than an int64.
const PONumAccountBalance = 1073742082
const PODFMaskAccountBalance = `64.0.1.2/32`
const PODFAccountBalance = `64.0.1.2`
const POMaskAccountBalance = 32

//SpawnpointConfig (67.0.2.0/32): SpawnPoint config
//A configuration file for SpawnPoint (github.com/immesys/spawnpoint)
const PONumSpawnpointConfig = 1124073984
const PODFMaskSpawnpointConfig = `67.0.2.0/32`
const PODFSpawnpointConfig = `67.0.2.0`
const POMaskSpawnpointConfig = 32

//SpawnpointHeartbeat (67.0.2.1/32): SpawnPoint heartbeat
//A heartbeat message from spawnpoint
const PONumSpawnpointHeartbeat = 1124073985
const PODFMaskSpawnpointHeartbeat = `67.0.2.1/32`
const PODFSpawnpointHeartbeat = `67.0.2.1`
const POMaskSpawnpointHeartbeat = 32

