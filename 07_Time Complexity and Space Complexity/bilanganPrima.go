package main

import (
	"math/big"
	"math/rand"
	"time"
)

// Fungsi untuk menghitung (a^b) % m
func powerMod(a, b, m *big.Int) *big.Int {
	res := big.NewInt(1)
	for b.Cmp(big.NewInt(0)) > 0 {
		if b.Bit(0) == 1 {
			res.Mul(res, a)
			res.Mod(res, m)
		}
		a.Mul(a, a)
		a.Mod(a, m)
		b.Rsh(b, 1)
	}
	return res
}

// Fungsi untuk menentukan apakah sebuah bilangan termasuk bilangan prima atau tidak
func isPrime(n *big.Int, k int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < k; i++ {
		// Pilih bilangan acak a dalam rentang (2, n-2)
		a := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(n, big.NewInt(4)))
		a.Add(a, big.NewInt(2))

		// Hitung a^(n-1) % n
		x := powerMod(a, new(big.Int).Sub(n, big.NewInt(1)), n)
		if x.Cmp(big.NewInt(1)) != 0 {
			return false
		}
	}

	return true
}

func main() {
	n := big.NewInt(1000000000000000007) // Contoh: uji keprimaan untuk bilangan 1000000000000000007
	k := 10                              // Jumlah iterasi uji keprimaan

	if isPrime(n, k) {
		println(n, "adalah bilangan prima")
	} else {
		println(n, "bukan bilangan prima")
	}
}
