package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/schemas"
	"time"
)

func GetSummonerInfoByName() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf(`https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s`, summonerName)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("status code: %d\n", response.StatusCode)
	fmt.Printf("summoner info: %s\n", string(responseBody))
}

func GetSummonerStatsBySummonerId() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf(`https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s`, summonerId)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	var summonerStats []schemas.SummonerStatsBySummonerId
	err = json.NewDecoder(response.Body).Decode(&summonerStats)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("status code: %d\n", response.StatusCode)
	for _, stat := range summonerStats {
		fmt.Printf("solo/flex summoner stat: %+v\n", stat)
	}
}

func GetSummonerMatchesListByPuuid() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf(`https://europe.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids`, puuid)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	query := request.URL.Query()
	query.Add("start", lastMatch)
	query.Add("count", lengthListOfMatches) // 100 max
	request.URL.RawQuery = query.Encode()

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("status code: %d\n", response.StatusCode)
	fmt.Printf("matches list: %s\n", string(responseBody))
}

func GetSummonerMatchInfoByMatchId() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf(`https://europe.api.riotgames.com/lol/match/v5/matches/%s`, matchId)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	var matchData schemas.MatchInfoByMatchId
	err = json.NewDecoder(response.Body).Decode(&matchData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("status code: %d\n", response.StatusCode)
	fmt.Printf("match metadata: %s\n", matchData.Metadata)
	fmt.Printf("match info: %+v\n", matchData.Info)
}

func GetSingleChampionInfoBySummonerIdAndChampionId() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf(`https://euw1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/by-champion/%s`, summonerId, championId)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	var championInfo schemas.ChampionInfoBySummonerId
	err = json.NewDecoder(response.Body).Decode(&championInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("status code: %d\n", response.StatusCode)
	fmt.Printf("champion info: %+v\n", championInfo)
}

func GetAllChampionsInfoBySummonerId() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf(`https://euw1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s`, summonerId)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	var matchData []schemas.ChampionInfoBySummonerId
	err = json.NewDecoder(response.Body).Decode(&matchData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("status code: %d\n", response.StatusCode)
	fmt.Printf("all champions info: %+v\n", matchData)
}

func GetTopChampionsInfoBySummonerIdAndCount() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf(`https://euw1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/top`, summonerId)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	query := request.URL.Query()
	query.Add("count", championsCount)
	request.URL.RawQuery = query.Encode()

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	var matchData []schemas.ChampionInfoBySummonerId
	err = json.NewDecoder(response.Body).Decode(&matchData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("status code: %d\n", response.StatusCode)
	fmt.Printf("top champions info: %+v\n", matchData)
}

func GetCurrentGameInformationBySummonerIdA() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf(`https://euw1.api.riotgames.com/lol/spectator/v4/active-games/by-summoner/%s`, summonerId)

	bbody := bytes.NewReader([]byte{})

	request, err := http.NewRequest("GET", url, bbody)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("X-Riot-Token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("status code: %d\n", response.StatusCode)
	fmt.Printf("current game info: %s\n", string(responseBody))
}

func main() {
	fmt.Println("project started")
	//GetSummonerInfoByName()
	//GetSummonerStatsBySummonerId()
	//GetSummonerMatchesListByPuuid()
	GetSummonerMatchInfoByMatchId()
	//GetSingleChampionInfoBySummonerIdAndChampionId()
	//GetAllChampionsInfoBySummonerId()
	//GetTopChampionsInfoBySummonerIdAndCount()
	//GetCurrentGameInformationBySummonerIdA()
}
