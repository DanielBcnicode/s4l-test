package booking

import (
	"errors"
	"github.com/danielbcnicode/timeslot/internal"
	"log"
	"sort"
	"time"
)

var (
	IndexNotFoundError = errors.New("no more indexes found")
)

type Maximizer interface {
	Maximize(bookings []Request) (MaximizeResponse, error)
}

type Maximize struct{}

func NewMaximizer() *Maximize {
	return &Maximize{}
}

func (m *Maximize) Maximize(bookings []Request) (MaximizeResponse, error) {
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i].StartDate().Unix() < bookings[j].StartDate().Unix()
	})
	// Create the tree

	rootNode := internal.NewNode(nil)
	currentCounter := -1
	createTree(bookings, rootNode, currentCounter)
	// Get Leafs
	leafs := rootNode.GetLeafs()
	// Choose the best case
	bestProfit, bestLeaf := getBestLeaf(leafs, rootNode)

	var ids []string
	curNode := bestLeaf
	nodes := 0
	minPPN := bestLeaf.ProfitPerNight()
	maxPPN := bestLeaf.ProfitPerNight()
	sumPPN := float32(0)
	for curNode != rootNode {
		ids = append([]string{curNode.Data().(string)}, ids...)
		nodes += 1
		if curNode.ProfitPerNight() > maxPPN {
			maxPPN = curNode.ProfitPerNight()
		}
		if curNode.ProfitPerNight() < minPPN {
			minPPN = curNode.ProfitPerNight()
		}
		sumPPN += curNode.ProfitPerNight()

		curNode = curNode.Father()
	}

	log.Printf("Best Profit %.3f and Leaf %v", bestProfit, bestLeaf)

	// Return data
	return MaximizeResponse{
		RequestIDs:   ids,
		TotalProfit:  bestProfit,
		AverageNight: internal.FloatRoundPrecision(sumPPN/float32(nodes), 2),
		MinNight:     minPPN,
		MaxNight:     maxPPN,
	}, nil
}

func getBestLeaf(leafs []*internal.Node, rootNode *internal.Node) (float32, *internal.Node) {
	var bestLeaf *internal.Node
	bestProfit := float32(0)
	for _, leaf := range leafs {
		curNode := leaf
		curProfit := float32(0)
		for curNode != rootNode {
			curProfit += curNode.Profit()
			curNode = curNode.Father()
		}
		if curProfit > bestProfit {
			bestProfit = curProfit
			bestLeaf = leaf
		}
	}

	return bestProfit, bestLeaf
}

func createTree(bookings []Request, father *internal.Node, currentCounter int) {
	var err error
	if currentCounter < 0 {
		currentCounter, err = findNodeNextToTime(0, bookings[0].StartDate(), bookings)
	} else {
		currentCounter, err = findNodeNextToTime(currentCounter, bookings[currentCounter].EndDate(), bookings)
	}
	if err == IndexNotFoundError {
		return // no more Nodes to add to current branch
	}

	brothers := findBrothers(currentCounter, bookings)
	brothers = append([]int{currentCounter}, brothers...)
	for _, brother := range brothers {
		// add node and data
		newFather := internal.NewNode(father)
		newFather.SetProfitPerNight(bookings[brother].ProfitPerNight)
		newFather.SetProfit(bookings[brother].Profit)
		newFather.SetData(bookings[brother].ID)
		createTree(bookings, newFather, brother)
	}
}

// findNodeNextToTime return the slide position for the next booking in the time.
// It assumes the bookings array is ordered by StartDate
func findNodeNextToTime(fromPos int, fromTime time.Time, bookings []Request) (int, error) {
	for i := fromPos; i < len(bookings); i++ {
		if bookings[i].StartDate().Unix() >= fromTime.Unix() {
			return i, nil
		}
	}

	return -1, IndexNotFoundError
}

// findBrothers returns the Request indexes of different time branches than the counter => Brother in the tree
func findBrothers(counter int, bookings []Request) []int {
	var incompatible []int
	for i := counter + 1; i < len(bookings); i++ {
		if bookings[counter].Overlaps(&bookings[i].DaySlot) && bookings[counter].StartDate().Unix() <= bookings[i].StartDate().Unix() {
			incompatible = append(incompatible, i)
		}
	}

	return incompatible
}
