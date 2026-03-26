package ui

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/schollz/progressbar/v3"
)

var (
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF4672")).
			Bold(true)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A0A0A0"))

	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Bold(true)

	msgBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7B2FBE")).
			Padding(0, 2).
			MarginTop(1).
			MarginBottom(1)

	doneStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#04B575")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#04B575")).
			Padding(0, 2)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A0A0A0")).
			Italic(true).
			MarginLeft(1)

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7B2FBE")).
			Bold(true)
)

func Banner() {
	art := []string{
		"██████╗ ██╗   ██╗██████╗       ███████╗██╗   ██╗██████╗ ",
		"██╔══██╗██║   ██║██╔══██╗      ██╔════╝██║   ██║██╔══██╗",
		"██████╔╝██║   ██║██████╔╝█████╗███████╗██║   ██║██████╔╝",
		"██╔═══╝ ██║   ██║██╔══██╗╚════╝╚════██║██║   ██║██╔══██╗",
		"██║     ╚██████╔╝██████╔╝      ███████║╚██████╔╝██████╔╝",
		"╚═╝      ╚═════╝ ╚═════╝       ╚══════╝ ╚═════╝ ╚═════╝ ",
	}

	gradient := []string{
		"#0000FF", "#0800FF", "#1000FF", "#1800FF", "#2000FF",
		"#2800FF", "#3000FF", "#3800FF", "#4000FF", "#4800FF",
		"#5000FF", "#5800FF", "#6000FF", "#6800FF", "#7000FF",
		"#7800FF", "#8000FF", "#8800FF", "#9000FF", "#9800FF",
		"#A000FF", "#A800FF", "#B000FF", "#B800EE", "#C000DD",
		"#C800CC", "#CC00BB", "#CC00AA", "#CC0099", "#CC0088",
		"#CC0077", "#CC0066", "#CC0055", "#CC0044", "#CC0033",
		"#DD0033", "#EE0033", "#FF0033", "#FF0022", "#FF0011",
		"#FF0000", "#EE0011", "#DD0022", "#CC0033", "#BB0044",
		"#AA0055", "#990066", "#880077", "#770088", "#660099",
	}

	fmt.Println()
	for _, line := range art {
		colored := ""
		colorIdx := 0
		for _, ch := range line {
			color := gradient[colorIdx%len(gradient)]
			colored += lipgloss.NewStyle().
				Foreground(lipgloss.Color(color)).
				Bold(true).
				Render(string(ch))
			if ch != ' ' {
				colorIdx++
			}
		}
		fmt.Println("  " + colored)
	}
	fmt.Println()
	fmt.Println(subtitleStyle.Render("  LocalStack • SNS • SQS • Go • Docker"))
	fmt.Println(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7B2FBE")).
		Render("  ────────────────────────────────────────"))
	fmt.Println()
}
func Bar(label string) {
	pb := progressbar.NewOptions(100,
		progressbar.OptionSetDescription(fmt.Sprintf("  %-40s", label)),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "█",
			SaucerPadding: "░",
			BarStart:      "│",
			BarEnd:        "│",
		}),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetPredictTime(false),
	)
	for i := 0; i < 100; i++ {
		pb.Add(1)
		time.Sleep(8 * time.Millisecond)
	}
}

func Ok(label, value string) {
	fmt.Println(successStyle.Render("  ✓ ") + labelStyle.Render(label+": ") + valueStyle.Render(value))
}

func Fail(label, errMsg string) {
	fmt.Println(errorStyle.Render("  ✗ " + label + ": " + errMsg))
}

func Input(prompt string, dest *string) {
	promptBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7B2FBE")).
		Padding(0, 2).
		Render(inputStyle.Render("➜  " + prompt))

	fmt.Println(promptBox)
	fmt.Print("  ╰─▶ ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*dest = scanner.Text()

	if *dest == "" {
		*dest = "Hello from AWS Pub/Sub Demo!"
	}
}

func MessageBox(direction, message string) {
	fmt.Println(msgBoxStyle.Render("  " + direction + "\n\n  " + message))
}

func Done() {
	fmt.Println()
	fmt.Println(doneStyle.Render(" ✓ SNS → SQS flow completed successfully"))
	fmt.Println()
}

func Spacer() {
	fmt.Println()
}
