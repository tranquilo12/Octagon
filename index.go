package main

import (
	"Octagon/structs"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	padding        = 2
	maxWidth       = 80
	indexState int = iota
	progressState
)

type ContinueMsg struct {
	url   string
	state int
}
type StopMsg struct{}
type Model struct {
	state    StateModel
	index    IdxModel
	progress PbModel
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
var tickerNamesResults []structs.TickersStruct

func (m *Model) SetState(state int) tea.Cmd {
	return func() tea.Msg {
		return ContinueMsg{state: state}
	}
}

func InitialModel() Model {
	return Model{
		StateModel{state: indexState},
		NewIndexModel(),
		NewProgressbarModel(progress.New(progress.WithDefaultGradient())),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Looking for the type of message that we received, it could be a StateModel, ContinueMsg, StopMsg, ReadyMsg
	// Looking for the type of state that we are in right now, that way we can maintain the keystrokes into IndexUpdate
	switch m.state.state {
	case indexState:
		m, cmd := m.IndexUpdate(msg)
		return m, cmd
	case progressState:
		m, cmd := m.progressbarUpdate(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m Model) View() string {
	s := ""
	switch m.state.state {
	case indexState:
		s += m.IndexView()
	case progressState:
		s += m.progressbarView()
	}
	// The footer
	s += helpStyle("\nPress q or 'ctr+c' to quit.")
	// Send the UI for rendering
	return s
}

const basePath = "https://api.polygon.io"

var baseUrl = fmt.Sprintf("%s/v3/reference/tickers?active=true&sort=ticker&order=asc&limit=1000", basePath)

//GetTickerNamesFromUrlAndClean is a function that gets data from an url, cleans it and returns it as a TickersStruct
func GetTickerNamesFromUrlAndClean(url string, apiKey string) structs.TickersStruct {
	tickersStruct := structs.TickersStruct{}
	url = fmt.Sprintf("%s&apiKey=%s", url, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &tickersStruct)
	if err != nil {
		panic(err)
	}
	return tickersStruct
}

// GetAllTickerNames returns all the ticker names
func (m *Model) GetAllTickerNames(url string, apiKey string, results *[]structs.TickersStruct) tea.Cmd {
	return func() tea.Msg {
		if url == "" {
			url = baseUrl
		}

		tickers := GetTickerNamesFromUrlAndClean(url, apiKey)
		*results = append(*results, tickers)

		if tickers.NextURL != "" {
			return ContinueMsg{url: tickers.NextURL}
		} else {
			return StopMsg{}
		}
	}
}

// WriteTickerNamesToJsonFile writes the ticker names to a file
func WriteTickerNamesToJsonFile(tickers []structs.TickersStruct) {
	jsonData, err := json.Marshal(tickers)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("data/tickerNames.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}
