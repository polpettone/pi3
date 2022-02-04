package main

import (
	"fmt"
	"log"
	"sort"

	"go.i3wm.org/i3/v4"
)

type Overview struct {
	WorkspaceOverviews []WorkspaceOverview
}

func (overview Overview) printFormatted() {
	m := make(map[string][]WorkspaceOverview)

	for _, w := range overview.WorkspaceOverviews {
		m[w.Screen] = append(m[w.Screen], w)
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s \n", k)
		for _, w := range m[k] {
			fmt.Printf("\n")
			fmt.Printf("- %s \n", w.Name)
			for _, p := range w.Programs {
				if p.Focused {
					fmt.Printf("  %s \n", "-----------------")
					fmt.Printf("  %s : %s \n", p.Title, p.Instance)
					fmt.Printf("  %s \n", "-----------------")
				} else {
					fmt.Printf("  %s : %s \n", p.Title, p.Instance)
				}
			}
		}
		fmt.Printf("\n")
	}
}

type WorkspaceOverview struct {
	Screen   string
	Name     string
	Programs []Program
}

type Program struct {
	Name     string
	Instance string
	Title    string
	Focused  bool
}

func collect() {

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	workspaces, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	var workspaceOverviews []WorkspaceOverview

	for _, w := range workspaces {
		found := tree.Root.FindChild(func(n *i3.Node) bool {
			return int64(n.ID) == int64(w.ID)
		})

		wo := WorkspaceOverview{
			Screen: w.Output,
			Name:   w.Name,
		}

		//TODO: walk down the tree
		for _, sn := range found.Nodes {
			if len(sn.Nodes) == 0 {
				program := Program{
					sn.Name,
					sn.WindowProperties.Instance,
					sn.WindowProperties.Title,
					sn.Focused,
				}
				wo.Programs = append(wo.Programs, program)
			} else {

			}
		}

		workspaceOverviews = append(workspaceOverviews, wo)
	}

	overview := Overview{
		WorkspaceOverviews: workspaceOverviews,
	}

	overview.printFormatted()
}

func printI3Tree() {
	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}
	printNode(tree.Root)
}

func printNode(node *i3.Node) {
	if len(node.Nodes) == 0 {

		if node.Type == "con" && node.Name != "" {
			fmt.Printf("--- %s \n", node.Name)
		}

	} else {
		for _, n := range node.Nodes {
			if node.Type == "workspace" {
				fmt.Println()
				fmt.Printf("%s:  %s \n", "Workspace", node.Name)
			}
			if node.Type == "output" {
				fmt.Println()
				fmt.Printf("%s:  %s \n", "Screen", node.Name)
			}
			printNode(n)
		}
	}
}

func main() {
	printI3Tree()
}
