package match

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gregtaole/lolapiwrapper/util"
)

const (
	rootURL      = "match/v4/"
	matchURL     = "matches"
	matchListURL = "matchlists/by-account"
	timelineURL  = "timelines/by-match"
)

type Match struct {
	APIKey string
	Region string
}

type MatchDTO struct {
	SeasonID              int                      `json:"seasonId"`
	QueueID               int                      `json:"queueId"`
	GameID                int                      `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"participantIdentities"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
	GameMode              string                   `json:"gameMode"`
	MapID                 int                      `json:"mapID"`
	GameType              string                   `json:"gameType"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	Participants          []ParticipantDTO         `json:"participants"`
	GameDuration          int                      `json:"gameDuration"`
	GameCreation          int                      `json:"gameCreation"`
}

type ParticipantIdentityDTO struct {
	Player        PlayerDTO `json:"player"`
	ParticipantID int       `json:"participantId"`
}

type PlayerDTO struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  string `json:"currentAccountId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        string `json:"summonerId"`
	AccountID         string `json:"accountId"`
}

type TeamStatsDTO struct {
	FirstDragon          bool          `json:"firstDragon"`
	FirstInhibitor       bool          `json:"firstInhibitor"`
	Bans                 []TeamBansDTO `json:"bans"`
	BaronKills           int           `json:"baronKills"`
	FirstRiftHerald      bool          `json:"firstRiftHerald"`
	FirstBaron           bool          `json:"firstBaron"`
	RiftHeraldKills      int           `json:"riftHeraldKills"`
	FirstBlood           bool          `json:"firstBlood"`
	TeamID               int           `json:"teamId"`
	FirstTower           bool          `json:"firstTower"`
	VilemawKills         int           `json:"vilemawKills"`
	InhibitorKills       int           `json:"inhibitorKills"`
	TowerKills           int           `json:"towerKills"`
	DominionVictoryScore int           `json:"dominionVictoryScore"`
	Win                  string        `json:"win"`
	DragonKills          int           `json:"dragonKills"`
}

type TeamBansDTO struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
}

type ParticipantDTO struct {
	Stats                     ParticipantStatsDTO    `json:"stats"`
	ParticipantID             int                    `json:"participantID"`
	Runes                     []RuneDTO              `json:"runes"`
	Timeline                  ParticipantTimelineDTO `json:"timeline"`
	TeamID                    int                    `json:"teamId"`
	Spell2ID                  int                    `json:"spell2Id"`
	Masteries                 []MasteryDTO           `json:"masteries"`
	HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
	Spell1ID                  int                    `json:"spell1Id"`
	ChampionID                int                    `json:"championId"`
}

type ParticipantStatsDTO struct {
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	VisionScore                     int  `json:"visionScore"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
	Perk1Var1                       int  `json:"perk1Var1"`
	Perk1Var3                       int  `json:"perk1Var3"`
	Perk1Var2                       int  `json:"perk1Var2"`
	TripleKills                     int  `json:"tripleKills"`
	Perk3Var3                       int  `json:"perk3Var3"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	Perk3Var2                       int  `json:"perk3Var2"`
	PlayerScore9                    int  `json:"playerScore9"`
	PlayerSocre8                    int  `json:"playerScore8"`
	Kills                           int  `json:"kills"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore6                    int  `json:"playerScore6"`
	Perk5Var1                       int  `json:"perk5Var1"`
	Perk5Var3                       int  `json:"perk5Var3"`
	Perk5Var2                       int  `json:"perk5Var2"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	NodeCapture                     int  `json:"nodeCapture"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	Perk2Var2                       int  `json:"perk2Var2"`
	Perk2Var3                       int  `json:"perk2Var3"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	Perk2Var1                       int  `json:"perk2Var1"`
	Perk4Var1                       int  `json:"perk4Var1"`
	Perk4Var2                       int  `json:"perk4Var2"`
	Perk4Var3                       int  `json:"perk4Var3"`
	WardsKilled                     int  `json:"wardsKilled"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	QuadraKills                     int  `json:"quadraKills"`
	TeamObjective                   int  `json:"teamObjective"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item0                           int  `json:"item0"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	Item6                           int  `json:"item6"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Perk1                           int  `json:"perk1"`
	Perk0                           int  `json:"perk0"`
	Perk3                           int  `json:"perk3"`
	Perk2                           int  `json:"perk2"`
	Perk5                           int  `json:"perk5"`
	Perk4                           int  `json:"perk4"`
	Perk3Var1                       int  `json:"perk3Var1"`
	DamageSeflMitigated             int  `json:"damageSelfMitigated"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	Assists                         int  `json:"assists"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
	GoldSpent                       int  `json:"goldSpent"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	ParticipantID                   int  `json:"participantId"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	Win                             bool `json:"win"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	Item1                           int  `json:"item1"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	Deaths                          int  `json:"deaths"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	PerkSubStyle                    int  `json:"perkSubStyle"`
	TurretKills                     int  `json:"turretKills"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	TrueDamageDealtToChampions      int  `json:"TrueDamageDealtToChampions"`
	GoldEarned                      int  `json:"goldEarned"`
	KillingSprees                   int  `json:"killingSprees"`
	UnrealKills                     int  `json:"unrealKills"`
	AltarsCaptured                  int  `json:"altarsCaptured"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	ChampLevel                      int  `json:"champLevel"`
	DoubleKills                     int  `json:"doubleKills"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	Perk0Var1                       int  `json:"perk0Var1"`
	Perk0Var2                       int  `json:"perk0Var2"`
	Perk0Var3                       int  `json:"perk0Var3"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	AltarsNeutralized               int  `json:"altarsNeutralized"`
	PentaKills                      int  `json:"pentaKills"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
}

type RuneDTO struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

type ParticipantTimelineDTO struct {
	Lane                        string             `json:"lane"`
	ParticipantID               int                `json:"participantId"`
	CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
	XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
}

type MasteryDTO struct {
	MasteryID int `json:"masteryId"`
	Rank      int `json:"rank"`
}

type MatchListDTO struct {
	Matches    []MatchReferenceDTO `json:"matches"`
	TotalGames int                 `json:"totalGames"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
}

