package main

func Answer(nums []uint) uint {
	var count uint
	count = 0
	for _, v := range nums {

    if v == 0 {
      count++
      return count
    }

    count++
	}
	return count
}
