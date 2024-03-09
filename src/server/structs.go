package main

type PacketHeader struct {
	PacketFormat            uint16  `json:"packetFormat"`
	GameYear                uint8   `json:"gameYear"`
	GameMajorVersion        uint8   `json:"gameMajorVersion"`
	GameMinorVersion        uint8   `json:"gameMinorVersion"`
	PacketVersion           uint8   `json:"packetVersion"`
	PacketID                uint8   `json:"packetID"`
	SessionUID              uint64  `json:"sessionUID"`
	SessionTime             float32 `json:"sessionTime"`
	FrameIdentifier         uint32  `json:"frameIdentifier"`
	PlayerCarIndex          uint8   `json:"playerCarIndex"`
	SecondaryPlayerCarIndex uint8   `json:"secondaryPlayerCarIndex"`
}

type CarMotionData struct {
	WorldPositionX     float32 `json:"worldPositionX"`
	WorldPositionY     float32 `json:"worldPositionY"`
	WorldPositionZ     float32 `json:"worldPositionZ"`
	WorldVelocityX     float32 `json:"worldVelocityX"`
	WorldVelocityY     float32 `json:"worldVelocityY"`
	WorldVelocityZ     float32 `json:"worldVelocityZ"`
	WorldForwardX      uint16  `json:"worldForwardX"`
	WorldForwardY      uint16  `json:"worldForwardY"`
	WorldForwardZ      uint16  `json:"worldForwardZ"`
	WorldRightDirX     uint16  `json:"worldRightDirX"`
	WorldRightDirY     uint16  `json:"worldRightDirY"`
	WorldRightDirZ     uint16  `json:"worldRightDirZ"`
	GForceLateral      float32 `json:"gForceLateral"`
	GForceLongitudinal float32 `json:"gForceLongitudinal"`
	GForceVertical     float32 `json:"gForceVertical"`
	Yaw                float32 `json:"yaw"`
	Pitch              float32 `json:"pitch"`
	Roll               float32 `json:"roll"`
}

type PacketMotionData struct {
	Header        PacketHeader  `json:"header"`
	CarMotionData CarMotionData `json:"carMotionData"`
}
