package model

func exist[T string | int | currency](p T, slc ...T) bool {
	for _, s := range slc {
		if s == p {
			return true
		}
	}
	return false
}
