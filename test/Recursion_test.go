package test

import (
	"fmt"
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestCapSo(t *testing.T) {
	assert := assert.New(t)
	a := algorithms.CapSoCong(4)
	assert.Equal(10, a, "f(n) = 1+2+3+...+n must is 10")
	assert.NotEqual(9, a, "f(n) = 1+2+3+...+n must is 10")

	a = algorithms.GiaiThua(9)
	assert.Equal(362_880, a, "9! must is 362.880")
	assert.NotEqual(1, a, "9! must is 362.880")
}

func TestTailAndHeadRecursion(t *testing.T) {
	fmt.Println("de quy duoi")
	algorithms.TailRecursion(3)
	fmt.Println("==========")
	fmt.Println("de quy dau")
	algorithms.HeadRecursion(3)
}

func TestCayPhanCap(t *testing.T) {
	listEmployee := []algorithms.NhanVien{
		{Id: 1, LineManager: 0},
		{Id: 2, LineManager: 0},
		{Id: 3, LineManager: 1},
		{Id: 4, LineManager: 1},
		{Id: 5, LineManager: 2},
		{Id: 6, LineManager: 2},
		{Id: 7, LineManager: 3},
		{Id: 8, LineManager: 3},
		{Id: 9, LineManager: 7},
		{Id: 10, LineManager: 4},
	}
	lookupNV := make(map[int]algorithms.NhanVien)
	for _, emp := range listEmployee {
		lookupNV[emp.Id] = emp
	}

	lookupLineManager := make(map[int][]algorithms.NhanVien)
	for _, emp := range listEmployee {
		if emp.LineManager != 0 {
			lineManager := emp.LineManager
			if _, ok := lookupLineManager[lineManager]; !ok {
				lookupLineManager[lineManager] = []algorithms.NhanVien{emp}
				continue
			}
			lookupLineManager[lineManager] = append(lookupLineManager[lineManager], emp)
		}
	}

	// for i, emp := range listEmployee {
	// 	if _, ok := lookupLineManager[emp.Id]; ok {
	// 		listEmployee[i].ListNV = lookupLineManager[emp.Id]
	// 	}
	// }

	for i, emp := range listEmployee {
		if emp.LineManager == 0 {
			listEmployee[i].ListNV = algorithms.AddNhanVienToLineManager(listEmployee[i], lookupLineManager)
		}
	}

	algorithms.CayPhanCap(9, lookupNV)
	fmt.Println()
	algorithms.CayPhanCap_3(1, listEmployee)
	fmt.Println()
	algorithms.CayPhanCap_2(1, listEmployee[0].ListNV)
}

func TestSn(t *testing.T) {
	assert := assert.New(t)
	a := algorithms.Sn_731(4)
	assert.Equal(30, a, "S(n) = 1^2 + 2^2 + 3^2 + … + (n-1)^2 + n^2 must is 5")
}

func TestSn_732_And_736(t *testing.T) {
	assert := assert.New(t)
	a := algorithms.Sn_732(4)
	assert.Equal(float32(1.0416667), a, "wrong result")

	a = algorithms.Sn_736(5)
	assert.Equal(float32(3.55), a, "wrong result")
}

func TestSn_741(t *testing.T) {
	assert := assert.New(t)
	a := algorithms.Sn_741_2(2, 7)
	assert.Equal(254, a, "wrong result")

	a = algorithms.Sn_741_2(3, 4)
	assert.Equal(120, a, "wrong result")

	luyThua := algorithms.LuyThua(3, 10)
	assert.Equal(59049, luyThua, "wrong result")
}

func TestTowerHaNoi(t *testing.T) {
	// assert := assert.New(t)
	algorithms.TowerHaNoi(2, "A", "C", "B")
}
