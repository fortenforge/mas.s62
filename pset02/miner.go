package main

import (
	"crypto/sha256"
	"fmt"
)

// This file is for the mining code.
// Note that "targetBits" for this assignment, at least initially, is 33.
// This could change during the assignment duration!  I will post if it does.

// Mine mines a block by varying the nonce until the hash has targetBits 0s in
// the beginning.  Could take forever if targetBits is too high.
// Modifies a block in place by using a pointer receiver.
func (self *Block) Mine(targetBits uint8, kill chan bool, out chan uint64) {
	// your mining code here
	// also feel free to get rid of this method entirely if you want to
	// organize things a different way; this is just a suggestion
	blockString := self.ToString()

	nonce := uint64(0)
	NPROC := 8

	for g := 0; g < NPROC; g++ {
		go func(nonce uint64) {
			for {
				select {
				case <- kill:
					return
				default:
					attemptString := fmt.Sprintf(blockString, nonce)
					if CheckWork(attemptString) {
						out <- nonce
						fmt.Printf("%v\n", attemptString)
						return
					}
					nonce++
				}
			}
		}(nonce)
		nonce += (1 << 30)
	}
	nonce = <- out
	for i := 0; i < 3; i++ {
		kill <- true
	}
	self.Nonce = fmt.Sprintf("%x", nonce)
	return
}

// CheckWork checks if there's enough work
func CheckWork(s string) bool {
	// your checkwork code here
	// feel free to inline this or do something else.  I just did it this way
	// so I'm giving empty functions here.
	hash := sha256.Sum256([]byte(s))
	return hash[0] == 0 &&
				 hash[1] == 0 &&
				 hash[2] == 0 &&
				 hash[3] == 0 &&
				 hash[4] < 0x80
}
