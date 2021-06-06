package protocol

import (
	"fmt"
	. "reflect"

	. "github.com/junglemc/JungleTree/internal/pkg/net/packets"
)

var Registry = registry{
	Handshake: func() (clientbound, serverbound map[int32]Type) {
		return HandshakeClientboundIds, HandshakeServerboundIds
	},
	Status: func() (clientbound, serverbound map[int32]Type) {
		return StatusClientboundIds, StatusServerboundIds
	},
	Login: func() (clientbound, serverbound map[int32]Type) {
		return LoginClientboundIds, LoginServerboundIds
	},
	Play: func() (clientbound, serverbound map[int32]Type) {
		return PlayClientboundIds, PlayServerboundIds
	},
}

type registryMap func() (clientbound map[int32]Type, serverbound map[int32]Type)

type registry struct {
	Handshake registryMap
	Status    registryMap
	Login     registryMap
	Play      registryMap
}

func (r *registry) ClientboundID(t Type, p Protocol) int32 {
	var clientbound map[int32]Type

	switch p {
	case Handshake:
		clientbound, _ = r.Handshake()
	case Status:
		clientbound, _ = r.Status()
	case Login:
		clientbound, _ = r.Login()
	case Play:
		clientbound, _ = r.Play()
	}

	for id, pkt := range clientbound {
		if pkt.Name() == t.Name() {
			return id
		}
	}

	panic("not found") // TODO: Cleanup
}

func (r *registry) ServerboundType(id int32, p Protocol) Type {
	var serverbound map[int32]Type

	switch p {
	case Handshake:
		_, serverbound = r.Handshake()
	case Status:
		_, serverbound = r.Status()
	case Login:
		_, serverbound = r.Login()
	case Play:
		_, serverbound = r.Play()
	}

	result := serverbound[id]
	if result == nil {
		panic(fmt.Sprintf("not found: packetID=0x%02X", id))
	}
	return result
}

var HandshakeClientboundIds = map[int32]Type{}

var HandshakeServerboundIds = map[int32]Type{
	0x00: TypeOf(ServerboundHandshakePacket{}),
	0xFE: TypeOf(ServerboundHandshakeLegacyPingPacket{}),
}

var StatusClientboundIds = map[int32]Type{
	0x00: TypeOf(ClientboundStatusResponsePacket{}),
	0x01: TypeOf(ClientboundStatusPongPacket{}),
}

var StatusServerboundIds = map[int32]Type{
	0x00: TypeOf(ServerboundStatusRequestPacket{}),
	0x01: TypeOf(ServerboundStatusPingPacket{}),
}

var LoginClientboundIds = map[int32]Type{
	0x00: TypeOf(ClientboundLoginDisconnectPacket{}),
	0x01: TypeOf(ClientboundLoginEncryptionRequest{}),
	0x02: TypeOf(ClientboundLoginSuccess{}),
	0x03: TypeOf(ClientboundLoginCompressionPacket{}),
	0x04: TypeOf(ClientboundLoginPluginRequest{}),
}

var LoginServerboundIds = map[int32]Type{
	0x00: TypeOf(ServerboundLoginStartPacket{}),
	0x01: TypeOf(ServerboundLoginEncryptionResponsePacket{}),
	0x02: TypeOf(ServerboundLoginPluginResponsePacket{}),
}

