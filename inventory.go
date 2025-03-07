package main

import (
	"fmt"
	"image/color"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const EmptyCell = 0

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

func (i *Inventory) Draw(screen *ebiten.Image) {
	rowSize := 9
	cellSize := 32
	gapSize := 2

	for idx, _ := range i.Cells {
		gridX := idx % rowSize
		gridY := idx / rowSize
		pixelX := gridX * (cellSize + gapSize)
		pixelY := gridY * (cellSize + gapSize)

		vector.DrawFilledRect(
			screen,
			float32(pixelX),
			float32(pixelY),
			float32(cellSize),
			float32(cellSize),
			color.RGBA{20, 20, 20, 255},
			true)
	}

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

func (i *Inventory) PrintIventory(out io.Writer) {
	for _, cell := range i.Cells {
		fmt.Fprintf(out, "cell: %d, %d\n", cell.ItemId, cell.Amount)
	}
}
