package exercise

import (
	"fmt"
)

// Create a Player struct (Name and Inventory => Slice of Item)
type Player struct {
	Name string
	Inventory []Item
}

// Create an Item struct (Name and Type)
type Item struct {
	Name string
    Type string
}

// Methods
func (p *Player) PickUpItem(item Item) {
	// Create a method for Player (PickUpItem) to add an item to their inventory
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s item!\n", p.Name, item.Name)
}

func (p *Player) DropItem(itemName string) {
	// Create a method for Player (DropItem) to add an item to their inventory
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s item!\n", p.Name, itemName)
   			return
		}
	}
	 fmt.Printf("%s does not have %s item in inventory.\n", p.Name, itemName)
}

func (p *Player) UseItem(itemName string) {
	// Create a method for Player (UseItem) to add an item to their inventory
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				fmt.Printf("%s used %s and feels rejuvenated!\n", p.Name, itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			} else {
				fmt.Printf("%s used %s item!\n", p.Name, itemName)
			}
			return
		}
	}
	fmt.Printf("%s does not have %s item in inventory.\n", p.Name, itemName)
}
