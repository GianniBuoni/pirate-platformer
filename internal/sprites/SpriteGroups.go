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

func (sg *SpriteGroup) GetSpritesID(name string, ids []int) ([]Sprite, error) {
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

func (sg *SpriteGroup) GetSpritesName(name string) ([]Sprite, error) {
	ids, err := sg.GetIDs(name)
	if err != nil {
		return nil, err
	}
	out, err := sg.GetSpritesID(name, ids)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (sg *SpriteGroup) Update(name string) error {
	ids, err := sg.GetIDs(name)
	if err != nil {
		return err
	}
	sprites, err := sg.GetSpritesID(name, ids)
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
	sprites, err := sg.GetSpritesName(name)
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
		ids, err := sg.GetIDs(linkedName)
		if err != nil {
			// should be ok if linked group isn't initialized yet?
			continue
		}
		for i, id := range ids {
			_, ok := sg.Sprites[id]
			// ok should fail if id is deleted by parent group
			if !ok {
				sg.IDs[linkedName] = removeSliceIndex(i, ids)
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
