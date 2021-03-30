package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"go.i3wm.org/i3/v4"
	"log"
	"sort"
	"time"
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
					color.Blue("  %s ", p.Name)
				} else {
					fmt.Printf("  %s \n", p.Name)
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
	Name    string
	Focused bool
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

		for _, sn := range found.Nodes {
			program := Program{Name: sn.Name, Focused: sn.Focused}
			wo.Programs = append(wo.Programs, program)
		}

		workspaceOverviews = append(workspaceOverviews, wo)
	}

	overview := Overview{
		WorkspaceOverviews: workspaceOverviews,
	}

	overview.printFormatted()
}

func main() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		collect()
		time.Sleep(1000 * time.Millisecond)
	}

}
