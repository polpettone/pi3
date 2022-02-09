package cmd

import (
	"fmt"
	"log"

	"github.com/bclicn/color"
	"go.i3wm.org/i3/v4"
)

func PrintOverview(showInstanceNames bool) {

	focusedWorkspaceName, err := getFocusedWorkspaceName()
	if err != nil {
		log.Fatal(err)
	}

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	outputs := []*i3.Node{}
	outputs = allNodesOfType("output", tree.Root, outputs)

	for _, output := range outputs {

		if output.Name != "__i3" {

			fmt.Printf("%s %s \n", MONITOR_ICON, output.Name)

			workspaces := []*i3.Node{}
			workspaces = allNodesOfType("workspace", output, workspaces)

			for _, workspace := range workspaces {

				icon := WORKSPACE_ICON
				if workspace.Name == focusedWorkspaceName {
					icon = color.Green(WORKSPACE_ICON)
				}

				fmt.Printf("  %s %s \n", icon, workspace.Name)

				contents := []*i3.Node{}
				contents = allContentNodes(workspace, contents)

				for _, content := range contents {

					icon := iconFor(content.WindowProperties.Instance)
					title := truncateString(content.WindowProperties.Title, 80)

					if content.Focused {
						icon = color.Red(icon)
						title = color.LightBlue(title)
					}

					if showInstanceNames {
						fmt.Printf("    %s %d %-18s %s\n",
							icon,
							content.ID,
							content.WindowProperties.Instance,
							title,
						)
					} else {
						fmt.Printf("    %s %s\n",
							icon,
							title,
						)
					}
				}
			}
		}
	}
}

const TERMINAL_ICON = ""
const MONITOR_ICON = ""
const GOPHER_ICON = ""
const WORKSPACE_ICON = ""
const BROWSER_ICON = ""
const ZOOM_ICON = ""
const IDE_ICON = ""
const NYXT_ICON = ""
const DEFAULT_ICON = ""
const SLACK_ICON = ""

func iconFor(class string) string {

	iconMap := map[string]string{
		"terminator":     TERMINAL_ICON,
		"zoom":           ZOOM_ICON,
		"jetbrains-idea": IDE_ICON,
		"Navigator":      BROWSER_ICON,
		"nyxt":           NYXT_ICON,
		"slack":          SLACK_ICON,
	}

	icon, exist := iconMap[class]

	if exist {
		return icon
	} else {
		return DEFAULT_ICON
	}
}

func getFocusedWorkspaceName() (string, error) {
	ws, err := i3.GetWorkspaces()
	if err != nil {
		return "", err
	}
	focusedWorkspaceName := "1"
	for _, w := range ws {
		if w.Focused {
			focusedWorkspaceName = w.Name
		}
	}
	return focusedWorkspaceName, err
}

func truncateString(value string, max int) string {
	if len(value) > max {
		return value[:max]
	} else {
		return value
	}
}

func allContentNodes(node *i3.Node, nodes []*i3.Node) []*i3.Node {
	if node.Type == "con" && node.Name != "" {
		nodes = append(nodes, node)
	} else {
		for _, n := range node.Nodes {
			nodes = allContentNodes(n, nodes)
		}

		for _, n := range node.FloatingNodes {
			nodes = allContentNodes(n, nodes)
		}
	}
	return nodes
}

func allNodesOfType(nodeType string, node *i3.Node, nodes []*i3.Node) []*i3.Node {
	if string(node.Type) != nodeType {
		for _, n := range node.Nodes {
			nodes = allNodesOfType(nodeType, n, nodes)
		}
	} else {
		nodes = append(nodes, node)
	}
	return nodes
}
