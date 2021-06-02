package packet

import (
	"github.com/google/uuid"
	"github.com/junglemc/JungleTree"
	"github.com/junglemc/JungleTree/chat"
	"github.com/junglemc/JungleTree/world"
)

type ServerboundPlayEnchantItem struct {
	WindowId    int8
	Enchantment int8
}

type ServerboundPlayFlying struct {
	OnGround bool
}

type ServerboundPlayPickItem struct {
	Slot int32 `type:"varint"`
}

type ServerboundAdvancementTabPacket struct {
	Action int32  `type:"varint"`
	TabId  string `optional:"Action=0"`
}

type ServerboundPlaySelectTrade struct {
	Slot int32 `type:"varint"`
}

type ServerboundPlayUpdateStructureBlock struct {
	Location world.BlockPosition
	Action   int32 `type:"varint"`
	Mode     int32 `type:"varint"`
	Name     string
	OffsetX  uint8
	OffsetY  uint8
	OffsetZ  uint8
	SizeX    uint8
	SizeY    uint8
	SizeZ    uint8
	Mirror   int32 `type:"varint"`
	Rotation  int32 `type:"varint"`
	Metadata  string
	Integrity float32
	Seed      int32 `type:"varint"`
	Flags     uint8
}

type ServerboundPlaySpectate struct {
	Target uuid.UUID
}

type ServerboundPlayCraftRecipeRequest struct {
	WindowId int8
	Recipe   string
	MakeAll  bool
}

type ServerboundPlayResourcePackReceive struct {
	Result int32 `type:"varint"`
}

type ServerboundPlaySetBeaconEffect struct {
	PrimaryEffect   int32 `type:"varint"`
	SecondaryEffect int32 `type:"varint"`
}

type ServerboundPlayArmAnimation struct {
	Hand int32 `type:"varint"`
}

type ServerboundPlayCloseWindow struct {
	WindowId uint8
}

type ServerboundPlayEditBook struct {
	NewBook interface{} // TODO: Slot
	Signing bool
	Hand    int32 `type:"varint"`
}

type ServerboundInteractEntityPacket struct {
	Target   int32   `type:"varint"`
	Mouse    int32   `type:"varint"`
	X        float32 `optional:"Type=2"`
	Y        float32 `optional:"Type=2"`
	Z        float32 `optional:"Type=2"`
	Hand     int32   `type:"varint" optional:"Type=0=2"`
	Sneaking bool
}

type ServerboundPlayBlockDig struct {
	Status   int8
	Location world.BlockPosition
	Face     int8
}

type ServerboundPlayEntityAction struct {
	EntityId  int32 `type:"varint"`
	ActionId  int32 `type:"varint"`
	JumpBoost int32 `type:"varint"`
}

type ServerboundPlayUseItem struct {
	Hand int32 `type:"varint"`
}

type ServerboundPlaySetDifficulty struct {
	NewDifficulty uint8
}

type ServerboundPlayQueryEntityNbt struct {
	TransactionId int32 `type:"varint"`
	EntityId      int32 `type:"varint"`
}

type ServerboundPlayLockDifficulty struct {
	Locked bool
}

type ServerboundPlayPositionLook struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

type ServerboundPlayLook struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

type ServerboundPlayRecipeBook struct {
	BookId       int32 `type:"varint"`
	BookOpen     bool
	FilterActive bool
}

type ServerboundPlayUpdateCommandBlockMinecart struct {
	EntityId    int32 `type:"varint"`
	Command     string
	TrackOutput bool
}

type ServerboundPlayTeleportConfirm struct {
	TeleportId int32 `type:"varint"`
}

type ServerboundPlayTabComplete struct {
	TransactionId int32 `type:"varint"`
	Text          string
}

type ServerboundPlaySteerBoat struct {
	LeftPaddle  bool
	RightPaddle bool
}

type ServerboundPlayDisplayedRecipe struct {
	RecipeId string
}

type ServerboundPlayNameItem struct {
	Name string
}

type ServerboundPlaySetCreativeSlot struct {
	Slot int16
	Item interface{} // TODO: Slot
}

type ServerboundPlayUpdateSign struct {
	Location world.BlockPosition
	Text1    string
	Text2    string
	Text3    string
	Text4    string
}

type ServerboundPluginMessagePacket struct {
	Channel JungleTree.Identifier
	Data    []byte `size:"infer"`
}

type ServerboundPlayVehicleMove struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

type ServerboundPlaySteerVehicle struct {
	Sideways float32
	Forward  float32
	Jump     uint8
}

type ServerboundPlayBlockPlace struct {
	Hand        int32 `type:"varint"`
	Location    world.BlockPosition
	Direction   int32 `type:"varint"`
	CursorX     float32
	CursorY     float32
	CursorZ     float32
	InsideBlock bool
}

type ServerboundPlayChat struct {
	Message string
}

type ServerboundPlayClientCommand struct {
	ActionId int32 `type:"varint"`
}

type ServerboundPlayTransaction struct {
	WindowId int8
	Action   int16
	Accepted bool
}

type ServerboundPlayGenerateStructure struct {
	Location    world.BlockPosition
	Levels      int32 `type:"varint"`
	KeepJigsaws bool
}

type ServerboundPlayHeldItemSlot struct {
	SlotId int16
}

type ServerboundPlayUpdateCommandBlock struct {
	Location world.BlockPosition
	Command  string
	Mode     int32 `type:"varint"`
	Flags    uint8
}

type ServerboundPlayQueryBlockNbt struct {
	TransactionId int32 `type:"varint"`
	Location      world.BlockPosition
}

type ServerboundClientSettingsPacket struct {
	Locale            string
	ViewDistance      byte
	ChatMode          *chat.Mode
	ChatColorsEnabled bool
	SkinParts         byte
	MainHand          *JungleTree.Hand
}

type ServerboundPlayWindowClick struct {
	WindowId    uint8
	Slot        int16
	MouseButton int8
	Action      int16
	Mode        int8
	Item        interface{} // TODO: Slot
}

type ServerboundPlayKeepAlive struct {
	KeepAliveId int64
}

type ServerboundPlayPosition struct {
	X        float64
	Y        float64
	Z        float64
	OnGround bool
}

type ServerboundPlayAbilities struct {
	Flags int8
}

type ServerboundPlayUpdateJigsawBlock struct {
	Location   world.BlockPosition
	Name       string
	Target     string
	Pool       string
	FinalState string
	JointType  string
}
