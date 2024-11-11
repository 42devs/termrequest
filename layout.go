package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// colorscheme

var (
	white   = "#ffffff"
	gray100 = "#f8f9fa"
	gray200 = "#e9ecef"
	gray300 = "#dee2e6"
	gray400 = "#ced4da"
	gray500 = "#adb5bd"
	gray600 = "#6c757d"
	gray700 = "#495057"
	gray800 = "#343a40"
	gray900 = "#170229"
	black   = "#000000"

	blue       = "#1ba2f6"
	indigo     = "#6610f2"
	purple     = "#6f42c1"
	pink       = "#ea39b8"
	red        = "#e44c55"
	orange     = "#f1b633"
	yellow     = "#ffc107"
	green      = "#3cf281"
	teal       = "#3f81a2"
	cyan       = "#32fbe2"
	lightColor = "#44d9e8"
	mainBg     = "#1a0933"

	primary   = lipgloss.Color(purple)
	secondary = lipgloss.Color(pink)
	success   = lipgloss.Color(green)
	info      = lipgloss.Color(blue)
	warning   = lipgloss.Color(yellow)
	danger    = lipgloss.Color(red)
	light     = lipgloss.Color(lightColor)
	dark      = lipgloss.Color(gray900)

	bodybg    = lipgloss.Color(mainBg)
	bodyColor = lipgloss.Color(cyan)
	linkColor = bodyColor
)

var (
	commonStyle = lipgloss.NewStyle().
			Foreground(bodyColor).
			Background(bodybg).
			Border(lipgloss.ThickBorder(), true).
			BorderForeground(bodyColor)

	sidebarStyle = commonStyle.
			Padding(1, 2).
			Width(20)

	navbarStyle = commonStyle.
			Padding(0, 1).
			Height(3).
			Align(lipgloss.Left)

	bottomStyle = commonStyle.
			Padding(0, 1).
			Height(2).
			Align(lipgloss.Left)

	contentStyle = commonStyle.
			Padding(1).
			Height(40).
			Width(15).Align(lipgloss.Left)
)

func mainLayout() {

	sidebar := sidebarStyle.Render("Sidebar\n")
	navbar := navbarStyle.Render("Term Request")
	bottom := bottomStyle.Render("This is a bottom")
	contentRight := contentStyle.Render("Rigth")
	contentLeft := contentStyle.Render("left")
	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		contentRight,
		contentLeft,
	)

	layout := lipgloss.JoinVertical(lipgloss.Top,
		navbar,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			sidebar,
			content,
		),
		bottom,
	)

	fmt.Println(layout)
}
