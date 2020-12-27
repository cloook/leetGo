func isIsomorphic(s string, t string) bool {
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}
	for i := range s {
		sv, tv := s[i], t[i]
		if sMap[sv] > 0 && sMap[sv] != tv || tMap[tv] > 0 && tMap[tv] != sv {
			return false
		}
		sMap[sv], tMap[tv] = tv, sv
	}
	return true
}