type MatchReferenceDTO struct {
	Lane       string `json:"lane"`
	GameID     int    `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Timestamp  int    `json:"timestamp"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Season     int    `json:"season"`
}

type MatchTimelineDTO struct {
	Frames        []MatchFrameDTO `json:"frames"`
	FrameInterval int             `json:"FrameInterval"`
}

type MatchFrameDTO struct {
	Timestamp         int                                 `json:"timestamp"`
	ParticipantFrames map[string]MatchParticipantFrameDTO `json:"participantFrames"`
	Events            []MatchEventDTO                     `json:"events"`
}

type MatchParticipantFrameDTO struct {
	TotalGold           int              `json:"totalGold"`
	TeamScore           int              `json:"teamScore"`
	ParticipantID       int              `json:"participantId"`
	Level               int              `json:"level"`
	CurrentGold         int              `json:"currentGold"`
	MinionsKilled       int              `json:"minionsKilled"`
	DominionScore       int              `json:"dominionScore"`
	Position            MatchPositionDTO `json:"position"`
	Xp                  int              `json:"xp"`
	JungleMinionsKilled int              `json:"jungleMinionsKilled"`
}

type MatchPositionDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MatchEventDTO struct {
	EventType               string           `json:"eventType"`
	TowerType               string           `json:"towerType"`
	TeamID                  int              `json:"teamId"`
	AscendedType            string           `json:"ascendedType"`
	KillerID                int              `json:"killerId"`
	LevelUpType             string           `json:"levelUpType"`
	PointCaptured           string           `json:"pointCaptured"`
	AssistingParticipantIDs []int            `json:"assistingParticipantIds"`
	WardType                string           `json:"wardType"`
	MonsterType             string           `json:"monsterType"`
	Type                    string           `json:"type"`
	SkillSlot               int              `json:"skillSlot"`
	VictimID                int              `json:"victimId"`
	Timestamp               int              `json:"timestamp"`
	AfterID                 int              `json:"afterId"`
	MonsterSubType          string           `json:"monsterSubType"`
	LaneType                string           `json:"laneType"`
	ItemID                  int              `json:"itemId"`
	ParticipantID           int              `json:"participantId"`
	BuildingType            string           `json:"buildingType"`
	CreatorID               int              `json:"creatorId"`
	Position                MatchPositionDTO `json:"position"`
	BeforeID                int              `json:"beforeId"`
}

