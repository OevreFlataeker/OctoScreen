package ui

var idleMenuPanelInstance *idleMenuPanel

type idleMenuPanel struct {
	CommonPanel
}

func IdleMenuPanel(ui *UI, parent Panel) Panel {
	if idleMenuPanelInstance == nil {
		m := &idleMenuPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		idleMenuPanelInstance = m
	}

	return idleMenuPanelInstance
}

func (m *idleMenuPanel) initialize() {
	defer m.Initialize()

	m.Grid().Attach(MustButtonImageStyle("Network", "network.svg", "color1", m.showNetwork), 1, 0, 1, 1)
	m.Grid().Attach(MustButtonImageStyle("ToolChanger", "toolchanger.svg", "color2", m.showToolchanger), 2, 0, 1, 1)
	m.Grid().Attach(MustButtonImageStyle("Control", "control.svg", "color4", m.showControl), 3, 0, 1, 1)
	m.Grid().Attach(MustButtonImageStyle("System", "settings.svg", "color3", m.showSystem), 4, 0, 1, 1)
	m.Grid().Attach(MustButtonImageStyle("Temperature", "heat-up.svg", "color4", m.showTemperature), 1, 1, 1, 1)
}

func (m *idleMenuPanel) showTemperature() {
	m.UI.Add(TemperaturePanel(m.UI, m))
}

func (m *idleMenuPanel) showControl() {
	m.UI.Add(ControlPanel(m.UI, m))
}

func (m *idleMenuPanel) showNetwork() {
	m.UI.Add(NetworkPanel(m.UI, m))
}

func (m *idleMenuPanel) showToolchanger() {
	m.UI.Add(ToolchangerPanel(m.UI, m))
}

func (m *idleMenuPanel) showSystem() {
	m.UI.Add(SystemPanel(m.UI, m))
}
