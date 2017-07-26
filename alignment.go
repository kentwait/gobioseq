package gobioseq

import (
	"sort"
)

type Alignment interface {
	Init() // Initializes Alignment
	Keys() []string
	Values() [][]string
	Ids() []string
	Sequences() [][]string
	Col(cols ...int) [][]string
	Append(id string, seq []string) bool
	Remove(id string) bool
	Pop(id string) []string
	SubsetById(ids ...string) Alignment
	DropAnyGap() Alignment
	DropAllGap() Alignment
	// ResampleCols() Alignment
	ReorderBy(l []string) Alignment
	// PSSM() [][]int
	// ConsensusSequence() []string
}

type alignment struct {
	Name        string
	Description string
	seqType     string
	charSize    int
	ids         []string
	sequences   [][]string
	sequencesT  [][]*string // transpose of sequences, values are pointers to sequences elements
	mapping     map[string]*[][]string
}

func (a *alignment) Init() {
	a.mapping = make(map[string]*[][]string)
	a.sequencesT = make([][]*string, len(a.sequences[0]))
	for i := range a.sequencesT {
		a.sequencesT[i] = make([]*string, len(a.sequences))
	}
	for i, row := range a.sequences {
		for j, value := range row {
			a.sequencesT[j][i] = &value
		}
	}
}

func (a *alignment) Keys() []string {
	return a.ids
}

func (a *alignment) Values() [][]string {
	return a.sequences
}

func (a *alignment) Ids() []string {
	return a.ids
}

func (a *alignment) Sequences() [][]string {
	return a.sequences
}

func (a *alignment) Col(cols ...int) [][]string {
	var res = make([][]string, a.sequences)
	for i := range res {
		res[i] = make([]string, len(cols))
	}

	for _, colId := range cols {
		for _, rowVal := range a.sequencesT[colId] {
			for rowId, ptr := range rowVal {
				res[rowId][col] = *ptr
			}
		}
	}
	return res
}

func (a *alignment) Append(id string, seq []string) bool {
	a.ids = append(a.ids, id)
	a.sequences = append(a.sequences)
	i := len(a.sequences)
	row := a.sequences[i]
	for j, value := range row {
		a.sequencesT[j] = append(a.sequencesT[j], &value)
	}
}

func (a *alignment) Remove(id string) bool {
	if FindID(id, a.ids) == false {
		return false
	}
	var index int
	for i, v := range a.ids {
		if v == id {
			index = i
			break
		}
	}
	copy(a.ids[index:], a.ids[index+1:])
	a.ids[len(a.ids)-1] = nil
	a.ids = a.ids[:len(a.ids)-1]

	copy(a.sequences[index:], a.sequences[index+1:])
	a.sequences[len(a.sequences)-1] = nil
	a.sequences = a.sequences[:len(a.sequences)-1]
}

func (a *alignment) SubsetById(ids ...string) Alignment {
	newIds := ids
	newSequences := make([][]string, len(ids))
	for i, id := range ids {
		newSequences[i] = a.sequences[id]
	}
	return NewAlignment(a.Name, a.Description, a.charSize, newIds, newSequences)
}

func (a *alignment) DropAnyGap() Alignment {

}

func NewAlignment(name, desc, seqType string, charSize int, ids []string, sequences [][]string) Alignment {
	return &alignment{
		Name:        name,
		Description: desc,
		seqType:     seqType,
		charSize:    charSize,
		ids:         ids,
		sequences:   sequences,
	}
}

func FindID(id string, s []string) bool {
	sortedIds := s
	sort.Strings(sortedIds)
	i := sort.Search(len(sortedIds), func(i int) bool { return sortedIds[i] >= id })
	if i < len(sortedIds) && sortedIds[i] == id {
		return true
	}
	return false
}
