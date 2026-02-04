package goass

import (
	"fmt"

	"github.com/Mericusta/go-assistant/src/common"
	"github.com/Mericusta/go-assistant/src/ui"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	Record_TriggerTick = iota + 1
)

type ui_MainOperationList struct {
	*ui.List

	debugData *data_Debug
	logicData *data_Logic
	uiData    *data_UI
}

func NewUIMainOperationList(debugData *data_Debug, logicData *data_Logic, uiData *data_UI) *ui_MainOperationList {
	uiMainOperationList := &ui_MainOperationList{
		List:      ui.NewList(0),
		debugData: debugData,
		logicData: logicData,
		uiData:    uiData,
	}
	uiGenerate := NewUISearch(debugData, logicData)
	// uiSearch := NewUISearch(debugData, logicData)
	// uiGrowList := NewUIGrowList(debugData, logicData)
	// uiSaveData := NewUISaveData(debugData, logicData)
	uiMainOperationList.AddChoices(uiGenerate)
	return uiMainOperationList
}

func (m *ui_MainOperationList) Init() tea.Cmd { return nil }

func (m *ui_MainOperationList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	_list, _cmd := m.List.Update(msg)
	m.List, cmds = _list.(*ui.List), tea.Batch(_cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case common.KEY_interrupt, common.KEY_quit:
			return m, common.CMD_back()
		}
	}
	return m, cmds
}

func (m ui_MainOperationList) View() string {
	s := fmt.Sprintf("Main Operation List: %v\n\n", m.GetChosenIndex())
	s += m.List.View()
	return s
}
