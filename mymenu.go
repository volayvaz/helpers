package helpers

import (
	"fmt"
)

type Menu struct {
	Prompt    string
	CursorPos int
	MenuItems []*MenuItem
}

type MenuItem struct {
	Text string
	ID   int
}

func NewMenu(prmt string, itemCount int) *Menu {
	return &Menu{
		Prompt:    prmt,
		CursorPos: 0,
		MenuItems: make([]*MenuItem, 0, itemCount),
	}
}

func (mnu *Menu) AddItem(txt string, id int) *Menu {
	menuItem := &MenuItem{
		Text: txt,
		ID:   id,
	}

	mnu.MenuItems = append(mnu.MenuItems, menuItem)
	return mnu
}

func (mnu *Menu) DisplayMenu() {
	fmt.Printf("\033[%dA", len(mnu.MenuItems)+1)
	fmt.Printf("%s\n", mnu.Prompt)
	var newline = "\n"
	var cursor = "  "
	for i := 0; i < len(mnu.MenuItems); i++ {
		newline = "\n"
		cursor = "  "
		if i == len(mnu.MenuItems) {
			newline = " "
		}
		if i == mnu.CursorPos {
			cursor = "> "
		}
		fmt.Printf("%s %s%s", cursor, mnu.MenuItems[i].Text, newline)
	}
}
