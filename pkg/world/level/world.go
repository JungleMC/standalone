package level

import (
	. "github.com/junglemc/JungleTree/pkg/block"
)

type World struct {
	Name   string
	Seed   uint64
	Height uint
	Chunks map[uint64]*Chunk
}

func (w *World) ChunkAt(x int32, z int32) *Chunk {
	index := uint64(x)<<32 | uint64(z)

	chunk := w.Chunks[index]
	if chunk == nil {
		chunk = &Chunk{
			World:     w,
			X:         byte(x),
			Z:         byte(z),
			sections:  make([]ChunkSection, w.Height/chunkSectionSize),
			heightMap: [256]int32{},
			biomes:    [1024]int32{},
		}
		w.Chunks[index] = chunk
	}
	return chunk
}

func (w *World) BlockAt(x, y, z int32) *Block {
	modX := x % chunkSize
	modZ := z % chunkSize

	return w.ChunkAt(x-modX, z-modZ).BlockAt(modX, y, modZ)
}

func (w *World) SetBlock(x, y, z int32, block *Block) {
	modX := x % chunkSize
	modZ := z % chunkSize

	w.ChunkAt(x-modX, z-modZ).SetBlock(modX, y, modZ, block)
}