type MatchQueryParams struct {
	Champion   []int
	Queue      []int
	Season     []int
	EndTime    int
	BeginTime  int
	EndIndex   int
	BeginIndex int
}

func NewMatch(APIKey string, region string) Match {
	return Match{
		APIKey: APIKey,
		Region: region,
	}
}

func (m Match) MatchesByID(matchID string) (*MatchDTO, error) {
	resp, err := util.GetResponse(m.APIKey, m.Region, rootURL+matchURL, matchID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var matchDTO MatchDTO
	if err = json.Unmarshal(body, &matchDTO); err != nil {
		return nil, err
	}
	return &matchDTO, nil

}

func (m Match) MatchListByAccount(accountID string, params MatchQueryParams) (*MatchListDTO, error) {
	if params.BeginIndex >= 0 && params.EndIndex >= 0 && params.BeginIndex < params.EndIndex {
		return nil, fmt.Errorf("MatchQueryParams.BeginIndex should be greater than MatchQueryParams.EndIndex when both are passed: %v < %v", params.EndIndex, params.BeginIndex)
	}
	if params.BeginTime >= 0 && params.EndTime >= 0 && params.BeginTime < params.EndTime {
		return nil, fmt.Errorf("MatchQueryParams.BeginTime should be greater than MatchQueryParams.EndTime when both are passed: %v < %v", params.EndTime, params.BeginTime)
	}
	queryParams := params.String()
	client := &http.Client{}
	url := fmt.Sprintf("https://%v.api.riotgames.com/lol/%v/%v%v", m.Region, rootURL+matchListURL, accountID, queryParams)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for %v: %v", url, err)
	}
	fmt.Println(req.URL)
	req.Header.Set("X-Riot-Token", m.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get %v: %v", req.URL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var matchListDTO MatchListDTO
	if err = json.Unmarshal(body, &matchListDTO); err != nil {
		return nil, err
	}
	return &matchListDTO, nil
}

func (m Match) TimelineByMatch(matchID string) (*MatchTimelineDTO, error) {
	resp, err := util.GetResponse(m.APIKey, m.Region, rootURL+timelineURL, matchID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var matchTimelineDTO MatchTimelineDTO
	if err = json.Unmarshal(body, &matchTimelineDTO); err != nil {
		return nil, err
	}
	return &matchTimelineDTO, nil
}

func (m MatchDTO) String() string {
	return fmt.Sprintf("%v\n", m.Participants)
}

func (p ParticipantDTO) String() string {
	return fmt.Sprintf("%v\n", p.ChampionID)
}

func (m MatchQueryParams) String() string {
	var queryParams strings.Builder
	if len(m.Champion) == 0 && len(m.Queue) == 0 && len(m.Season) == 0 && m.BeginTime < 0 && m.EndTime < 0 && m.BeginIndex < 0 && m.EndIndex < 0 {
		return ""
	}
	queryParams.WriteString("?")
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
	if m.BeginTime >= 0 {
		queryParams.WriteString(fmt.Sprintf("beginTime=%v&", m.BeginTime))
	}
	if m.EndTime >= 0 {
		queryParams.WriteString(fmt.Sprintf("endTime=%v&", m.EndTime))
	}
	if m.BeginIndex >= 0 {
		queryParams.WriteString(fmt.Sprintf("beginIndex=%v&", m.BeginIndex))
	}
	if m.EndIndex >= 0 {
		queryParams.WriteString(fmt.Sprintf("endIndex=%v&", m.EndIndex))
	}
	s := queryParams.String()
	if strings.HasSuffix(s, "&") {
		s = s[:len(s)-1]
	}
	return s
}
