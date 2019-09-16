package main

import "fmt"

//算法导论 11.1

//DAS Direct-address tables
type DAS struct {
	T [100]Data
}

// Data ..
type Data struct {
	key       int
	satellite string
}

// Search ...
func (das *DAS) Search(k int) Data {
	return das.T[k]
}

// Insert ...
func (das *DAS) Insert(x Data) {
	das.T[x.key] = x
}

// Delete ...
func (das *DAS) Delete(k int) {
	das.T[k] = Data{0, ""}
}

func main() {
	das1 := DAS{}
	das1.Insert(Data{1, "a"})
	das1.Insert(Data{2, "b"})
	das1.Insert(Data{3, "c"})
	das1.Insert(Data{4, "d"})
	das1.Insert(Data{5, "e"})

	fmt.Println(das1)
	fmt.Println(das1.Search(3))
	das1.Delete(2)
	fmt.Println(das1)
}
