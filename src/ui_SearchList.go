package goass

import (
	"fmt"

	"github.com/Mericusta/go-assistant/src/common"
	"github.com/Mericusta/go-assistant/src/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type ui_SearchList struct {
	*ui.List

	debugData *data_Debug
	logicData *data_Logic

	title   string
	content string
}

func NewUISearch(debugData *data_Debug, logicData *data_Logic) *ui_SearchList {
	uiSearchList := &ui_SearchList{
		List:      ui.NewList(0),
		logicData: logicData,
		title:     "Search",
		content:   "this is search model",
	}
	uiSearchFunc := NewUISearchFunc(debugData, logicData)
	// uiGenerateBenchmark := NewUIGenerateBenchmark(debugData, logicData)
	// uiGenerateAST := NewUIGenerateAST(debugData, logicData)
	// uiGenerateSecret := NewUIGenerateSecret(debugData, logicData)
	uiSearchList.AddChoices(uiSearchFunc)
	return uiSearchList
}

func (m *ui_SearchList) Init() tea.Cmd { return nil }

func (m *ui_SearchList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	_list, _cmd := m.List.Update(msg)
	m.List, cmds = _list.(*ui.List), tea.Batch(_cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case common.KEY_backspace:
			return m, common.CMD_back()
		}
	}
	return m, cmds
}

func (m *ui_SearchList) View() string {
	s := fmt.Sprintf("Search List: %v\n\n", m.GetChosenIndex())
	s += m.List.View()
	return s
}

func (m *ui_SearchList) ChoiceTitle() string {
	return m.title
}
