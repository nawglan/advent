package main

import (
	"fmt"
	p "advent/util/parse"
)

var (
	filter map[int]bool
)

type Image struct {
	width  int
	length int
	pixels map[int]bool
	bg     bool
}

func (i Image) dump() {
	for y := 0; y < i.length; y++ {
		for x := 0; x < i.width; x++ {
			if i.pixels[i.coords2pos(x, y)] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (i Image) enhancePixel(pos int) bool {
	posX := pos % i.width
	posY := pos / i.width

	inf_value := "0"
	if i.bg {
		inf_value = "1"
	}

	binary := ""

	var nY, nX int

	// north-west
	if posX > 0 && posY > 0 {
		nY = posY - 1
		nX = posX - 1
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else { // it would be off the map
		binary += inf_value
	}

	// north
	if posY > 0 {
		nY = posY - 1
		nX = posX
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else {
		binary += inf_value
	}

	// north-east
	if posX < i.width-1 && posY > 0 {
		nY = posY - 1
		nX = posX + 1
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else {
		binary += inf_value
	}

	// west
	if posX > 0 {
		nY = posY
		nX = posX - 1
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else {
		binary += inf_value
	}

	// self
	nY = posY
	nX = posX
	if i.pixels[i.coords2pos(nX, nY)] {
		binary += "1"
	} else {
		binary += "0"
	}

	// east
	if posX < i.width-1 {
		nY = posY
		nX = posX + 1
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else {
		binary += inf_value
	}

	// south-west
	if posX > 0 && posY < i.length-1 {
		nY = posY + 1
		nX = posX - 1
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else {
		binary += inf_value
	}

	// south
	if posY < i.length-1 {
		nY = posY + 1
		nX = posX
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else {
		binary += inf_value
	}

	// south-east
	if posX < i.width-1 && posY < i.length-1 {
		nY = posY + 1
		nX = posX + 1
		if i.pixels[i.coords2pos(nX, nY)] {
			binary += "1"
		} else {
			binary += "0"
		}
	} else {
		binary += inf_value
	}

	idx := p.Bin2dec(binary)
	result := filter[int(idx)]

	return result
}

func (i Image) coords2pos(x, y int) int {
	return x + y*i.width
}

func (i Image) addPixel(x, y int, val bool) {
	i.pixels[i.coords2pos(x, y)] = val
}

func (i Image) resize() Image {
	image := Image{
		width:  i.width + 2,
		length: i.length + 2,
		pixels: map[int]bool{},
		bg:     i.bg,
	}

	for idx := range i.pixels {
		posX := 1 + idx%i.width
		posY := 1 + idx/i.width
		image.addPixel(posX, posY, true)
	}

	for y := 0; y < image.length; y++ {
		image.addPixel(0, y, true)
		image.addPixel(image.width-1, y, true)
	}
	for x := 1; x < image.width-1; x++ {
		image.addPixel(x, 0, true)
		image.addPixel(x, image.length-1, true)
	}

	return image
}

func (i Image) enhance() Image {
	image := i.resize()
	new_image := Image{
		width:  image.width,
		length: image.length,
		pixels: map[int]bool{},
		bg:     image.bg,
	}

	for idx := range i.pixels {
		x := idx % i.width
		y := idx / i.width
		new_image.pixels[new_image.coords2pos(x+1, y+1)] = image.enhancePixel(idx)
	}

	return new_image
}

func day20(puzzle_data []string) {

	filter := map[int]bool{}
	image := Image{
		width:  len(puzzle_data[2]),
		length: len(puzzle_data),
		pixels: map[int]bool{},
	}

	// parse the filter
	for i, char := range puzzle_data[0] {
		if char == '#' {
			filter[i] = true
		}
	}

	if filter[0] {
		image.bg = true
	}

	// only keep track of lit pixels
	for y := 2; y < len(puzzle_data); y++ {
		dataVal := puzzle_data[y]
		for x, char := range dataVal {
			if char == '#' {
				image.addPixel(x, y-2, true)
			}
		}
	}

	part_1 := image.enhance().enhance()
	fmt.Printf("Day 20 (part 1): There are %d lit pixels.\n", len(part_1.pixels))

	part_2 := Image{}
	for i := 0; i < 50; i++ {
		part_2 = part_2.enhance()
	}
	fmt.Printf("Day 20 (part 2): There are %d lit pixels.\n", len(part_2.pixels))
}
