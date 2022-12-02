package ValorantAPI

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MatchUserPlaytime struct {
	Minutes      int `json:"minutes"`
	Seconds      int `json:"seconds"`
	Milliseconds int `json:"milliseconds"`
}

type MatchUserBehaviorFriendlyFire struct {
	Incoming int `json:"incoming"`
	Outgoing int `json:"outgoing"`
}

type MatchUserPlatformOS struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type MatchUserPlatform struct {
	PC string               `json:"type"`
	OS *MatchUserPlatformOS `json:"os"`
}

type MatchUserAbilityCasts struct {
	C int `json:"c_cast"`
	Q int `json:"q_cast"`
	E int `json:"e_cast"`
	X int `json:"x_cast"`
}

type MatchUserBehavior struct {
	RoundsAFK     float32                        `json:"afk_rounds"`
	FriendlyFire  *MatchUserBehaviorFriendlyFire `json:"friendly_fire"`
	RoundsInSpawn float32                        `json:"rounds_in_spawn"`
}

type MatchUserAssetsCard struct {
	Small string `json:"small"`
	Large string `json:"large"`
	Wide  string `json:"wide"`
}

type MatchUserAssetsAgent struct {
	Small    string `json:"small"`
	Bust     string `json:"bust"`
	Full     string `json:"full"`
	Killfeed string `json:"killfeed"`
}

type MatchUserAssets struct {
	Card  *MatchUserAssetsCard  `json:"card"`
	Agent *MatchUserAssetsAgent `json:"agent"`
}

type MatchUserEconomySpent struct {
	Overall int `json:"overall"`
	Average int `json:"average"`
}
type MatchUserEconomyLoadoutValue struct {
	Overall int `json:"overall"`
	Average int `json:"average"`
}

type MatchUserEconomy struct {
	Spent        *MatchUserEconomySpent        `json:"spent"`
	LoadoutValue *MatchUserEconomyLoadoutValue `json:"loadout_value"`
}

type MatchUser struct {
	PUUID              string                 `json:"puuid"`
	Name               string                 `json:"name"`
	Tag                string                 `json:"tag"`
	Level              int                    `json:"level"`
	Character          string                 `json:"character"`
	CurrentTier        int                    `json:"currenttier"`
	CurrentTierPatched string                 `json:"currenttier_patched"`
	PlayerCard         string                 `json:"player_card"`
	PlayerTitle        string                 `json:"player_title"`
	PartyID            string                 `json:"party_id"`
	SessionPlaytime    *MatchUserPlaytime     `json:"session_playtime"`
	Behavior           *MatchUserBehavior     `json:"behavior"`
	Platform           *MatchUserPlatform     `json:"platform"`
	AbilityCasts       *MatchUserAbilityCasts `json:"ability_casts"`
	Assets             *MatchUserAssets       `json:"assets"`
	Economy            *MatchUserEconomy      `json:"economy"`
	DamageMade         int                    `json:"damage_made"`
	DamageRecieved     int                    `json:"damage_received"`
}

type MatchMetadata struct {
	Map              string `json:"map"`
	GameVersion      string `json:"game_version"`
	GameLength       int    `json:"game_length"`
	GameStart        int64  `json:"game_start"`
	GameStartPatched string `json:"game_start_patched"`
	RoundsPlayed     int    `json:"rounds_played"`
	Mode             string `json:"mode"`
	Queue            string `json:"queue"`
	SeasonID         string `json:"season_id"`
	Platform         string `json:"platform"`
	MatchID          string `json:"matchid"`
	Region           string `json:"region"`
	Cluster          string `json:"cluster"`
}

type MatchPlayers struct {
	AllPlayers []*MatchUser `json:"all_players"`
	Red        []*MatchUser `json:"red"`
	Blue       []*MatchUser `json:"blue"`
}

type MatchTeamsData struct {
	HasWon     bool `json:"has_won"`
	RoundsWon  int  `json:"rounds_won"`
	RoundsLost int  `json:"rounds_lost"`
}

