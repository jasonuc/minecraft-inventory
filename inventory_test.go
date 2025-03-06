package main

import (
	"bytes"
	"testing"
)

func TestPlaceItem(t *testing.T) {
	tests := []struct {
		name           string
		inventory      Inventory
		index          int
		expectedCells  []Cell
		expectedHand   Cell
		expectedResult bool
	}{
		{
			name: "place first item in empty cell",
			inventory: Inventory{
				Cells: []Cell{{ItemId: EmptyCell, Amount: 0}},
				Hand:  Cell{ItemId: 1, Amount: 5},
			},
			index: 0,
			expectedCells: []Cell{
				{ItemId: 1, Amount: 5},
			},
			expectedHand:   Cell{ItemId: EmptyCell, Amount: 0},
			expectedResult: true,
		},
		{
			name: "add items to a cell with the same item",
			inventory: Inventory{
				Cells: []Cell{{ItemId: 1, Amount: 5}},
				Hand:  Cell{ItemId: 1, Amount: 3},
			},
			index: 0,
			expectedCells: []Cell{
				{ItemId: 1, Amount: 8},
			},
			expectedHand:   Cell{ItemId: EmptyCell, Amount: 0},
			expectedResult: true,
		},
		{
			name: "swap items in a cell with items in the hand",
			inventory: Inventory{
				Cells: []Cell{{ItemId: 2, Amount: 5}},
				Hand:  Cell{ItemId: 1, Amount: 3},
			},
			index: 0,
			expectedCells: []Cell{
				{ItemId: 1, Amount: 3},
			},
			expectedHand:   Cell{ItemId: 2, Amount: 5},
			expectedResult: true,
		},
		{
			name: "invalid index",
			inventory: Inventory{
				Cells: []Cell{{ItemId: 1, Amount: 5}},
				Hand:  Cell{ItemId: 1, Amount: 3},
			},
			index:          -1,
			expectedCells:  []Cell{{ItemId: 1, Amount: 5}},
			expectedHand:   Cell{ItemId: 1, Amount: 3},
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.inventory.placeItem(tt.index)
			if result != tt.expectedResult {
				t.Errorf("expected result %v, got %v", tt.expectedResult, result)
			}
			for i, cell := range tt.inventory.Cells {
				if cell != tt.expectedCells[i] {
					t.Errorf("expected cell %+v, got %+v", tt.expectedCells[i], cell)
				}
			}
			if tt.inventory.Hand != tt.expectedHand {
				t.Errorf("expected hand %+v, got %+v", tt.expectedHand, tt.inventory.Hand)
			}
		})
	}
}

func TestPrintInventory(t *testing.T) {
	inventory := Inventory{
		Cells: []Cell{
			{ItemId: 1, Amount: 5},
			{ItemId: 2, Amount: 3},
		},
	}

	var buf bytes.Buffer
	inventory.PrintIventory(&buf)
	expectedOutput := "cell: 1, 5\ncell: 2, 3\n"
	if buf.String() != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, buf.String())
	}
}
