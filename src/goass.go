package goass

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Mericusta/go-assistant/src/common"

	tea "github.com/charmbracelet/bubbletea"
)

type logic_MainProgram struct {
	// debug data
	debugData *data_Debug

	// logic data
	logicData *data_Logic

	// ui data
	uiData *data_UI
}

func NewLogicMainProgram(saveDataPath string, debugMode bool) *logic_MainProgram {
	var debugData *data_Debug
	if debugMode {
		debugData = newDebugData()
	}
	logicData := newLogicData(saveDataPath)
	uiData := newDataUI()
	logicMainProgram := &logic_MainProgram{debugData: debugData, logicData: logicData, uiData: uiData}
	logicMainProgram.uiData.AppendModelStack(NewUIMainOperationList(debugData, logicData, uiData))
	return logicMainProgram
}

func (m *logic_MainProgram) String() string {
	if m.debugData == nil {
		return ""
	}
	return fmt.Sprintf(`time: %v
- modelStack: %v
- triggerMap: %v
`, time.Now().Format(time.DateTime),
		m.uiData.ModelStackDesc(),
		m.debugData.triggerCmdTickMap,
	)
}

func (m *logic_MainProgram) Init() tea.Cmd {
	return m.uiData.ModelTop().Init()
}

func (m *logic_MainProgram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 从栈顶开始处理消息，如果某一层不拦截，则需要在 model 的 update 中实现

	// 消息
	// - 按键消息：只有栈顶处理
	// - 非按键消息，层层传递

	var cmds []tea.Cmd
	switch msg.(type) {
	case tea.KeyMsg:
		_, _cmd := m.uiData.ModelTop().Update(msg)
		if _cmd != nil {
			cmds = append(cmds, _cmd)
		}
	default:
		// 迭代 model 栈
		m.uiData.RangeModelStack(func(m tea.Model) bool {
			_, _cmd := m.Update(msg)
			cmds = append(cmds, _cmd)
			return true
		})

		// 栈底处理消息
		switch msg := msg.(type) {
		case common.MSG_enter:
			m.uiData.AppendModelStack(msg.Model())
			cmds = append(cmds, msg.Model().Init())
		case common.MSG_back:
			if m.uiData.ModelStackLen() == 1 {
				return m, tea.Quit
			}
			m.uiData.PopModelStack()
			cmds = append(cmds, m.uiData.ModelTop().Init())
		}
	}

	if len(cmds) > 0 {
		return m, tea.Batch(cmds...)
	}

	return m, nil
}

func (m *logic_MainProgram) View() string {
	modelView := m.uiData.ModelTop().View()
	return fmt.Sprintf("logic main program\n%v\ntop model\n\n%v\n", m, modelView)
}

const debugMode = true

func GoAssistant() {
	p := tea.NewProgram(NewLogicMainProgram(os.Getenv("SAVE_DATA_PATH"), debugMode), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
