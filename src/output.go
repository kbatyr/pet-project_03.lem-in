package lemin

import (
	"fmt"
	"strconv"
)

// Распределение муравьев по выбранным путям
func (p *Allpaths) AntsAllocation(ants int) []int {

	a := make([]int, len(p.OptPath))

	for ants != 0 {

		for i, path := range p.OptPath {
			num := p.Tick - (len(path.P) - 1)

			if num > ants {
				num = ants
			}
			ants -= num

			a[i] += num
		}
	}
	return a
}

// Печать передвижения муравьев по путям
func (p *Allpaths) Output(ants int) {

	out := make([][]string, (p.Tick - 1))
	antName := 1
	locations := p.AntsAllocation(ants)
	p.removeStart()

	// цикл по кол-ву путей в комбинации
	for i := range locations {

		//цикл по кол-ву муравьев для каждого пути
		for j := 0; j < locations[i]; j++ {

			// пробегаемся по комнатам пути
			for k := range p.OptPath[i].P {

				out[k+j] = append(out[k+j], "L"+strconv.Itoa(antName)+"-"+p.OptPath[i].P[k].Key)
			}
			antName++
		}
	}
	for _, lvl := range out {
		for _, v := range lvl {
			fmt.Print(v + " ")
		}
		fmt.Println()
	}
}

// удаляем стартовые комнаты в путях
func (p *Allpaths) removeStart() {
	for _, path := range p.OptPath {
		path.P = path.P[1:]
	}
}

// Printing the migration of ants with only start and end rooms (only 1 step migration)
func (p *Allpaths) PrintOneStep(ants int) {
	out := []string{}
	antsName := 1
	locations := p.AntsAllocation(ants)
	p.removeStart()

	defer fmt.Println()

	for j := 0; j < locations[0]; j++ {

		for i := range p.OptPath[0].P {
			out = append(out, "L"+strconv.Itoa(antsName)+"-"+p.OptPath[0].P[i].Key)
		}
		antsName++
	}

	for _, v := range out {
		fmt.Print(v, " ")
	}
}
