package main

import (
	"fmt"
	"go.i3wm.org/i3/v4"
	"log"
	"sort"
)

type Overview struct {
	WorkspaceOverviews []WorkspaceOverview
}

func (overview Overview) format() string {
	m := make(map[string][]WorkspaceOverview)

	for _, w := range overview.WorkspaceOverviews {
		m[w.Screen] = append(m[w.Screen], w)
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var output string
	for _, k := range keys {
		output += fmt.Sprintf("%s \n", k)
		for _, w := range m[k] {
			output += "\n"
			output += fmt.Sprintf("- %s \n", w.Name)
			for _, p := range w.Programs {
				output += fmt.Sprintf("  %s \n", p)
			}
		}
		output += "\n"
	}
	return output
}

type WorkspaceOverview struct {
	Screen   string
	Name     string
	Programs []string
}

func main() {
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
			wo.Programs = append(wo.Programs, sn.Name)
		}

		workspaceOverviews = append(workspaceOverviews, wo)
	}

	overview := Overview{
		WorkspaceOverviews: workspaceOverviews,
	}

	fmt.Printf("%s", overview.format())

}
