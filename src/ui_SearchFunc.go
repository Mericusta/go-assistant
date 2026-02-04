package goass

import (
	"fmt"

	"github.com/Mericusta/go-assistant/src/common"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ui_SearchFunc struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode

	debugData *data_Debug
	logicData *data_Logic

	title   string
	content string
}

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func NewUISearchFunc(debugData *data_Debug, logicData *data_Logic) *ui_SearchFunc {
	uiSearchFunc := &ui_SearchFunc{
		logicData: logicData,
		title:     "Func",
		content:   "this is search func model",
	}

	uiSearchFunc.inputs = make([]textinput.Model, 2)

	var t textinput.Model
	for i := range uiSearchFunc.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "func ident"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "project root"
		}

		uiSearchFunc.inputs[i] = t
	}

	return uiSearchFunc
}

func (m *ui_SearchFunc) Init() tea.Cmd { return textinput.Blink }

func (m *ui_SearchFunc) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	_textInput, _cmd := m.textInput.Update(msg)
	m.textInput, cmds = &_textInput, tea.Batch(_cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case common.KEY_quit:
			return m, common.CMD_back()
		}
	}
	return m, cmds
}

func (m *ui_SearchFunc) View() string {
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func (m *ui_SearchFunc) ChoiceTitle() string {
	return m.title
}
