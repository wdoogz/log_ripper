package scrape

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Scrape scrapes
func Scrape() {
	dir := "/home/pi/logs_toscrape"
	//parDir, err := ioutil.ReadDir(dir)

	hopDown(dir)

}

func hopDown(parentDir string) {
	tree := filepath.Walk(parentDir, treeDir)
	if tree != nil {
		fmt.Println(tree)
	}
}

func treeDir(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !info.IsDir() {
		file, _ := ioutil.ReadFile(path)
		fmt.Println(string(file))
	}
	fmt.Println(path, info.IsDir())
	return nil
}
