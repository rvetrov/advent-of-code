package day09

import (
	"fmt"
	"strings"

	"adventofcode.com/internal/utils"
)

type Node struct {
	id   int
	len  int
	free bool
}

type DiskMap []Node

func (d DiskMap) line() string {
	var parts []string
	for _, n := range d {
		ch := "."
		if !n.free {
			ch = fmt.Sprint(n.id)
		}
		parts = append(parts, strings.Repeat(ch, n.len))
	}
	return strings.Join(parts, "")
}

func parseDiskMap(line string, unpackFiles bool) DiskMap {
	var res DiskMap
	var id int
	isFile := true
	for i := 0; i < len(line); i++ {
		rep := int(line[i]) - int('0')
		if isFile {
			if unpackFiles {
				for ; rep > 0; rep-- {
					res = append(res, Node{id, 1, false})
				}
			} else {
				res = append(res, Node{id, rep, false})
			}
			id++
		} else {
			if unpackFiles {
				for ; rep > 0; rep-- {
					res = append(res, Node{0, 1, true})
				}
			} else {
				res = append(res, Node{0, rep, true})
			}
		}
		isFile = !isFile
	}
	return res
}

func compactV1(disk DiskMap) DiskMap {
	left, right := 0, len(disk)-1
	for left < right {
		for left < right && !disk[left].free {
			left++
		}
		for left < right && disk[right].free {
			right--
		}
		if left < right {
			disk[left], disk[right] = disk[right], disk[left]
		}
	}
	return disk
}

func compactV2(disk DiskMap) DiskMap {
	left, right := 0, len(disk)-1
	for left <= right {
		if disk[right].free {
			right--
			continue
		}
		if !disk[left].free {
			left++
			continue
		}

		found := false
		for i := left; i < right; i++ {
			if disk[i].free && disk[i+1].free {
				disk[i+1].len += disk[i].len
				disk[i].len = 0
				continue
			}

			if disk[i].free && disk[i].len >= disk[right].len {
				rest := disk[i].len - disk[right].len
				disk[i] = disk[right]
				disk[right].free = true

				if rest > 0 {
					var tmp []Node
					tmp = append(tmp, disk[:i+1]...)
					tmp = append(tmp, Node{0, rest, true})
					tmp = append(tmp, disk[i+1:]...)
					disk = tmp
				} else {
					right--
				}
				found = true
				break
			}
		}
		if !found {
			right--
		}
	}
	return disk
}

func checksum(disk DiskMap) int {
	ret := 0
	i := 0
	for _, node := range disk {
		if !node.free {
			for range node.len {
				ret += node.id * i
				i++
			}
		} else {
			i += node.len
		}
	}
	return ret
}

func SolveV1(input string) int {
	line := utils.NonEmptyLines(input)[0]
	disk := parseDiskMap(line, true)
	disk = compactV1(disk)
	return checksum(disk)
}

func SolveV2(input string) int {
	line := utils.NonEmptyLines(input)[0]
	disk := parseDiskMap(line, false)
	disk = compactV2(disk)
	return checksum(disk)
}
