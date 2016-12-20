package main

type rule struct {
	from uint
	to   uint
}

/** Sort interface for []rule **/
type byRange []rule

func (rules byRange) Len() int {
	return len(rules)
}

func (rules byRange) Less(i, j int) bool {
	if rules[i].from == rules[j].from {
		return rules[i].to < rules[j].to
	}
	return rules[i].from < rules[j].from
}

func (rules byRange) Swap(i, j int) {
	rules[i], rules[j] = rules[j], rules[i]
}

// Assume rules are sorted and they have no overlap
func normalize(rules []rule, newRule rule) []rule {
	for i := range rules {
		if overlappingOrAdjacent(rules[i], newRule) {
			rules[i].to = max(newRule.to, rules[i].to)
			return rules
		}
	}
	rules = append(rules, newRule)
	return rules
}

func overlappingOrAdjacent(one rule, other rule) bool {
	return (one.from <= other.to && other.from <= one.to) ||
		one.to+1 == other.from ||
		other.to+1 == one.from
}
