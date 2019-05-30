package lolapiwrapper

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

const (
	matchRootURL = "match/v4/"
	matchURL     = "matches"
	matchListURL = "matchlists/by-account"
	timelineURL  = "timelines/by-match"
)

/*MatchDTO …*/
type MatchDTO struct {
	GameID                int                      `json:"gameId"`
	SeasonID              int                      `json:"seasonId"`
	QueueID               int                      `json:"queueId"`
	GameType              string                   `json:"gameType"`
	GameMode              string                   `json:"gameMode"`
	MapID                 int                      `json:"mapID"`
	Participants          []ParticipantDTO         `json:"participants"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"participantIdentities"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	GameDuration          int                      `json:"gameDuration"`
	GameCreation          int                      `json:"gameCreation"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
}

/*ParticipantIdentityDTO …*/
type ParticipantIdentityDTO struct {
	ParticipantID int       `json:"participantId"`
	Player        PlayerDTO `json:"player"`
}

/*PlayerDTO …*/
type PlayerDTO struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	SummonerID        string `json:"summonerId"`
	AccountID         string `json:"accountId"`
	CurrentAccountID  string `json:"currentAccountId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	ProfileIcon       int    `json:"profileIcon"`
	PlatformID        string `json:"platformId"`
}

/*TeamStatsDTO …*/
type TeamStatsDTO struct {
	TeamID               int           `json:"teamId"`
	Win                  string        `json:"win"`
	Bans                 []TeamBansDTO `json:"bans"`
	FirstBlood           bool          `json:"firstBlood"`
	FirstTower           bool          `json:"firstTower"`
	TowerKills           int           `json:"towerKills"`
	FirstInhibitor       bool          `json:"firstInhibitor"`
	InhibitorKills       int           `json:"inhibitorKills"`
	FirstDragon          bool          `json:"firstDragon"`
	DragonKills          int           `json:"dragonKills"`
	FirstRiftHerald      bool          `json:"firstRiftHerald"`
	RiftHeraldKills      int           `json:"riftHeraldKills"`
	FirstBaron           bool          `json:"firstBaron"`
	BaronKills           int           `json:"baronKills"`
	VilemawKills         int           `json:"vilemawKills"`
	DominionVictoryScore int           `json:"dominionVictoryScore"`
}

/*TeamBansDTO …*/
type TeamBansDTO struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
}

/*ParticipantDTO …*/
type ParticipantDTO struct {
	ParticipantID             int                    `json:"participantID"`
	TeamID                    int                    `json:"teamId"`
	ChampionID                int                    `json:"championId"`
	Runes                     []RuneDTO              `json:"runes"`
	Spell1ID                  int                    `json:"spell1Id"`
	Spell2ID                  int                    `json:"spell2Id"`
	Masteries                 []MasteryDTO           `json:"masteries"`
	Stats                     ParticipantStatsDTO    `json:"stats"`
	Timeline                  ParticipantTimelineDTO `json:"timeline"`
	HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
}

/*ParticipantStatsDTO …*/
type ParticipantStatsDTO struct {
	ParticipantID                   int  `json:"participantId"`
	ChampLevel                      int  `json:"champLevel"`
	Win                             bool `json:"win"`
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
	PerkSubStyle                    int  `json:"perkSubStyle"`
	Perk0                           int  `json:"perk0"`
	Perk0Var1                       int  `json:"perk0Var1"`
	Perk0Var2                       int  `json:"perk0Var2"`
	Perk0Var3                       int  `json:"perk0Var3"`
	Perk1                           int  `json:"perk1"`
	Perk1Var1                       int  `json:"perk1Var1"`
	Perk1Var2                       int  `json:"perk1Var2"`
	Perk1Var3                       int  `json:"perk1Var3"`
	Perk2                           int  `json:"perk2"`
	Perk2Var1                       int  `json:"perk2Var1"`
	Perk2Var2                       int  `json:"perk2Var2"`
	Perk2Var3                       int  `json:"perk2Var3"`
	Perk3                           int  `json:"perk3"`
	Perk3Var1                       int  `json:"perk3Var1"`
	Perk3Var2                       int  `json:"perk3Var2"`
	Perk3Var3                       int  `json:"perk3Var3"`
	Perk4                           int  `json:"perk4"`
	Perk4Var1                       int  `json:"perk4Var1"`
	Perk4Var2                       int  `json:"perk4Var2"`
	Perk4Var3                       int  `json:"perk4Var3"`
	Perk5                           int  `json:"perk5"`
	Perk5Var1                       int  `json:"perk5Var1"`
	Perk5Var2                       int  `json:"perk5Var2"`
	Perk5Var3                       int  `json:"perk5Var3"`
	Kills                           int  `json:"kills"`
	Deaths                          int  `json:"deaths"`
	Assists                         int  `json:"assists"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore6                    int  `json:"playerScore6"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore8                    int  `json:"playerScore8"`
	PlayerScore9                    int  `json:"playerScore9"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	VisionScore                     int  `json:"visionScore"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int  `json:"TrueDamageDealtToChampions"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	DamageSeflMitigated             int  `json:"damageSelfMitigated"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TotalHeal                       int  `json:"totalHeal"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	DoubleKills                     int  `json:"doubleKills"`
	TripleKills                     int  `json:"tripleKills"`
	QuadraKills                     int  `json:"quadraKills"`
	PentaKills                      int  `json:"pentaKills"`
	KillingSprees                   int  `json:"killingSprees"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	WardsKilled                     int  `json:"wardsKilled"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	TurretKills                     int  `json:"turretKills"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	NodeCapture                     int  `json:"nodeCapture"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	AltarsCaptured                  int  `json:"altarsCaptured"`
	AltarsNeutralized               int  `json:"altarsNeutralized"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	TeamObjective                   int  `json:"teamObjective"`
	UnrealKills                     int  `json:"unrealKills"`
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
}

/*RuneDTO …*/
type RuneDTO struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

/*ParticipantTimelineDTO …*/
type ParticipantTimelineDTO struct {
	ParticipantID               int                `json:"participantId"`
	Lane                        string             `json:"lane"`
	Role                        string             `json:"role"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
	XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
}

