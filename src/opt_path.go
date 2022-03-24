package lemin

// Находит комбинации непересекающихся путей
func (p *Allpaths) Combinations() {

	for k := range p.Paths {

		uniqeRooms := make(map[string]*Room)
		tmp := []*Path{}

		if intersectCheck(uniqeRooms, p.Paths[k]) {

			addToUniqeRooms(uniqeRooms, p.Paths[k])
			tmp = append(tmp, p.Paths[k])
		}

		for i := k + 1; i < len(p.Paths); i++ {

			if intersectCheck(uniqeRooms, p.Paths[i]) {

				addToUniqeRooms(uniqeRooms, p.Paths[i])
				tmp = append(tmp, p.Paths[i])
			} else {
				continue
			}
		}
		p.Combo = append(p.Combo, tmp)
	}
}

// Расчет кол-ва времени для перемещения муравьев из точки А в точку Б
func (p *Allpaths) Ticks(ants int) []int {

	ticks := []int{}

	for _, elem := range p.Combo {

		roomLen := 0
		roomTick := 0
		for _, paths := range elem {
			roomLen += len(paths.P)
		}
		roomTick = (ants + (roomLen - 1)) / len(elem)
		ticks = append(ticks, roomTick)
	}
	return ticks
}

// Оптимальный путь на основе кол-ва муравьев и времени на их перемещение
func (p *Allpaths) OptimalPath(ticks []int, g *Graph) {

	p.Tick = ticks[0]
	p.OptPath = p.Combo[0]

	// Если в комбинации присутствует путь сост. только из стратовой и конечной точки,
	// используем только данный путь, отсекая остальные пути в комбинации
	if p.OptPath[0].P[0] == g.Start && p.OptPath[0].P[1] == g.End {

		p.Combo[0] = p.Combo[0][0:1]
		p.OptPath = p.Combo[0]
	}

	for i, num := range ticks {
		if num < p.Tick {
			p.Tick = num
			p.OptPath = p.Combo[i]
		}
	}
}

// Проверка пересекаются ли комнаты путей
func intersectCheck(uniqeRooms map[string]*Room, path *Path) bool {

	for i := 1; i < len(path.P)-1; i++ {

		if _, ok := uniqeRooms[path.P[i].Key]; ok {
			return false
		}
	}
	return true
}

// Собирает в мапу уникальные комнаты
func addToUniqeRooms(uniqeRooms map[string]*Room, path *Path) {

	for _, room := range path.P {
		uniqeRooms[room.Key] = room
	}
}
