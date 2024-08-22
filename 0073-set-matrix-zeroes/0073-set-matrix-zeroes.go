func setZeroes(matrix [][]int)  {
    row := len(matrix)
    col := len(matrix[0])

    rows := make([]bool, row)
    cols := make([]bool, col)

    for i := 0; i< row; i++  {
        for j :=0 ; j< col;j++  {
            if(matrix[i][j] == 0 ){
                rows[i] = true
                cols[j] = true
            }
        }
    }

    for i := 0; i< row;i++ {
        for j :=0 ; j< col;j++ {
            if(rows[i] || cols[j] ){
                matrix[i][j] = 0;
            }
        }
    }
    
}

// brute solution 
// traerse evryt element and if 0 mark respective row and col as -1 and then move to others 

// better 
// traverse and figure out if there's atleast one 0 in the entire row and entire col in a n col matric for row and col each and if there's atleast one 0 then mark as true and later on mark it as 0 

// optomised 
