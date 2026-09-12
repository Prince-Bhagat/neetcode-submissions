func checkSubGroup(board [][]byte)bool{
	for i := 0; i < 9; i= i+ 3 {
		for j := 0; j < 9; j=j+3 {
			tempArray := [9]bool{}
			
			for a := 0; a < 3; a++ {
				
				for b := 0; b < 3; b++ {
					row := a+j
					column := b+ i
					num := int(board[row][column]) - '0' -1
					
					if num < 0{
						continue
					}
					if (tempArray[num]){
						return false
					}else{
						tempArray[num] = true
					}

				}

			}
		}
	}
	return true

}

func isValidSudoku(board [][]byte) bool {
	for i := 0 ;i < 9;i++{
		boolTableRow := [9]bool{}
		for j:= 0;j< 9;j++{
			if board[i][j] == '.' {
				continue
			}
			num:= int(board[i][j]) - '0' -1
			if boolTableRow[num] {
				return false
			}else{
				boolTableRow[num] = true
			}
		}
	}

	for j := 0 ;j < 9;j++{
		boolTableColumn := [9]bool{}
		for i:= 0;i< 9;i++{
			if board[i][j] == '.' {
				continue
			}
			num:= int(board[i][j]) - '0' -1
			if boolTableColumn[num] {
				return false
			}else{
				boolTableColumn[num] = true
			}
		}
	}
	
	return checkSubGroup(board)
}


