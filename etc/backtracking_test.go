package etc

import "testing"

func Test8Hau(t *testing.T) {
	KhoiTao()
	BacktrackingQueen(0)
}

func TestLietKeNhiPhan(t *testing.T) {
	LietKeNhiPhan(0, 8)
}

func TestChinhHopKhongLap(t *testing.T) {
	for i := range choose {
		choose[i] = true
	}
	LietKeChinhHopKoLap(0, 3, 2)
}

func TestPhanTichSo(t *testing.T) {
	KhoiTaoPhanTichSo()
	PhanTichSo(1, 6)
}
