package level

import (
	"math"

	. "github.com/junglemc/JungleTree/pkg/block"
)

const chunkSectionSize = 16

type ChunkSection struct {
	palette     []uint
	blockData   [chunkSize][chunkSize][chunkSize]*Block
	blockStates [chunkSize][chunkSize][chunkSize]uint
}

func (s *ChunkSection) Palette() (palette []uint) {
	if s.palette == nil {
		s.updatePalette()
	}
	return s.palette
}

func (s *ChunkSection) BitsPerBlock() byte {
	if s.palette == nil {
		s.updatePalette()
	}
	// TODO: Check logic for indirect palette bits per block
	return byte(math.Ceil(math.Log2(float64(len(s.palette)))))
}

func (s *ChunkSection) BlockCount() (count int32) {
	for y := 0; y < chunkSectionSize; y++ {
		for z := 0; z < chunkSectionSize; z++ {
			for x := 0; x < chunkSectionSize; x++ {
				if s.blockData[y][z][x] != nil && s.blockData[y][z][x].Name != "air" {
					count++
				}
			}
		}
	}
	return
}

func (s *ChunkSection) BlockAt(x, y, z int32) *Block {
	if s.blockData[y][z][x] == nil {
		s.blockData[y][z][z] = Empty()
	}
	return s.blockData[y][z][x]
}

func (s *ChunkSection) SetBlock(x, y, z int32, block *Block) {
	s.blockData[y][z][x] = block
}

func (s *ChunkSection) HighestBlockAt(x, z int) (y int32, ok bool) {
	for y := len(s.blockData) - 1; y >= 0; y-- {
		if s.blockData[y][z][x].Id != 0 {
			return int32(y), true
		}
	}
	return 0, false
}

func (s *ChunkSection) updatePalette() {
	exists := struct{}{}
	paletteSet := make(map[uint]struct{})

	for y := 0; y < chunkSectionSize; y++ {
		for z := 0; z < chunkSectionSize; z++ {
			for x := 0; x < chunkSectionSize; x++ {
				if s.blockData[y][z][x] == nil {
					continue
				}

				if s.blockStates[y][z][x] == 0 {
					paletteSet[s.blockData[y][z][x].DefaultStateId] = exists
				}
				paletteSet[s.blockStates[y][z][x]] = exists
			}
		}
	}

	s.palette = make([]uint, 0, len(paletteSet))
	for k := range paletteSet {
		s.palette = append(s.palette, k)
	}
	return
}
