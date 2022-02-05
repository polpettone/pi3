package main

import (
	"fmt"
	"log"

	"go.i3wm.org/i3/v4"
)

const TERMINAL_ICON = ""
const MONITOR_ICON = ""
const GOPHER_ICON = ""
const WORKSPACE_ICON = ""

func printOverview() {

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
				fmt.Printf("  %s %s \n", WORKSPACE_ICON, workspace.Name)

				contents := []*i3.Node{}
				contents = allContentNodes(workspace, contents)

				for _, content := range contents {
					fmt.Printf("    %s %-30s \n",
						TERMINAL_ICON,
						truncateString(content.WindowProperties.Title, 30))
				}
			}
		}
	}
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

func main() {
	printOverview()
}
