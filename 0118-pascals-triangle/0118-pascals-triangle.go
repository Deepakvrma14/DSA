func generate(nums int) [][]int {
    if(nums == 0) {
        return [][]int{}
    }
    if(nums == 1){
        return [][]int{{1}}
    }

    prevRows := generate(nums - 1)

    newRow := make([]int, nums)
    newRow[0] = 1
    newRow[nums-1] = 1

    prevRow := prevRows[nums-2]

    for i := 1; i < nums-1; i++ {
        newRow[i] = prevRow[i-1] + prevRow[i]
    }
    // prevRow.append(newRow)
    prevRows = append(prevRows, newRow)
    return prevRows
}


// find value at pos r and c can be done by formula of combination .. r-1 C c-1 and 