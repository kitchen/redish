package engine

type setValueStore struct {
	storage map[string]interface{}
}

func (s *setValueStore) isMember(member string) bool {
	_, ok := s.storage[member]
	return ok
}

func (s *setValueStore) add(members []string) int64 {
	var added int64
	for _, member := range members {
		if !s.isMember(member) {
			s.storage[member] = nil
			added++
		}
	}
	return added
}

func (s *setValueStore) rem(members []string) int64 {
	var deleted int64
	for _, member := range members {
		if s.isMember(member) {
			delete(s.storage, member)
			deleted++
		}
	}
	return deleted
}

func (s *setValueStore) members() []string {
	members := make([]string, len(s.storage))
	var i int64
	for member := range s.storage {
		members[i] = member
		i++
	}
	return members
}
