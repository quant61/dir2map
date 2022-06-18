package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
	"syscall"
	"unicode"
)

func IsPrintable(s string) bool {
	for _, r := range []rune(s) {
		if !unicode.IsPrint(r) && r != '\n' && r != '\t' {
			//fmt.Println("wrong char", r)
			return false
		}
	}
	return true
}

func main() {
	dir := "."
	if len(os.Args) > 1{
		dir = os.Args[1]
	}
	ff, err := ioutil.ReadDir(dir)
	if err != nil{
		panic(err)
	}
	for _, f := range ff {
		b, err := ioutil.ReadFile(path.Join(dir, f.Name()))
		if err != nil{
			switch {
			case errors.Is(err, syscall.EISDIR):
				fmt.Println(f.Name(), "<Directory>")
			case errors.Is(err, syscall.EACCES):
				fmt.Println(f.Name(), "[Access Denied]")
			default:
				var errno syscall.Errno
				is_errno := errors.As(err, &errno)
				if is_errno {
					fmt.Printf("%s Errno(%d): %s\n", f.Name(), errno, errno)
				} else {
					fmt.Println(f.Name(), reflect.TypeOf(err),  err)
					for err != nil {
						err = errors.Unwrap(err)
						fmt.Println("...", reflect.TypeOf(err), err)
					}
				}
			}
			continue
		}
		s := string(b)
		if(!IsPrintable(s)){
			if len(b) > 128 {
				fmt.Println(f.Name(), "=0x" + hex.EncodeToString(b[:128]), "...", len(b))
			} else {
				fmt.Println(f.Name(), "=0x" + hex.EncodeToString(b))
			}
		} else {
			s = strings.Trim(s, "\n")
			fmt.Println(f.Name(), "=", s)
		}
	}
}
