package packets

import (
	"github.com/google/uuid"
	"github.com/junglemc/JungleTree/pkg/chat"
	"github.com/junglemc/JungleTree/pkg/crafting"
	"github.com/junglemc/JungleTree/pkg/util"
	"github.com/junglemc/JungleTree/pkg/world"
)

type ClientboundSpawnEntityPacket struct {
	EntityId   int32 `type:"varint"`
	ObjectUUID uuid.UUID
	Type       int32 `type:"varint"`
	X          float64
	Y          float64
	Z          float64
	Pitch      int8
	Yaw        int8
	ObjectData int32
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

type ClientboundSpawnEntityExperienceOrbPacket struct {
	EntityId int32 `type:"varint"`
	X        float64
	Y        float64
	Z        float64
	Count    int16
}

type ClientboundSpawnEntityLivingPacket struct {
	EntityId   int32 `type:"varint"`
	EntityUUID uuid.UUID
	Type       int32 `type:"varint"`
	X          float64
	Y          float64
	Z          float64
	Yaw        int8
	Pitch      int8
	HeadPitch  int8
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

type ClientboundSpawnEntityPaintingPacket struct {
	EntityId   int32 `type:"varint"`
	EntityUUID uuid.UUID
	Title      int32 `type:"varint"`
	Location   world.BlockPosition
	Direction  byte
}

type ClientboundPlaySpawnPlayerPacket struct {
	EntityId   int32 `type:"varint"`
	PlayerUUID uuid.UUID
	X          float64
	Y          float64
	Z          float64
	Yaw        int8
	Pitch      int8
}

type ClientboundPlayEntityAnimationPacket struct {
	EntityId  int32 `type:"varint"`
	Animation byte
}

type ClientboundPlayStatisticsPacket struct {
	Entries []interface{} ``
}

type ClientboundPlayAcknowledgePlayerDiggingPacket struct {
	Location   world.BlockPosition
	Block      int32 `type:"varint"`
	Status     int32 `type:"varint"`
	Successful bool
}

type ClientboundPlayBlockBreakAnimationPacket struct {
	EntityId     int32 `type:"varint"`
	Location     world.BlockPosition
	DestroyStage byte
}

type ClientboundPlayBlockEntityDataPacket struct {
	Location world.BlockPosition
	Action   byte
	Data     map[string]interface{} `type:"nbt"`
}

type ClientboundPlayBlockActionPacket struct {
	Location        world.BlockPosition
	ActionID        byte
	ActionParameter byte
	BlockId         int32 `type:"varint"`
}

type ClientboundJoinGamePacket struct {
	EntityId            int32
	IsHardcore          bool
	GameMode            util.GameMode
	PreviousGameMode    int8
	WorldNames          []string
	DimensionCodec      interface{} `type:"nbt"`
	Dimension           interface{} `type:"nbt"`
	WorldName           string
	HashedSeed          int64
	MaxPlayers          int32 `type:"varint"`
	ViewDistance        int32 `type:"varint"`
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool
	IsFlat              bool
}

type ClientboundPlayEntityEffect struct {
	EntityId      int32 `type:"varint"`
	EffectId      int8
	Amplifier     int8
	Duration      int32 `type:"varint"`
	HideParticles int8
}

type ClientboundPlayNamedSoundEffect struct {
	SoundName     string
	SoundCategory int32 `type:"varint"`
	X             int32
	Y             int32
	Z             int32
	Volume        float32
	Pitch         float32
}

type ClientboundPlayBlockChange struct {
	Location world.BlockPosition
	Type     int32 `type:"varint"`
}

type ClientboundPlayPosition struct {
	X          float64
	Y          float64
	Z          float64
	Yaw        float32
	Pitch      float32
	Flags      int8
	TeleportId int32 `type:"varint"`
}

type ClientboundPlayFacePlayer struct {
	FeetEyes       int32 `type:"varint"`
	X              float64
	Y              float64
	Z              float64
	IsEntity       bool
	EntityId       interface{}
	EntityFeetEyes interface{}
}

type ClientboundPlayEntityVelocity struct {
	EntityId  int32 `type:"varint"`
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

type ClientboundPlayEntitySoundEffect struct {
	SoundId       int32 `type:"varint"`
	SoundCategory int32 `type:"varint"`
	EntityId      int32 `type:"varint"`
	Volume        float32
	Pitch         float32
}

type ClientboundPlayStopSound struct {
	Flags  int8
	Source interface{}
	Sound  interface{}
}

type ClientboundPlayWorldParticles struct {
	ParticleId   int32
	LongDistance bool
	X            float64
	Y            float64
	Z            float64
	OffsetX      float32
	OffsetY      float32
	OffsetZ      float32
	ParticleData float32
	Particles    int32
	Data         interface{}
}

type ClientboundPlayEntityLook struct {
	EntityId int32 `type:"varint"`
	Yaw      int8
	Pitch    int8
	OnGround bool
}

type ClientboundPlayUpdateLight struct {
	ChunkX              int32 `type:"varint"`
	ChunkZ              int32 `type:"varint"`
	TrustEdges          bool
	SkyLightMask        int32 `type:"varint"`
	BlockLightMask      int32 `type:"varint"`
	EmptySkyLightMask   int32 `type:"varint"`
	EmptyBlockLightMask int32 `type:"varint"`
	Data                interface{}
}

type ClientboundPlayMap struct {
	ItemDamage       int32 `type:"varint"`
	Scale            int8
	TrackingPosition bool
	Locked           bool
	Icons            []interface{} ``
	Columns          int8
	Rows             interface{}
	X                interface{}
	Y                interface{}
	Data             interface{}
}

type ClientboundPlayResourcePackSend struct {
	Url  string
	Hash string
}

type ClientboundPlayScoreboardDisplayObjective struct {
	Position int8
	Name     string
}

type ClientboundPlayExperience struct {
	ExperienceBar   float32
	Level           int32 `type:"varint"`
	TotalExperience int32 `type:"varint"`
}

type ClientboundPlayUpdateHealth struct {
	Health         float32
	Food           int32 `type:"varint"`
	FoodSaturation float32
}

type ClientboundPlayBossBar struct {
	EntityUUID uuid.UUID
	Action     int32 `type:"varint"`
	Title      interface{}
	Health     interface{}
	Color      interface{}
	Dividers   interface{}
	Flags      interface{}
}

type ClientboundPlayEntityTeleport struct {
	EntityId int32 `type:"varint"`
	X        float64
	Y        float64
	Z        float64
	Yaw      int8
	Pitch    int8
	OnGround bool
}

type ClientboundPlayEntity struct {
	EntityId int32 `type:"varint"`
}

type ClientboundPlayerAbilitiesPacket struct {
	Flags        byte
	FlyingSpeed  float32
	WalkingSpeed float32
}

type ClientboundPlayCombatEvent struct {
	Event    int32 `type:"varint"`
	Duration interface{}
	PlayerId interface{}
	EntityId interface{}
	Message  interface{}
}

type ClientboundPlayUpdateViewDistance struct {
	ViewDistance int32 `type:"varint"`
}

type ClientboundPlaySetSlot struct {
	WindowId int8
	Slot     int16
	Item     interface{} // TODO: Slot
}

type ClientboundPlayEntityHeadRotation struct {
	EntityId int32 `type:"varint"`
	HeadYaw  int8
}

type ClientboundPlayNbtQueryResponse struct {
	TransactionId int32                  `type:"varint"`
	NBT           map[string]interface{} `type:"nbt"`
}

type ClientboundPlayTags struct {
	BlockTags  interface{}
	ItemTags   interface{}
	FluidTags  interface{}
	EntityTags interface{}
}

type ClientboundPlayPlayerInfo struct {
	Action int32         `type:"varint"`
	Data   []interface{} ``
}

type ClientboundPlayUnlockRecipes struct {
	Action                int32 `type:"varint"`
	CraftingBookOpen      bool
	FilteringCraftable    bool
	SmeltingBookOpen      bool
	FilteringSmeltable    bool
	BlastFurnaceOpen      bool
	FilteringBlastFurnace bool
	SmokerBookOpen        bool
	FilteringSmoker       bool
	Recipes1              []interface{} ``
	Recipes2              interface{}
}

type ClientboundPlayCamera struct {
	CameraId int32 `type:"varint"`
}

type ClientboundHeldItemChangePacket struct {
	Slot byte
}

type ClientboundPlayTeams struct {
	Team              string
	Mode              int8
	Name              interface{}
	FriendlyFire      interface{}
	NameTagVisibility interface{}
	CollisionRule     interface{}
	Formatting        interface{}
	Prefix            interface{}
	Suffix            interface{}
	Players           interface{}
}

type ClientboundPlayGameStateChange struct {
	Reason   uint8
	GameMode float32
}

type ClientboundPlayCraftRecipeResponse struct {
	WindowId int8
	Recipe   string
}

type ClientboundPlayDeclareCommands struct {
	Nodes     []interface{} ``
	RootIndex int32         `type:"varint"`
}

type ClientboundPlaySetCooldown struct {
	ItemID        int32 `type:"varint"`
	CooldownTicks int32 `type:"varint"`
}

type ClientboundPlaySelectAdvancementTab struct {
	Id interface{}
}

type ClientboundPlayTransaction struct {
	WindowId int8
	Action   int16
	Accepted bool
}

type ClientboundPlayOpenHorseWindow struct {
	WindowId uint8
	NbSlots  int32 `type:"varint"`
	EntityId int32
}

type ClientboundPlayMapChunk struct {
	X             int32
	Z             int32
	GroundUp      bool
	BitMap        int32 `type:"varint"`
	Heightmaps    interface{}
	Biomes        interface{}
	ChunkData     []byte        ``
	BlockEntities []interface{} ``
}

type ClientboundPlayOpenBook struct {
	Hand int32 `type:"varint"`
}

type ClientboundPlaySetPassengers struct {
	EntityId   int32         `type:"varint"`
	Passengers []interface{} ``
}

type ClientboundPlayScoreboardScore struct {
	ItemName  string
	Action    int8
	ScoreName string
	Value     interface{}
}

type ClientboundPlayWindowItems struct {
	WindowId uint8
	Items    []interface{} `size:"int16"`
}

type ClientboundPlayEntityMoveLook struct {
	EntityId int32 `type:"varint"`
	DX       int16
	DY       int16
	DZ       int16
	Yaw      int8
	Pitch    int8
	OnGround bool
}

type ClientboundPlayRemoveEntityEffect struct {
	EntityId int32 `type:"varint"`
	EffectId int8
}

type ClientboundPlayUpdateTime struct {
	Age  int64
	Time int64
}

type ClientboundPlayPlayerlistHeader struct {
	Header string
	Footer string
}

type ClientboundPlayChat struct {
	Message  string
	Position int8
	Sender   uuid.UUID
}

type ClientboundPlayTabComplete struct {
	TransactionId int32         `type:"varint"`
	Start         int32         `type:"varint"`
	Length        int32         `type:"varint"`
	Matches       []interface{} ``
}

type ClientboundPlayMultiBlockChange struct {
	ChunkCoordinates interface{}
	NotTrustEdges    bool
	Records          []interface{} ``
}

type ClientboundPlayAdvancements struct {
	Reset              bool
	AdvancementMapping []interface{} ``
	Identifiers        []interface{} ``
	ProgressMapping    []interface{} ``
}

type ClientboundServerDifficultyPacket struct {
	Difficulty       util.Difficulty
	DifficultyLocked bool
}

type ClientboundPlayEntityDestroy struct {
	EntityIds []interface{} ``
}

type ClientboundPlayExplosion struct {
	X                    float32
	Y                    float32
	Z                    float32
	Radius               float32
	AffectedBlockOffsets []interface{} `size:"int32"`
	PlayerMotionX        float32
	PlayerMotionY        float32
	PlayerMotionZ        float32
}

type ClientboundPlayTradeList struct {
	WindowId          int32         `type:"varint"`
	Trades            []interface{} `size:"byte"`
	VillagerLevel     int32         `type:"varint"`
	Experience        int32         `type:"varint"`
	IsRegularVillager bool
	CanRestock        bool
}

type ClientboundPlayOpenWindow struct {
	WindowId      int32 `type:"varint"`
	InventoryType int32 `type:"varint"`
	WindowTitle   string
}

type ClientboundPlayCraftProgressBar struct {
	WindowId uint8
	Property int16
	Value    int16
}

type ClientboundPlayKickDisconnect struct {
	Reason *chat.Message
}

type ClientboundPlayAttachEntity struct {
	EntityId  int32
	VehicleId int32
}

type ClientboundPlayWorldEvent struct {
	EffectId int32
	Location world.BlockPosition
	Data     int32
	Global   bool
}

type ClientboundPlayEntityMetadata struct {
	EntityId int32       `type:"varint"`
	Metadata interface{} // TODO: EntityMetadata
}

type ClientboundPlayUnloadChunk struct {
	ChunkX int32
	ChunkZ int32
}

type ClientboundPlayKeepAlive struct {
	KeepAliveId int64
}

type ClientboundPlayRelEntityMove struct {
	EntityId int32 `type:"varint"`
	DX       int16
	DY       int16
	DZ       int16
	OnGround bool
}

type ClientboundPlayOpenSignEntity struct {
	Location world.BlockPosition
}

type ClientboundPlayRespawn struct {
	Dimension        interface{}
	WorldName        string
	HashedSeed       int64
	Gamemode         uint8
	PreviousGamemode uint8
	IsDebug          bool
	IsFlat           bool
	CopyMetadata     bool
}

type ClientboundPlayWorldBorder struct {
	Action         int32 `type:"varint"`
	Radius         interface{}
	X              interface{}
	Z              interface{}
	OldRadius      interface{}
	NewRadius      interface{}
	Speed          interface{}
	PortalBoundary interface{}
	WarningTime    interface{}
	WarningBlocks  interface{}
}

type ClientboundDeclareRecipesPacket struct {
	Recipes []*crafting.Recipe
}

type ClientboundPlaySpawnPosition struct {
	Location world.BlockPosition
}

type ClientboundPlayEntityEquipment struct {
	EntityId   int32 `type:"varint"`
	Equipments interface{}
}

type ClientboundPlayUpdateViewPosition struct {
	ChunkX int32 `type:"varint"`
	ChunkZ int32 `type:"varint"`
}

type ClientboundPlayScoreboardObjective struct {
	Name        string
	Action      int8
	DisplayText interface{}
	Type        interface{}
}

type ClientboundPlayEntityUpdateAttributes struct {
	EntityId   int32         `type:"varint"`
	Properties []interface{} `size:"int32"`
}

type ClientboundPluginMessagePacket struct {
	Channel util.Identifier
	Data    []byte `size:"infer"`
}

type ClientboundPlayVehicleMove struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

type ClientboundPlayEntityStatus struct {
	EntityId     int32
	EntityStatus int8
}

type ClientboundPlayTitle struct {
	Action  int32 `type:"varint"`
	Text    interface{}
	FadeIn  interface{}
	Stay    interface{}
	FadeOut interface{}
}

type ClientboundPlaySoundEffect struct {
	SoundId       int32 `type:"varint"`
	SoundCategory int32 `type:"varint"`
	X             int32
	Y             int32
	Z             int32
	Volume        float32
	Pitch         float32
}

type ClientboundPlayCollect struct {
	CollectedEntityId int32 `type:"varint"`
	CollectorEntityId int32 `type:"varint"`
	PickupItemCount   int32 `type:"varint"`
}

type ClientboundPlayCloseWindow struct {
	WindowId uint8
}
