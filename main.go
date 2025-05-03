package main

import (
	"fmt"
	"math/rand/v2"
)

type dataSeed struct {
	visits []*visit
	orders []*order
}

type visit struct {
	visitId   string
	partnerId int
}

type order struct {
	visitId string
	amount  float64
}

type partnerVisit struct {
	ordersTotal float64
	visitsCount int
}

// getPartnerCommissions the original solution I submitted during the interview
func getPartnerCommissions(visits []visit, orders []order) map[int]float64 {
	partners := make(map[int]partnerVisit)

	for _, v := range visits {
		partner, ok := partners[v.partnerId]
		if !ok {
			partner = partnerVisit{}
			partners[v.partnerId] = partner
		}
		partner.visitsCount++

		for _, o := range orders {
			if o.visitId == v.visitId {
				partner.ordersTotal += o.amount
			}
			partners[v.partnerId] = partner
		}
	}

	coms := make(map[int]float64)
	for k, v := range partners {
		if v.visitsCount >= 5000 {
			coms[k] = v.ordersTotal * 0.1
		}
	}

	return coms
}

// moreOptimalCommissions is significantly faster by only looping over each slice a single time.
// I first iterate over the orders slice and store the order->visit joined amount
// in a map for a faster lookup.
//
// then I iterate over the visits slice and store a structure of the partner containing total amount
// and visits count using the constant time lookup of ther ordermap to lookup by visitId
// while in this loop I check if the current partner I'm looking at has me the 5K critieria
// if so, then that partnerid gets added to the comms structs to be returned.
//
//	finally, I calculate the 10% commision for all partners in the comms map
func moreOptimalCommissions(visits []visit, orders []order) map[int]float64 {
	comms := make(map[int]float64)
	pv := make(map[int]partnerVisit)
	ordermap := make(map[string]float64)

	// convert the orders list to a map for O(1) lookup time
	// tradeoff, this will increase the space conplexity given
	// the storage of the new data structure
	for _, o := range orders {
		ordermap[o.visitId] = o.amount
	}

	// now iterate over all visists and update the
	// partner data for that visit, getting the order amount
	// from the ordermap
	for _, v := range visits {
		p, ok := pv[v.partnerId]
		if !ok {
			p = partnerVisit{}
			pv[v.partnerId] = p
		}

		p.visitsCount++
		p.ordersTotal += ordermap[v.visitId]
		_, exists := comms[v.partnerId]

		// add the partner if they've met the min criteria
		// and isn't already in the comms map
		if p.visitsCount >= 5000 && !exists {
			comms[v.partnerId] = p.ordersTotal
		}

		pv[v.partnerId] = p
	}

	// calculate the commissions
	for id := range comms {
		partner := pv[id]
		comms[id] = partner.ordersTotal * 0.1
	}

	return comms
}

func seedDataForPartner(pId int, visists *[]visit, orders *[]order) {
	for i := 0; i < 25000; i++ {
		visitId := fmt.Sprintf("%d_%d", pId, i)
		v := visit{visitId: visitId, partnerId: pId}
		*visists = append(*visists, v)

		if rand.IntN(1000) < 10 {
			price := float64(rand.IntN(100)) + rand.Float64()
			o := order{visitId: visitId, amount: price}
			*orders = append(*orders, o)
		}
	}
}

func main() {

	orders := make([]order, 0)
	visits := make([]visit, 0)

	for i := 0; i < 5; i++ {
		seedDataForPartner(i, &visits, &orders)
	}

	orig := getPartnerCommissions(visits, orders)
	fmt.Printf("%+v\n", orig)

	coms := moreOptimalCommissions(visits, orders)
	fmt.Printf("%+v\n", coms)
}
