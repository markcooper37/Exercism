package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank 

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    _, rankExists := cb[rank]
    if !rankExists {
        return 0
    }
	count := 0
    for _, squareOccupied := range cb[rank] {
        if squareOccupied {
            count++
        }
    }
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    if file < 1 || file > 8 {
        return 0
    }
    count := 0
	for _, rank := range cb {
        if rank[file - 1] {
            count++
        }
    }
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for _, rank := range cb {
        for range rank {
            count++
        }
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
    for _, rank := range cb {
        for _, squareOccupied := range rank {
            if squareOccupied {
                count++
            }
        }
    }
	return count
}
