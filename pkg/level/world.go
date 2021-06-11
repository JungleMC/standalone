package level

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/errors"

	"github.com/junglemc/JungleTree/internal/storage"
	. "github.com/junglemc/JungleTree/pkg/block"
	. "github.com/junglemc/JungleTree/pkg/util"
)

type World struct {
	Name   Identifier
	Seed   uint64
	Height uint
}

func ListWorlds() []Identifier {
	worlds := make([]Identifier, 0, 0)
	err := storage.Get("jungletree:worlds", &worlds, nil)
	if err != nil {
		panic(err)
	}
	return worlds
}

func NewWorld(name string, seed uint64, height uint) *World {
	id := Identifier(fmt.Sprintf("world:%s", name))

	worlds := make([]Identifier, 0, 0)
	err := storage.Get("jungletree:worlds", &worlds, nil)
	if err != nil && err != errors.ErrNotFound {
		panic(err)
	}

	worlds = append(worlds, id)
	err = storage.Put("jungletree:worlds", worlds, nil)
	if err != nil {
		panic(err)
	}

	result := World{
		Name:   id,
		Seed:   seed,
		Height: height,
	}

	err = storage.Put(result.worldKey(), result, nil)
	if err != nil {
		panic(err)
	}
	return &result
}

func GetWorld(name Identifier) *World {
	id := Identifier(fmt.Sprintf("world:%s", name.Name()))
	ok, err := storage.Has(id, nil)
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	result := World{}
	err = storage.Get(id, &result, nil)
	if err != nil {
		panic(err)
	}
	return &result
}

func DefaultWorld() *World {
	var id Identifier
	err := storage.Get("jungletree:default_world", &id, nil)
	if err != nil {
		panic(err)
	}
	return GetWorld(id)
}

func (w *World) ChunkAt(x int32, z int32) *Chunk {
	index := uint64(x)<<32 | uint64(z)

	chunk := w.chunkAt(index)
	if chunk == nil {
		chunk = &Chunk{
			World:     w,
			X:         byte(x),
			Z:         byte(z),
			sections:  make([]ChunkSection, w.Height/chunkSectionSize),
			heightMap: [256]int32{},
			biomes:    [1024]int32{},
		}
		if err := chunk.Save(); err != nil {
			panic(err)
		}
	}
	return chunk
}

func (w *World) worldKey() Identifier {
	return Identifier(fmt.Sprintf("world:%s", w.Name.Name()))
}

func (w *World) chunkKey(index uint64) Identifier {
	return Identifier(fmt.Sprintf("jungletree:%s_%d", w.Name.Name(), index))
}

func (w *World) chunkAt(index uint64) *Chunk {
	var chunk Chunk
	err := storage.Get(w.chunkKey(index), &chunk, nil)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return &chunk
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
