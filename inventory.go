package main

import "fmt"

const EmptyCell = -1

type Cell struct {
	ItemId int
	Amount uint // usinged int because a cell can't hold negative items
}

type Inventory struct {
	Cells []Cell
	Hand  Cell
}

func resetCell(cell *Cell) {
	cell.ItemId = EmptyCell
	cell.Amount = 0
}

func (i *Inventory) placeItem(index int) bool {
	/* there are 3 possibilities:
	 * putting the first item in a cell
	 * adding items to a cell
	 * swapping items in a cell with items in the hand
	 */

	if index < 0 || index > len(i.Cells)-1 {
		return false
	}

	cell := &i.Cells[index]

	switch {
	case cell.ItemId == EmptyCell:
		cell.ItemId = i.Hand.ItemId
		cell.Amount = i.Hand.Amount
		resetCell(&i.Hand)
		return true
	case cell.ItemId == i.Hand.ItemId:
		cell.Amount += i.Hand.Amount
		resetCell(&i.Hand)
		return true
	default:
		i.Hand.ItemId, cell.ItemId = cell.ItemId, i.Hand.ItemId
		cell.Amount, i.Hand.Amount = i.Hand.Amount, cell.Amount
		return true
	}
}

func (i *Inventory) PrintIventory() {
	for _, cell := range i.Cells {
		fmt.Printf("cell: %d, %d\n", cell.ItemId, cell.Amount)
	}
}
