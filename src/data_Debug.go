package goass

type data_Debug struct {
	triggerCmdTickMap map[string]int
}

func newDebugData() *data_Debug {
	return &data_Debug{triggerCmdTickMap: make(map[string]int)}
}