var PlayClientboundIds = map[int32]Type{
	0x00: TypeOf(ClientboundSpawnEntityPacket{}),
	0x01: TypeOf(ClientboundSpawnEntityExperienceOrbPacket{}),
	0x02: TypeOf(ClientboundSpawnEntityLivingPacket{}),
	0x03: TypeOf(ClientboundSpawnEntityPaintingPacket{}),
	0x04: TypeOf(ClientboundPlaySpawnPlayerPacket{}),
	0x05: TypeOf(ClientboundPlayEntityAnimationPacket{}),
	0x06: TypeOf(ClientboundPlayStatisticsPacket{}),
	0x07: TypeOf(ClientboundPlayAcknowledgePlayerDiggingPacket{}),
	0x08: TypeOf(ClientboundPlayBlockBreakAnimationPacket{}),
	0x09: TypeOf(ClientboundPlayBlockEntityDataPacket{}),
	0x0A: TypeOf(ClientboundPlayBlockActionPacket{}),
	0x0B: TypeOf(ClientboundPlayBlockChange{}),
	0x0C: TypeOf(ClientboundPlayBossBar{}),
	0x0D: TypeOf(ClientboundServerDifficultyPacket{}),
	0x0E: TypeOf(ClientboundPlayChat{}),
	0x0F: TypeOf(ClientboundPlayTabComplete{}),
	0x10: TypeOf(ClientboundPlayDeclareCommands{}),
	0x11: TypeOf(ClientboundPlayTransaction{}),
	0x12: TypeOf(ClientboundPlayCloseWindow{}),
	0x13: TypeOf(ClientboundPlayWindowItems{}),
	0x14: TypeOf(ClientboundPlayCraftProgressBar{}),
	0x15: TypeOf(ClientboundPlaySetSlot{}),
	0x16: TypeOf(ClientboundPlaySetCooldown{}),
	0x17: TypeOf(ClientboundPluginMessagePacket{}),
	0x18: TypeOf(ClientboundPlayNamedSoundEffect{}),
	0x19: TypeOf(ClientboundPlayKickDisconnect{}),
	0x1A: TypeOf(ClientboundPlayEntityStatus{}),
	0x1B: TypeOf(ClientboundPlayExplosion{}),
	0x1C: TypeOf(ClientboundPlayUnloadChunk{}),
	0x1D: TypeOf(ClientboundPlayGameStateChange{}),
	0x1E: TypeOf(ClientboundPlayOpenHorseWindow{}),
	0x1F: TypeOf(ClientboundPlayKeepAlive{}),
	0x20: TypeOf(ClientboundPlayMapChunk{}),
	0x21: TypeOf(ClientboundPlayWorldEvent{}),
	0x22: TypeOf(ClientboundPlayWorldParticles{}),
	0x23: TypeOf(ClientboundPlayUpdateLight{}),
	0x24: TypeOf(ClientboundJoinGamePacket{}),
	0x25: TypeOf(ClientboundPlayMap{}),
	0x26: TypeOf(ClientboundPlayTradeList{}),
	0x27: TypeOf(ClientboundPlayRelEntityMove{}),
	0x28: TypeOf(ClientboundPlayEntityMoveLook{}),
	0x29: TypeOf(ClientboundPlayEntityLook{}),
	0x2A: TypeOf(ClientboundPlayEntity{}),
	0x2B: TypeOf(ClientboundPlayVehicleMove{}),
	0x2C: TypeOf(ClientboundPlayOpenBook{}),
	0x2D: TypeOf(ClientboundPlayOpenWindow{}),
	0x2E: TypeOf(ClientboundPlayOpenSignEntity{}),
	0x2F: TypeOf(ClientboundPlayCraftRecipeResponse{}),
	0x30: TypeOf(ClientboundPlayerAbilitiesPacket{}),
	0x31: TypeOf(ClientboundPlayCombatEvent{}),
	0x32: TypeOf(ClientboundPlayerInfoPacket{}),
	0x33: TypeOf(ClientboundPlayFacePlayer{}),
	0x34: TypeOf(ClientboundPositionLookPacket{}),
	0x35: TypeOf(ClientboundPlayUnlockRecipes{}),
	0x36: TypeOf(ClientboundPlayEntityDestroy{}),
	0x37: TypeOf(ClientboundPlayRemoveEntityEffect{}),
	0x38: TypeOf(ClientboundPlayResourcePackSend{}),
	0x39: TypeOf(ClientboundPlayRespawn{}),
	0x3A: TypeOf(ClientboundPlayEntityHeadRotation{}),
	0x3B: TypeOf(ClientboundPlayMultiBlockChange{}),
	0x3C: TypeOf(ClientboundPlaySelectAdvancementTab{}),
	0x3D: TypeOf(ClientboundPlayWorldBorder{}),
	0x3E: TypeOf(ClientboundPlayCamera{}),
	0x3F: TypeOf(ClientboundHeldItemChangePacket{}),
	0x40: TypeOf(ClientboundPlayUpdateViewPosition{}),
	0x41: TypeOf(ClientboundPlayUpdateViewDistance{}),
	0x42: TypeOf(ClientboundPlaySpawnPosition{}),
	0x43: TypeOf(ClientboundPlayScoreboardDisplayObjective{}),
	0x44: TypeOf(ClientboundPlayEntityMetadata{}),
	0x45: TypeOf(ClientboundPlayAttachEntity{}),
	0x46: TypeOf(ClientboundPlayEntityVelocity{}),
	0x47: TypeOf(ClientboundPlayEntityEquipment{}),
	0x48: TypeOf(ClientboundPlayExperience{}),
	0x49: TypeOf(ClientboundPlayUpdateHealth{}),
	0x4A: TypeOf(ClientboundPlayScoreboardObjective{}),
	0x4B: TypeOf(ClientboundPlaySetPassengers{}),
	0x4C: TypeOf(ClientboundPlayTeams{}),
	0x4D: TypeOf(ClientboundPlayScoreboardScore{}),
	0x4E: TypeOf(ClientboundPlayUpdateTime{}),
	0x4F: TypeOf(ClientboundPlayTitle{}),
	0x50: TypeOf(ClientboundPlayEntitySoundEffect{}),
	0x51: TypeOf(ClientboundPlaySoundEffect{}),
	0x52: TypeOf(ClientboundPlayStopSound{}),
	0x53: TypeOf(ClientboundPlayPlayerlistHeader{}),
	0x54: TypeOf(ClientboundPlayNbtQueryResponse{}),
	0x55: TypeOf(ClientboundPlayCollect{}),
	0x56: TypeOf(ClientboundPlayEntityTeleport{}),
	0x57: TypeOf(ClientboundPlayAdvancements{}),
	0x58: TypeOf(ClientboundPlayEntityUpdateAttributes{}),
	0x59: TypeOf(ClientboundPlayEntityEffect{}),
	0x5A: TypeOf(ClientboundDeclareRecipesPacket{}),
	0x5B: TypeOf(ClientboundPlayTags{}),
}

