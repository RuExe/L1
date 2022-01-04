package set

type (
	Set struct {
		hash map[string]nothing
	}

	nothing struct{}
)

func New(slice []string) Set {
	s := Set{make(map[string]nothing)}

	for _, v := range slice {
		s.Insert(v)
	}
	return s
}

func (s *Set) Insert(element string) {
	s.hash[element] = nothing{}
}

func (s *Set) Intersection(set *Set) Set {
	res := make(map[string]nothing)

	for k := range s.hash {
		if _, exists := set.hash[k]; exists {
			res[k] = nothing{}
		}
	}
	return Set{res}
}

func (s *Set) ToSlice() []string {
	res := make([]string, 0, len(s.hash))

	for k := range s.hash {
		res = append(res, k)
	}
	return res
}
