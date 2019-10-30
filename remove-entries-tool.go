package main

import (
	"errors"
	"os"
	"flag"
	"bufio"
)

var(
	ReadFile string
	TargetFile string
	ErrParams = errors.New("A target file must be selected")
	ErrFileNotFound = errors.New("File not found")
)

func main() {
	if err := Init(); err == ErrParams{
		print(err)
		os.Exit(2)
	} else if err != nil {
		print(err)
		os.Exit(1)
	}
}

func Init() error{
	flag.StringVar(&ReadFile,"e","","File in which parameters will be read")
	flag.StringVar(&TargetFile,"t","", "File in which entries will be removed")
	flag.Parse()
	if err := RemoveEntriesFromTable(ReadFile,TargetFile); err != nil{
		return err
	}
	println(ReadFile)
	println(TargetFile)
	return nil
}

func RemoveEntriesFromTable(e string,t string) error{
	p, err := ReadEntry(e)
	if err != nil {
		return err
	}
	s, err := ReadEntry(t)
	if err != nil {
		return err
	}
	for i := 0; i < len(p); i++{
		for j := 0; j < len(s); j++{
			if p[i] == s[j]{
				s = append(s[:j], s[j+1:]...)
			}
		}
	}

	r := CreateResultTable(s)
	return r
}

func ReadEntry(f string) ([]string, error){
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entryArgs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entryArgs = append(entryArgs, scanner.Text())
	}
	return entryArgs, scanner.Err()
}

func CreateResultTable(r []string) error{
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	for _, line := range r {
		println(line)
		w.WriteString(line+ "\n")
	}
	w.Flush()
	file.Close()

	return nil
}