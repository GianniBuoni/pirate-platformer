package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type SpriteGroup struct {
	IDs     map[string][]int
	Sprites map[int]Sprite
}

func (sg *SpriteGroup) GetIDs(name string) ([]int, error) {
	groupIDs, ok := sg.IDs[name]
	if !ok {
		return nil, fmt.Errorf("Group \"%s\" not found.", name)
	}
	return groupIDs, nil
}

func (sg *SpriteGroup) GetSprites(ids []int, name string) ([]Sprite, error) {
	out := []Sprite{}
	for _, id := range ids {
		s, ok := sg.Sprites[id]
		if !ok {
			return nil, DeletedError(name, id)
		}
		out = append(out, s)
	}
	return out, nil
}

func (sg *SpriteGroup) Update(name string) error {
	ids, err := sg.GetIDs(name)
	if err != nil {
		return err
	}
	sprites, err := sg.GetSprites(ids, name)
	if err != nil {
		return err
	}
	for _, s := range sprites {
		err := s.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (sg *SpriteGroup) Draw(name string) error {
	ids, err := sg.GetIDs(name)
	if err != nil {
		return err
	}
	sprites, err := sg.GetSprites(ids, name)
	if err != nil {
		return err
	}
	for _, s := range sprites {
		s.Draw(s.GetID().Src, s.GetPos())
	}
	return nil
}

func (sg *SpriteGroup) Cleanup(name string, linked ...string) error {
	ids, err := sg.GetIDs(name)
	if err != nil {
		return err
	}
	for i, id := range ids {
		s, ok := sg.Sprites[id]
		if !ok {
			return DeletedError(name, id)
		}
		if s.GetID().Kill {
			sg.IDs[name] = removeSliceIndex(i, ids)
			delete(sg.Sprites, id)
		}
	}
	// check if deleted id is in other groups
	for _, linkedName := range linked {
		linkedName, ok := sg.IDs[linkedName]
		// continue if linked group is not initialized
		if !ok {
			continue
		}
		for i, id := range linkedName {
			_, ok := sg.Sprites[id]
			// ok should fail if id is deleted by parent group
			if !ok {
				sg.IDs[name] = removeSliceIndex(i, linkedName)
			}
		}
	}
	return nil
}

func removeSliceIndex(i int, src []int) []int {
	last := len(src) - 1
	src[i] = src[last]
	return src[:last]
}
