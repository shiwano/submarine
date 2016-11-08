package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

const cellSize = 5
const lightRange = 20

func main() {
	stageRootDir := getStageRootDir()
	stageDirs := getSubDirs(stageRootDir)

	var wg sync.WaitGroup
	for _, stageDir := range stageDirs {
		wg.Add(1)
		go func(stageDir string) {
			generateLightMapIfNeeded(stageDir)
			wg.Done()
		}(stageDir)
	}
	wg.Wait()
}

func generateLightMapIfNeeded(stageDir string) {
	navMeshJSONPath := path.Join(stageDir, "NavMesh.json")
	lightMapJSONPath := path.Join(stageDir, "LightMap.json")

	mesh, err := navmesh.LoadMeshFromJSONFile(navMeshJSONPath)
	if err != nil {
		log.Fatal(err)
	}

	lightMap, err := sight.LoadLightMapFromJSONFile(lightMapJSONPath)
	if err == nil &&
		lightMap.MeshVersion == mesh.Version &&
		lightMap.Helper.CellSize == cellSize &&
		lightMap.Helper.LightRange == lightRange {
		return
	}

	lightMapData, err := sight.GenerateLightMap(navmesh.New(mesh), cellSize, lightRange).ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(lightMapJSONPath, lightMapData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated", lightMapJSONPath)
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
	return path.Join(rootDir, "../../client/Assets/Art/Stages")
}
