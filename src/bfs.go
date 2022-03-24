package lemin

// Находит все возможные комбинации путей с точки А до точки Б
func (g *Graph) BFS(start, end string) []*Path {

	g.Start = g.getVertex(start)
	g.End = g.getVertex(end)

	// Очередь из путей
	queue := [][]*Room{{g.Start}}

	// Собираем в ней все варианты путей
	var allPaths []*Path

	for len(queue) != 0 {
		curPath := queue[0]

		// Последняя комната пути
		lastRoom := curPath[len(curPath)-1]

		// Удаляем первый первый элемент из очереди
		queue = queue[1:]

		// Пробегаемся по смежным комнатам
		for _, adj := range lastRoom.Children {

			// Если смежная комната является конечной,
			// то добавляем полученный путь в срез всех комбинаций путей
			if adj == g.End {

				path := &Path{P: curPath}
				path.P = append(path.P, adj)
				allPaths = append(allPaths, path)
				continue
			}

			newPath := make([]*Room, len(curPath))
			copy(newPath, curPath)

			if existRoom(adj, newPath) {

				newPath = append(newPath, adj)
				queue = append(queue, newPath)

			} else {
				continue
			}
		}
	}
	return allPaths
}

// Проверят не повтоярется ли комната в пути
// т.е. проверка на наличие петли
func existRoom(r *Room, newPath []*Room) bool {

	for _, v := range newPath {
		if v == r {
			return false
		}
	}
	return true
}
