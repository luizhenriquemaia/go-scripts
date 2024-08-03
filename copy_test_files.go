package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func getTempDir() string {
	work_dir, err := os.Getwd()
	if err != nil {
		log.Fatal("working directory wasn't setted")
	}
	_, err = os.Stat(work_dir + "/temp/tests")
	log.Printf("work dir: %v", work_dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(work_dir+"/temp/tests", 0777)
		if err != nil {
			log.Fatalf("couldn't create temp test dir 2: %v", err)
		}
	}
	return work_dir + "/temp/tests/"
}

func copyFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("couldn't open file to copy: %v", path)
	}
	defer file.Close()
	new_file_name := strings.ReplaceAll(filepath.ToSlash(path), "/", "_")
	dest_file, err := os.Create("temp/tests/" + new_file_name)
	if err != nil {
		log.Fatalf("couldn't create copy test file: %v", path)
	}
	defer dest_file.Close()
	_, err = io.Copy(dest_file, file)
	if err != nil {
		log.Fatalf("couldn't copy test file: %v", path)
	}
}

func CopyTestFiles() {
	path_temp_dir := getTempDir()
	log.Printf("creating temp dir %v", path_temp_dir)
	log.Print("copying test files")
	name_regex, err := regexp.Compile("^.+_test.go")
	if err != nil {
		log.Fatal("fail to parse test filename regex", err)
	}
	number_founded_files := 0
	err = filepath.Walk("./internal", func(path string, info os.FileInfo, err error) error {
		if name_regex.MatchString(info.Name()) {
			number_founded_files += 1
			copyFile(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("fail to walk directory ", err)
	}
	log.Println("founded files", number_founded_files)
}

func RemoveTestFiles() {
	dir := "temp/"
	os.RemoveAll(dir)
}
