package lemin

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fname string) ([]string, error) {
	file, err := os.Open("examples/" + fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arrFile := []string{}
	for scanner.Scan() {
		arrFile = append(arrFile, scanner.Text())
	}

	if len(arrFile) == 0 {
		return arrFile, errors.New("invalid file: the file is empty")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arrFile, nil
}

func (data *Input) GetData(arrFile []string) error {
	var existStart, existEnd bool
	var ctr int

	for i, v := range arrFile {
		if i == 0 {
			a, err := strconv.Atoi(v)
			if err != nil || a <= 0 {
				return errors.New("ERROR: invalid data format, invalid format of Ants")
			}
			data.Ants = a
			continue
		}

		if v == "##start" && i != len(arrFile)-1 {
			if err := data.isValidRoom(arrFile[i+1], "start"); err != nil {
				return err
			}
			existStart = true
			ctr++
			continue
		}

		if v == "##end" && i != len(arrFile)-1 {
			if err := data.isValidRoom(arrFile[i+1], "end"); err != nil {
				return err
			}
			existEnd = true
			ctr++
			continue
		}

		if strings.Contains(v, " ") {
			if err := data.isValidRoom(v, ""); err != nil {
				return err
			}
			continue
		}

		if strings.Contains(v, "-") {
			tmp := strings.Split(v, "-")
			if len(tmp) != 2 {
				return errors.New("ERROR: invalid data format, invalid link between rooms")
			}

			tmpLinks := []string{}
			tmpLinks = append(tmpLinks, tmp[0], tmp[1])

			data.Links = append(data.Links, tmpLinks)
		}

	}

	if !existStart {
		return errors.New("ERROR: the start room doesn't exist")
	}
	if !existEnd {
		return errors.New("ERROR: the end room doesn't exist")
	}
	if ctr > 2 {
		return errors.New("ERROR: more than one start/end room found")
	}

	return nil
}

func (data *Input) isValidRoom(v, s string) error {

	tmp := strings.Split(v, " ")
	if len(tmp) != 3 {
		return errors.New("ERROR: invalid format of the room")
	}

	if strings.Contains(tmp[0], " ") || tmp[0][0] == '#' || tmp[0][0] == 'L' {
		return errors.New("ERROR: invalid data format, the room can't start with L, #")
	}
	if x, err := strconv.Atoi(tmp[1]); err != nil || x < 0 {
		return errors.New("ERROR: invalid data format, the coord X must be positive and number")
	}
	if y, err := strconv.Atoi(tmp[2]); err != nil || y < 0 {
		return errors.New("ERROR: invalid data format, the coord Y must be positive and number")
	}

	tmpCoord := []string{tmp[1], tmp[2]}
	if !data.isValidCoords(tmpCoord) {
		return errors.New("ERROR: invalid data format, room with such coordinates already exists")
	}

	switch s {
	case "start":
		data.StartR = tmp[0]
	case "end":
		data.EndR = tmp[0]
	default:
		data.Rooms = append(data.Rooms, tmp[0])
		data.Coords = append(data.Coords, tmpCoord)
	}
	return nil
}

func (data *Input) isValidCoords(tmpCoord []string) bool {

	for _, coord := range data.Coords {

		if coord[0] == tmpCoord[0] && coord[1] == tmpCoord[1] {
			return false
		}
	}

	return true
}
