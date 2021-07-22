package renderer

import (
	"fmt"
	"log"

	"github.com/mackstann/exopolis/city/domain"

	"github.com/muesli/termenv"
)

func Render(city domain.City, n *domain.JobTransportNetwork) {
	if termenv.ColorProfile() != termenv.TrueColor {
		log.Fatalf("not enough color! %v, want %v", termenv.ColorProfile(), termenv.TrueColor)
	}

	for _, row := range textualize(city, n) {
		fmt.Println(row)
	}
}

func textualize(city domain.City, n *domain.JobTransportNetwork) []string {
	rows := make([]string, 0, len(city))
	for y, row := range city {
		rowOutput := ""
		for x, cell := range row {
			nCell := n.Grid[y][x]
			temp255 := int(nCell.Temperature * 255)
			intensity := fmt.Sprintf("%02x", temp255)
			c := "."
			color := ""
			switch cell.Typ {
			case domain.House:
				c = "■"
				color = intensity + "0000"
			case domain.Farm:
				c = "▤"
				color = "00" + intensity + "00"
			case domain.Road:
				c = "▪"
				color = intensity + intensity + intensity
			case domain.PowerPlant:
				c = "p"
				color = intensity + intensity + "00"
			case domain.Dirt:
				c = "░"
				color = "0000" + intensity
			}
			p := termenv.ColorProfile()
			if intensity != "" && c != " " && len(color) == 6 {
				rowOutput += termenv.String(c).Foreground(p.Color("#" + color)).String()
			} else {
				rowOutput += c
			}
		}
		rows = append(rows, rowOutput)
	}
	return rows
}
