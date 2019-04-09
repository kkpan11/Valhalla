package opcode

const (
	SendLoginResponce               byte = 0x01
	SendLoginWorldMeta              byte = 0x03
	SendLoginPinRegister            byte = 0x07 // Add 1 byte, 1 = register byte a pin
	SendLoginPinStuff               byte = 0x08 // Setting byte pin good
	SendLoginWorldList              byte = 0x09
	SendLoginRestarter              byte = 0x15
	SendLoginCharacterData          byte = 0x0A
	SendLoginCharacterMigrate       byte = 0x0B
	SendLoginNameCheckResult        byte = 0x0C
	SendLoginNewCharacterGood       byte = 0x0D
	SendLoginDeleteCharacter        byte = 0x0E
	SendChannelInventoryOperation   byte = 0x18
	SendChannelStatChange           byte = 0x1A
	SendChannelSkillRecordUpdate    byte = 0x1D
	SendChannelInfoMessage          byte = 0x20
	SendChannelLieDetectorTest      byte = 0x23
	SendChannelAvatarInfoWindow     byte = 0x2c
	SendChannelPartyInfo            byte = 0x2D
	SendChannelBroadcastMessage     byte = 0x32
	SendChannelWarpToMap            byte = 0x36
	SendChannelPortalClosed         byte = 0x3A
	SendChannelBubblessChat         byte = 0x3D
	SendChannelWhisper              byte = 0x3E
	SendChannelEmployee             byte = 0x43
	SendChannelQuizQAndA            byte = 0x44
	SendChannelCharacterEnterField  byte = 0x4E
	SendChannelCharacterLeaveField  byte = 0x4F
	SendChannelAllChatMsg           byte = 0x51
	SendChannelPlayerMovement       byte = 0x65
	SendChannelPlayerUseMeleeSkill  byte = 0x66
	SendChannelPlayerUseRangedSkill byte = 0x67
	SendChannelPlayerUseMagicSkill  byte = 0x68
	SendChannelPlayerTakeDmg        byte = 0x6B
	SendChannelPlayerEmoticon       byte = 0x6C
	SendChannelPlayerChangeAvatar   byte = 0x6F
	SendChannelPlayerAnimation      byte = 0x70
	SendChannelLevelUpAnimation     byte = 0x79
	SendChannelShowMob              byte = 0x86
	SendChannelRemoveMob            byte = 0x87
	SendChannelControlMob           byte = 0x88
	SendChannelMoveMob              byte = 0x8A
	SendChannelControlMobAck        byte = 0x8B
	SendChannelMobChangeHP          byte = 0x91
	SendChannelNpcShow              byte = 0x97
	SendChannelNpcRemove            byte = 0x98
	SendChannelNpcControl           byte = 0x99
	SendChannelNpcMovement          byte = 0x9B
	SendChannelSpawnDoor            byte = 0xB1
	SendChannelRemoveDoor           byte = 0xB2
	SendChannelNpcDialogueBox       byte = 0xC5
	SendChannelNpcShop              byte = 0xC8
	SendChannelNpcShopResult        byte = 0xC9
	SendChannelNpcStorage           byte = 0xCD
	SendChannelRoom                 byte = 0xDC
	SendChannelRoomBox              byte = 0x52
)
