package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	maxProduct := Product{Product: fmin*fmin - 1, Factorizations: [][2]int{}}
	minProduct := Product{Product: fmax*fmax + 1, Factorizations: [][2]int{}}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			stringProduct := strconv.Itoa(i * j)
			if isPalindrome(stringProduct) {
				if product > maxProduct.Product {
					maxProduct.Product = product
					maxProduct.Factorizations = [][2]int{{i, j}}
				} else if product == maxProduct.Product {
					maxProduct.Factorizations = append(maxProduct.Factorizations, [2]int{i,j})
				}
				if product < minProduct.Product {
					minProduct.Product = product
					minProduct.Factorizations = [][2]int{{i, j}}
				} else if product == minProduct.Product {
					minProduct.Factorizations = append(minProduct.Factorizations, [2]int{i,j})
				}
			}
		}
	}
	if minProduct.Product == fmax*fmax + 1 {
		return Product{}, Product{}, errors.New("no palindromes")
	}
	return minProduct, maxProduct, nil
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
