package main

import (
	"fmt"
	"testing"
)

var (
	orders []order
	visits []visit
)

func TestMain(t *testing.M) {
	for i := 0; i < 5; i++ {
		seedDataForPartner(i, &visits, &orders)
	}
	fmt.Printf("Benchmarking agains %d visits total and %d orders total\n", len(visits), len(orders))

	t.Run()
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPartnerCommissions(visits, orders)
	}
}

func BenchmarkMoreOptimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moreOptimalCommissions(visits, orders)
	}
}