var PlayServerboundIds = map[int32]Type{
	0x00: TypeOf(ServerboundPlayTeleportConfirm{}),
	0x01: TypeOf(ServerboundPlayQueryBlockNbt{}),
	0x02: TypeOf(ServerboundPlaySetDifficulty{}),
	0x03: TypeOf(ServerboundPlayChat{}),
	0x04: TypeOf(ServerboundPlayClientCommand{}),
	0x05: TypeOf(ServerboundClientSettingsPacket{}),
	0x06: TypeOf(ServerboundPlayTabComplete{}),
	0x07: TypeOf(ServerboundPlayTransaction{}),
	0x08: TypeOf(ServerboundPlayEnchantItem{}),
	0x09: TypeOf(ServerboundPlayWindowClick{}),
	0x0A: TypeOf(ServerboundPlayCloseWindow{}),
	0x0B: TypeOf(ServerboundPluginMessagePacket{}),
	0x0C: TypeOf(ServerboundPlayEditBook{}),
	0x0D: TypeOf(ServerboundPlayQueryEntityNbt{}),
	0x0E: TypeOf(ServerboundInteractEntityPacket{}),
	0x0F: TypeOf(ServerboundPlayGenerateStructure{}),
	0x10: TypeOf(ServerboundPlayKeepAlive{}),
	0x11: TypeOf(ServerboundPlayLockDifficulty{}),
	0x12: TypeOf(ServerboundPlayPosition{}),
	0x13: TypeOf(ServerboundPlayPositionLook{}),
	0x14: TypeOf(ServerboundPlayLook{}),
	0x15: TypeOf(ServerboundPlayFlying{}),
	0x16: TypeOf(ServerboundPlayVehicleMove{}),
	0x17: TypeOf(ServerboundPlaySteerBoat{}),
	0x18: TypeOf(ServerboundPlayPickItem{}),
	0x19: TypeOf(ServerboundPlayCraftRecipeRequest{}),
	0x1A: TypeOf(ServerboundPlayAbilities{}),
	0x1B: TypeOf(ServerboundPlayBlockDig{}),
	0x1C: TypeOf(ServerboundPlayEntityAction{}),
	0x1D: TypeOf(ServerboundPlaySteerVehicle{}),
	0x1E: TypeOf(ServerboundPlayDisplayedRecipe{}),
	0x1F: TypeOf(ServerboundPlayRecipeBook{}),
	0x20: TypeOf(ServerboundPlayNameItem{}),
	0x21: TypeOf(ServerboundPlayResourcePackReceive{}),
	0x22: TypeOf(ServerboundAdvancementTabPacket{}),
	0x23: TypeOf(ServerboundPlaySelectTrade{}),
	0x24: TypeOf(ServerboundPlaySetBeaconEffect{}),
	0x25: TypeOf(ServerboundPlayHeldItemSlot{}),
	0x26: TypeOf(ServerboundPlayUpdateCommandBlock{}),
	0x27: TypeOf(ServerboundPlayUpdateCommandBlockMinecart{}),
	0x28: TypeOf(ServerboundPlaySetCreativeSlot{}),
	0x29: TypeOf(ServerboundPlayUpdateJigsawBlock{}),
	0x2A: TypeOf(ServerboundPlayUpdateStructureBlock{}),
	0x2B: TypeOf(ServerboundPlayUpdateSign{}),
	0x2C: TypeOf(ServerboundPlayArmAnimation{}),
	0x2D: TypeOf(ServerboundPlaySpectate{}),
	0x2E: TypeOf(ServerboundPlayBlockPlace{}),
	0x2F: TypeOf(ServerboundPlayUseItem{}),
}