type MatchTeams struct {
	Red  *MatchTeamsData `json:"red"`
	Blue *MatchTeamsData `json:"blue"`
}

type MatchRoundsPlantAndDefuseEventPlayer struct {
	PUUID       string `json:"puuid"`
	DisplayName string `json:"display_name"`
	PlantSite   string `json:"plant_site"`
}

type MatchRoundsPlantEventPlayersLocationsOnPlantAndDefuse struct {
	PUUID       string       `json:"player_puuid"`
	DisplayName string       `json:"player_display_name"`
	PlayerTeam  string       `json:"player_team"`
	Location    *Coordinates `json:"location"`
	ViewRadians float32      `json:"view_radians"`
}

type MatchRoundsPlantEvent struct {
	PlantLocation           *Coordinates                                             `json:"plant_location"`
	PlantedBy               *MatchRoundsPlantAndDefuseEventPlayer                    `json:"planted_by"`
	PlantSite               string                                                   `json:"plant_site"`
	PlantTimeInRound        string                                                   `json:"plant_time_in_round"`
	PlayersLocationsOnPlant []*MatchRoundsPlantEventPlayersLocationsOnPlantAndDefuse `json:"player_locations_on_plant"`
}

type MatchRoundsDefuseEvent struct {
	Location                *Coordinates                                             `json:"defuse_location"`
	DefusedBy               []*MatchRoundsPlantAndDefuseEventPlayer                  `json:"defused_by"`
	DefuseTimeIntoRound     int                                                      `json:"defuse_time_in_round"`
	PlayerLocationsOnDefuse []*MatchRoundsPlantEventPlayersLocationsOnPlantAndDefuse `json:"player_locations_on_defuse"`
}

type MatchStatsDamageEvents struct {
	RecieverPUUID       string `json:"receiver_puuid"`
	RecieverDisplayName string `json:"receiver_display_name"`
	RecieverTeam        string `json:"receiver_team"`
	Bodyshots           int    `json:"bodyshots"`
	Damage              int    `json:"damage"`
	Headshots           int    `json:"headshots"`
	Legshots            int    `json:"legshots"`
}

type MatchStatsKillEvents struct {
	TimeIntoRound     int    `json:"kill_time_in_round"`
	TimeIntoMatch     int    `json:"kill_time_in_match"`
	KillerPUUID       string `json:"killer_puuid"`
	KillerDisplayName string `json:"killer_display_name"`
	KillerTeam        string `json:"killer_team"`
	VictimPUUID       string `json:"victim_puuid"`
	VitcimDisplayName string `json:"victim_display_name"`
	// KillerTeam        string `json:""`
	// KillerTeam        string `json:""`
}

type MatchRoundsAbilityCasts struct {
	AbilityCasts *MatchUserAbilityCasts    `json:"ability_casts"`
	PPUID        string                    `json:"player_puuid"`
	DisplayName  string                    `json:"player_display_name"`
	Team         string                    `json:"player_team"`
	DamageEvents []*MatchStatsDamageEvents `json:"damage_events"`
	KillEvents   []*MatchStatsKillEvents   `json:"kill_events"`
}

type MatchRounds struct {
	WinningTeam  string                  `json:"winning_team"`
	EndType      string                  `json:"end_type"`
	BombPlanted  bool                    `json:"bomb_planted"`
	BombDefused  bool                    `json:"bomb_defused"`
	PlantEvents  *MatchRoundsPlantEvent  `json:"plant_events"`
	DefuseEvents *MatchRoundsDefuseEvent `json:"defuse_events"`
	PlayerStats  string                  `json:"player_stats"`
}

type Match struct {
	Metadata *MatchMetadata `json:"metadata"`
	Players  *MatchPlayers  `json:"players"`
	Teams    *MatchTeams    `json:"teams"`
}

type API_Response struct {
	Status  int      `json:"status"`
	Matches []*Match `json:"data"`
}
