# JungleTree

Minecraft server software, written from the ground-up in Go.

NOTE: config to reproduce junglemc/JungleTree#9
```
DebugMode = true
Difficulty = "normal"
Gamemode = "survival"
MOTD = "JungleTree Debug"
MaxOnlinePlayers = 20
Verbose = true

[JavaEdition]
  OnlineMode = false

[Network]
  IP = ""
  NetworkCompressionThreshold = -1
  Port = 25565

[[BannedPlayers]]
Player = "hubofeverything"
Reason = "ab"
```