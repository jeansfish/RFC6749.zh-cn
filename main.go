package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var (
	_   = fmt.Printf
	dir = `D:\guotie\code\oauth2\RFC6749.zh-cn\`
	LF  = []byte("\r\n")
)

type ConvFile struct {
	ConvName string
	FileName string
}

type ConvFiles []ConvFile

func (cf ConvFiles) Len() int { return len(cf) }

func (cf ConvFiles) Swap(i, j int) { cf[i], cf[j] = cf[j], cf[i] }

func (cf ConvFiles) Less(i, j int) bool { return cf[i].ConvName < cf[j].ConvName }

// 把文件名转换为正确的排序格式
// 例如按照普通字母排序的话, 结果是这样的: 10.1 10.10 10.11 10.12 10.2 10.3
// 这个函数把10.1 10.10 10.11 10.12 10.2 10.3先转换为010.001 010.010 010.011 010.012 010.002 010.003
// 然后再按字母排序, 得到正确的顺序
func conv(files []string) []ConvFile {
	var cf []ConvFile = make([]ConvFile, len(files))
	for idx, file := range files {
		segs := strings.Split(strings.ToLower(file), ".md")
		if len(segs) != 2 {
			panic(fmt.Sprintf("file name format error: %s!\n", file))
		}

		var conv string
		nums := strings.Split(segs[0], ".")
		for _, num := range nums {
			no, err := strconv.Atoi(num)
			if err != nil {
				panic(fmt.Sprintf("file name %s format error: chapter is not number\n", segs[0]))
			}
			conv += fmt.Sprintf("%03d.", no)
		}
		cf[idx] = ConvFile{conv, file}
	}

	return cf
}

func sortFiles(files []string) []string {
	var (
		conv_files []ConvFile
		res        []string = make([]string, len(files))
	)
	conv_files = conv(files)
	sort.Sort(ConvFiles(conv_files))
	for i, f := range conv_files {
		res[i] = f.FileName
	}

	return res
}

func catSectionFiles(buf *bytes.Buffer, dir string) error {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error found: path: %s error: %s\n", path, err.Error())
			return err
		}
		if info.IsDir() && path != dir {
			return filepath.SkipDir
		}
		name := strings.ToLower(info.Name())
		if strings.HasSuffix(name, ".md") {
			files = append(files, info.Name())
		}

		return nil
	})

	if err != nil && err != filepath.SkipDir {
		panic(err)
	}

	//fmt.Printf("Section Dir: %s files: %v\n", dir, files)
	// 文件排序
	files = sortFiles(files)
	//fmt.Printf("Section Dir: %s sorted files: %v\n", dir, files)

	// 逐个读出文件的内容, 合并
	for _, fn := range files {
		file := path.Join(dir, fn)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(fmt.Sprintf("read file %s failed: %s\n", file, err.Error()))
		}
		_, err = buf.Write(content)
		if err != nil {
			panic("write file content to buffer failed: " + err.Error())
		}
		_, err = buf.Write(LF)
		if err != nil {
			panic("write LF to buffer failed: " + err.Error())
		}
	}

	return nil
}

func catSectionDir(dir string) {
	var (
		contents []byte
		buf      *bytes.Buffer
		sections []string
	)

	buf = bytes.NewBuffer(contents)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error found: path: %s error: %s\n", path, err.Error())
			return err
		}
		if info.IsDir() && path != dir {
			if strings.HasPrefix(info.Name(), "Section") {
				sections = append(sections, info.Name())
			}
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(sections)
	for _, sec := range sections {
		catSectionFiles(buf, dir+sec)
	}
	err = ioutil.WriteFile(path.Join(dir, "all-in-one.md"), buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	catSectionDir(dir)
}
