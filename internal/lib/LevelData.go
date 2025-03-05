package lib

import (
	"errors"
	"fmt"
)

type LevelData struct {
	MapProps MapProps  `json:"properties"`
	Layers   []Layer   `json:"layers"`
	Width    int       `json:"width"`
	TileRefs []TileRef `json:"tilesets"`
}

type TileParams struct {
	Data      []int
	GIDRanges map[GIDRange]string
	Columns   int
}

func (ld *LevelData) MapGIDRanges(a *Assets) (map[GIDRange]string, error) {
	out := map[GIDRange]string{}
	for _, tileRef := range ld.TileRefs {
		tilesetName := GetAssetKey(tileRef.Source)
		tileset, err := a.GetTileset(tilesetName)
		if err != nil {
			return nil, err
		}
		key := GIDRange{
			FirstGID: tileRef.FirstGID,
			LastGID:  tileRef.FirstGID + tileset.Count - 1,
		}
		out[key] = tilesetName
	}
	return out, nil
}

func getGIDSource(
	gid int, gidMap map[GIDRange]string,
) (tilesetName string, firstGID int, err error) {
	for k, v := range gidMap {
		if k.FirstGID <= gid && k.LastGID >= gid {
			tilesetName = v
			firstGID = k.FirstGID
		}
	}
	if tilesetName == "" {
		return "", 0, fmt.Errorf("gid: %d\n not found in tileset refs.", gid)
	}

	return tilesetName, firstGID, nil
}

func getGIDImgPos(
	gid, firstGID int, tilesetName string, a *Assets,
) (x, y float32, err error) {
	tileset, err := a.GetTileset(tilesetName)
	if err != nil {
		return 0, 0, err
	}

	tileID := gid - firstGID
	if tileID > tileset.Count {
		fmt.Printf(
			"index: %d out of range of tileset '%s'.\n", tileID, tilesetName,
		)
		return 0, 0,
			errors.New("Check if there are rotation flags on tile gid")
	}
	x = float32(tileID%tileset.Columns) * TileSize
	y = float32(tileID/tileset.Columns) * TileSize
	return x, y, nil
}

func ParseTileImage(params TileParams, index int, a *Assets) (Tile, error) {
	var (
		tile     Tile
		err      error
		firstGID int
	)
	// get tile image
	gid := params.Data[index]
	tile.Image, firstGID, err = getGIDSource(gid, params.GIDRanges)
	if err != nil {
		return Tile{}, err
	}
	// get inner posision
	tile.ImgX, tile.ImgY, err = getGIDImgPos(gid, firstGID, tile.Image, a)
	if err != nil {
		return Tile{}, err
	}
	// get map pos
	tile.X = float32(index%params.Columns) * TileSize
	tile.Y = float32(index/params.Columns) * TileSize
	return tile, nil
}
