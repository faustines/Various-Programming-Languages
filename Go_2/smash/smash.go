/*
ECS 140A Programming Hw5
Faustine Yiu 913062973
Source: https://golang.org, ECS 140A Go Lecture, StackOverflow
*/

package smash

import (
	"io"
	"bufio"
	"sync"
)

type word string

// Smash takes as input an io.Reader and a smasher function,
// and returns
func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint{
	m := make(map[uint32]uint)
	// TODO: Incomplete!
	var wg sync.WaitGroup
	var a uint32
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
	    b := word(scan.Text())
	    a = smasher(b)
	    wg.Add(1)
	    go func(key uint32,mp map[uint32]uint,wg *sync.WaitGroup){
	        mp[key] = mp[key] + 1
	        wg.Done()
	    }(a,m,&wg)
	    wg.Wait()
    }
	return m
}
