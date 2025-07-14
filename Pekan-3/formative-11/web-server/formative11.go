package formative11

import "math"

func LuasAlas(jariJari int, tinggi int) float64 {
	return math.Pi * float64(jariJari*jariJari) * float64(tinggi)
}

func KelilingAlas(jariJari int) float64 {
	return math.Pi * float64(jariJari*jariJari)
}

func VolumeTabung(jariJari int) float64 {
	return 2 * math.Pi * float64(jariJari)

}
