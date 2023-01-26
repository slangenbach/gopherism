package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var count int

	for _, col := range cb[file] {
		if col {
			count += 1
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count int

	if rank >= 1 && rank <= 8 {
		for file := range cb {
			// rank ranges from 1 to 8, while indexing ranges from 0 to 7
			if cb[file][rank-1] {
				count += 1
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	files := len(cb)
	return files * files
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var files int

	for file, _ := range cb {
		files += CountInFile(cb, file)
	}

	return files

}