/*MasteryDTO …*/
type MasteryDTO struct {
	MasteryID int `json:"masteryId"`
	Rank      int `json:"rank"`
}

/*MatchListDTO …*/
type MatchListDTO struct {
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
	TotalGames int                 `json:"totalGames"`
	Matches    []MatchReferenceDTO `json:"matches"`
}

/*MatchReferenceDTO …*/
type MatchReferenceDTO struct {
	GameID     int    `json:"gameId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Champion   int    `json:"champion"`
	Role       string `json:"role"`
	Lane       string `json:"lane"`
	Timestamp  int    `json:"timestamp"`
	PlatformID string `json:"platformId"`
}

/*MatchTimelineDTO …*/
type MatchTimelineDTO struct {
	FrameInterval int             `json:"FrameInterval"`
	Frames        []MatchFrameDTO `json:"frames"`
}

/*MatchFrameDTO …*/
type MatchFrameDTO struct {
	Timestamp         int                                 `json:"timestamp"`
	ParticipantFrames map[string]MatchParticipantFrameDTO `json:"participantFrames"`
	Events            []MatchEventDTO                     `json:"events"`
}

/*MatchParticipantFrameDTO …*/
type MatchParticipantFrameDTO struct {
	ParticipantID       int              `json:"participantId"`
	Position            MatchPositionDTO `json:"position"`
	MinionsKilled       int              `json:"minionsKilled"`
	JungleMinionsKilled int              `json:"jungleMinionsKilled"`
	Xp                  int              `json:"xp"`
	Level               int              `json:"level"`
	CurrentGold         int              `json:"currentGold"`
	TotalGold           int              `json:"totalGold"`
	TeamScore           int              `json:"teamScore"`
	DominionScore       int              `json:"dominionScore"`
}

/*MatchPositionDTO …*/
type MatchPositionDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

/*MatchEventDTO …*/
type MatchEventDTO struct {
	TeamID                  int              `json:"teamId"`
	Timestamp               int              `json:"timestamp"`
	Type                    string           `json:"type"`
	Position                MatchPositionDTO `json:"position"`
	BeforeID                int              `json:"beforeId"`
	AfterID                 int              `json:"afterId"`
	CreatorID               int              `json:"creatorId"`
	EventType               string           `json:"eventType"`
	TowerType               string           `json:"towerType"`
	AscendedType            string           `json:"ascendedType"`
	LevelUpType             string           `json:"levelUpType"`
	WardType                string           `json:"wardType"`
	MonsterType             string           `json:"monsterType"`
	MonsterSubType          string           `json:"monsterSubType"`
	BuildingType            string           `json:"buildingType"`
	LaneType                string           `json:"laneType"`
	KillerID                int              `json:"killerId"`
	AssistingParticipantIDs []int            `json:"assistingParticipantIds"`
	VictimID                int              `json:"victimId"`
	PointCaptured           string           `json:"pointCaptured"`
	SkillSlot               int              `json:"skillSlot"`
	ItemID                  int              `json:"itemId"`
	ParticipantID           int              `json:"participantId"`
}

// MatchesByID gets the match information for the given matchID
func (c *client) MatchesByID(ctx context.Context, matchID string) (MatchDTO, error) {
	var res MatchDTO
	url := filepath.Join(matchRootURL, matchURL, matchID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// MatchListByAccount gets the match list for the given accountID filtered by params
func (c *client) MatchListByAccount(ctx context.Context, accountID string, params MatchQueryParams) (MatchListDTO, error) {
	var res MatchListDTO
	if *params.BeginIndex >= 0 && *params.EndIndex >= 0 && *params.BeginIndex < *params.EndIndex {
		return res, fmt.Errorf("MatchQueryParams.BeginIndex should be greater than MatchQueryParams.EndIndex when both are passed: %v < %v", params.EndIndex, params.BeginIndex)
	}
	if *params.BeginTime >= 0 && *params.EndTime >= 0 && *params.BeginTime < *params.EndTime {
		return res, fmt.Errorf("MatchQueryParams.BeginTime should be greater than MatchQueryParams.EndTime when both are passed: %v < %v", params.EndTime, params.BeginTime)
	}
	queryURL := filepath.Join(matchRootURL, matchListURL, accountID)
	vals, err := url.ParseQuery(params.String())
	if err != nil {
		return res, err
	}
	err = c.query(ctx, queryURL, vals, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

//TimelineByMatch gets the match timeline for the given matchID
func (c *client) TimelineByMatch(ctx context.Context, matchID string) (MatchTimelineDTO, error) {
	var res MatchTimelineDTO
	url := filepath.Join(matchRootURL, timelineURL, matchID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (m MatchQueryParams) String() string {
	var queryParams strings.Builder
	if len(m.Champion) == 0 && len(m.Queue) == 0 && len(m.Season) == 0 && *m.BeginTime < 0 && *m.EndTime < 0 && *m.BeginIndex < 0 && *m.EndIndex < 0 {
		return ""
	}
	//queryParams.WriteString("?")
	if len(m.Champion) != 0 {
		for _, c := range m.Champion {
			queryParams.WriteString(fmt.Sprintf("champion=%v&", c))
		}
	}
	if len(m.Queue) != 0 {
		for _, q := range m.Queue {
			queryParams.WriteString(fmt.Sprintf("queue=%v&", q))
		}
	}
	if len(m.Season) != 0 {
		for _, s := range m.Season {
			queryParams.WriteString(fmt.Sprintf("season=%v&", s))
		}
	}
	if *m.BeginTime >= 0 {
		queryParams.WriteString(fmt.Sprintf("beginTime=%v&", m.BeginTime))
	}
	if *m.EndTime >= 0 {
		queryParams.WriteString(fmt.Sprintf("endTime=%v&", m.EndTime))
	}
	if *m.BeginIndex >= 0 {
		queryParams.WriteString(fmt.Sprintf("beginIndex=%v&", m.BeginIndex))
	}
	if *m.EndIndex >= 0 {
		queryParams.WriteString(fmt.Sprintf("endIndex=%v&", m.EndIndex))
	}
	s := queryParams.String()
	if strings.HasSuffix(s, "&") {
		s = s[:len(s)-1]
	}
	return s
}

func (m MatchDTO) String() string {
	var str strings.Builder
	str.WriteString("MatchDTO{\n")
	str.WriteString(fmt.Sprintf("	GameID: %v,\n", m.GameID))
	str.WriteString(fmt.Sprintf("	SeasonID: %v,\n", m.SeasonID))
	str.WriteString(fmt.Sprintf("	QueueID: %v,\n", m.QueueID))
	str.WriteString(fmt.Sprintf("	GameType: \"%v\",\n", m.GameType))
	str.WriteString(fmt.Sprintf("	GameMode: \"%v\",\n", m.GameMode))
	str.WriteString(fmt.Sprintf("	MapID: %v,\n", m.MapID))
	str.WriteString("	Participants: []ParticipantDTO{\n")
	for _, participant := range m.Participants {
		str.WriteString("		ParticipantDTO{\n")
		str.WriteString(fmt.Sprintf("			ParticipantID: %v,\n", participant.ParticipantID))
		str.WriteString(fmt.Sprintf("			TeamID: %v,\n", participant.TeamID))
		str.WriteString(fmt.Sprintf("			ChampionID: %v,\n", participant.ChampionID))
		str.WriteString("			Runes: []RuneDTO{\n")
		for _, runes := range participant.Runes {
			str.WriteString("				RuneDTO{\n")
			str.WriteString(fmt.Sprintf("					RuneID: %v,\n", runes.RuneID))
			str.WriteString(fmt.Sprintf("					Rank: %v,\n", runes.Rank))
			str.WriteString("				},\n")
		}
		str.WriteString("			},\n")
		str.WriteString(fmt.Sprintf("			Spell1ID: %v,\n", participant.Spell1ID))
		str.WriteString(fmt.Sprintf("			Spell2ID: %v,\n", participant.Spell2ID))
		str.WriteString("			Masteries: []MasteryDTO{\n")
		for _, mastery := range participant.Masteries {
			str.WriteString("				MasteryDTO{\n")
			str.WriteString(fmt.Sprintf("					MasteryID: %v", mastery.MasteryID))
			str.WriteString(fmt.Sprintf("					Rank: %v", mastery.Rank))
			str.WriteString("				},\n")
		}
		str.WriteString("			},\n")

		str.WriteString(fmt.Sprintf("			HighestAchievedSeasonTier: \"%v\",\n", participant.HighestAchievedSeasonTier))
		str.WriteString("		},\n")
	}
	str.WriteString("	},\n")
	str.WriteString("	ParticipantIdentities: []ParticipantIdentityDTO{\n")
	for _, participantIdentity := range m.ParticipantIdentities {
		str.WriteString("		ParticipantIdentityDTO{\n")
		str.WriteString(fmt.Sprintf("			ParticipantID: %v,\n", participantIdentity.ParticipantID))
		str.WriteString("			Player: PlayerDTO{\n")
		str.WriteString(fmt.Sprintf("				CurrentPlatformID: \"%v\",\n", participantIdentity.Player.CurrentPlatformID))
		str.WriteString(fmt.Sprintf("				SummonerName: \"%v\",\n", participantIdentity.Player.SummonerName))
		str.WriteString(fmt.Sprintf("				SummonerID: \"%v\",\n", participantIdentity.Player.SummonerID))
		str.WriteString(fmt.Sprintf("				AccountID: \"%v\",\n", participantIdentity.Player.AccountID))
		str.WriteString(fmt.Sprintf("				CurrentAccountID: \"%v\",\n", participantIdentity.Player.CurrentAccountID))
		str.WriteString(fmt.Sprintf("				MatchHistoryURI: \"%v\",\n", participantIdentity.Player.MatchHistoryURI))
		str.WriteString(fmt.Sprintf("				ProfileIcon: %v,\n", participantIdentity.Player.ProfileIcon))
		str.WriteString(fmt.Sprintf("				PlatformID: \"%v\",\n", participantIdentity.Player.PlatformID))
		str.WriteString("			},\n")
		str.WriteString("		},\n")
	}
	str.WriteString("	},\n")
	str.WriteString("	Teams: []TeamStatsDTO{\n")
	for _, team := range m.Teams {
		str.WriteString("		TeamStatsDTO{\n")
		str.WriteString(fmt.Sprintf("			TeamID: %v,\n", team.TeamID))
		str.WriteString(fmt.Sprintf("			Win: \"%v\",\n", team.Win))
		str.WriteString("			Bans: []TeamBansDTO{\n")
		for _, teamBans := range team.Bans {
			str.WriteString("				TeamBansDTO{\n")
			str.WriteString(fmt.Sprintf("					PickTurn: %v,\n", teamBans.PickTurn))
			str.WriteString(fmt.Sprintf("					ChampionID: %v,\n", teamBans.ChampionID))
			str.WriteString("				},\n")
		}
		str.WriteString("			},\n")
		str.WriteString(fmt.Sprintf("			FirstBlood: %v,\n", team.FirstBlood))
		str.WriteString(fmt.Sprintf("			FirstTower: %v,\n", team.FirstTower))
		str.WriteString(fmt.Sprintf("			TowerKills: %v,\n", team.TowerKills))
		str.WriteString(fmt.Sprintf("			FirstInhibitor: %v,\n", team.FirstInhibitor))
		str.WriteString(fmt.Sprintf("			InhibitorKills: %v,\n", team.InhibitorKills))
		str.WriteString(fmt.Sprintf("			FirstDragon: %v,\n", team.FirstDragon))
		str.WriteString(fmt.Sprintf("			DragonKills: %v,\n", team.DragonKills))
		str.WriteString(fmt.Sprintf("			FirstRiftHerald: %v,\n", team.FirstRiftHerald))
		str.WriteString(fmt.Sprintf("			RiftHeraldKills: %v,\n", team.RiftHeraldKills))
		str.WriteString(fmt.Sprintf("			FirstBaron: %v,\n", team.FirstBaron))
		str.WriteString(fmt.Sprintf("			BaronKills: %v,\n", team.BaronKills))
		str.WriteString(fmt.Sprintf("			VilemawKills: %v,\n", team.VilemawKills))
		str.WriteString(fmt.Sprintf("			DominionVictoryScore: %v,\n", team.DominionVictoryScore))
		str.WriteString("		},\n")
	}
	str.WriteString("	},\n")
	str.WriteString(fmt.Sprintf("	GameDuration: %v,\n", m.GameDuration))
	str.WriteString(fmt.Sprintf("	GameCreation: %v,\n", m.GameCreation))
	str.WriteString(fmt.Sprintf("	GameVersion: \"%v\",\n", m.GameVersion))
	str.WriteString(fmt.Sprintf("	PlatformID: \"%v\",\n", m.PlatformID))
	str.WriteString("}\n")
	return str.String()
}

func (m MatchListDTO) String() string {
	var str strings.Builder
	str.WriteString("MatchListDTO{\n")
	str.WriteString(fmt.Sprintf("	TotalGames: %v,\n", m.TotalGames))
	str.WriteString("}\n")
	return str.String()
}

func (m MatchTimelineDTO) String() string {
	var str strings.Builder
	str.WriteString("MatchTimelineDTO{\n")
	str.WriteString(fmt.Sprintf("	FrameInterval: %v,\n", m.FrameInterval))
	str.WriteString("}\n")
	return str.String()
}
