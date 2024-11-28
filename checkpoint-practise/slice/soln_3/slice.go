package main

func Slice(a []string, nbrs ...int) []string {
	var (
		start int
		end   int = len(a)
	)

	if len(nbrs) >= 1 {
		start = nbrs[0]

		if start > len(a) {
			start = len(a)
		}

		if start < 0 {
			start += len(a)
		}

		if start < 0 {
			start = 0
		}
	}

	if len(nbrs) >= 2 {
		end = nbrs[1]

		if end > len(a) {
			end = len(a)
		}

		if end < 0 {
			end += len(a)
		}

		if end < 0 {
			end = 0
		}

		if end < start {
			return []string(nil)
		}
	}
	return a[start:end]
}
