package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: ./lightmap cellSize lightRange")
	}
	cellSize, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	lightRange, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	stageRootDir := getStageRootDir()
	stageDirs := getSubDirs(stageRootDir)

	var wg sync.WaitGroup
	for _, stageDir := range stageDirs {
		wg.Add(1)
		go func(stageDir string) {
			generateLightMapIfNeeded(cellSize, lightRange, stageDir)
			wg.Done()
		}(stageDir)
	}
	wg.Wait()
}

func generateLightMapIfNeeded(cellSize, lightRange float64, stageDir string) {
	navMeshFilePath := path.Join(stageDir, "mesh.json")
	lightMapFilePath := path.Join(stageDir, "lightmap.mpac")

	mesh, err := navmesh.LoadMeshFromJSONFile(navMeshFilePath)
	if err != nil {
		log.Fatal(err)
	}

	lightMap, err := sight.LoadLightMapFromMessagePackFile(lightMapFilePath)
	if err == nil &&
		lightMap.MeshVersion == mesh.Version &&
		lightMap.Helper.CellSize == cellSize &&
		lightMap.Helper.LightRange == lightRange {
		return
	}

	lightMapData, err := sight.GenerateLightMap(navmesh.New(mesh), cellSize, lightRange).ToMessagePack()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(lightMapFilePath, lightMapData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated", lightMapFilePath)
}

func getSubDirs(dir string) (subDirs []string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			subDir := path.Join(dir, fileInfo.Name())
			subDirs = append(subDirs, subDir)
		}
	}
	return
}

func getStageRootDir() string {
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(rootDir, "../../server/battle/assets/stages")
}
