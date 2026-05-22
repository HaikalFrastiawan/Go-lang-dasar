package golang

// NumberConstraint menggunakan tilde (~) untuk type approximation.
// Constraint ini menerima int, float64, DAN semua custom type yang berbasis tipe tersebut.
type NumberConstraint interface {
	~int | ~float64
}

// Double akan mengalikan nilai input dengan 2.
// Return type-nya sama dengan input type-nya (T).
func Double[T NumberConstraint](val T) T {
	return val * 2
}