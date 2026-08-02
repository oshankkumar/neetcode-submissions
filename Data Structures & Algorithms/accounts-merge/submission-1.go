
import (
	"maps"
)

type EmailSet map[string]struct{}

func (e EmailSet) Emails() []string {
	var result []string
	for k := range e {
		result = append(result, k)
	}
	return result
}

func (e EmailSet) Len() int {
	return len(e)
}

func (e EmailSet) Intersection(e2 EmailSet) EmailSet {
	result := make(EmailSet)
	for key := range e2 {
		if _, ok := e[key]; ok {
			result[key] = struct{}{}
		}
	}
	return result
}

func (e EmailSet) Union(e2 EmailSet) EmailSet {
	result := maps.Clone(e)
	for key := range e2 {
		result[key] = struct{}{}
	}
	return result
}

func NewEmailSet(emails ...string) EmailSet {
	set := make(EmailSet)
	for _, email := range emails {
		set[email] = struct{}{}
	}
	return set
}

func accountsMerge(accounts [][]string) [][]string {
	userAccounts := make(map[string][]EmailSet)

	for _, account := range accounts {
		name, emailSets := account[0], NewEmailSet(account[1:]...)

		existingEmailSets, ok := userAccounts[name]
		if !ok {
			userAccounts[name] = append(userAccounts[name], emailSets)
			continue
		}

		var matchedIdx []int
		var newEmailSets []EmailSet
		for i, existing := range existingEmailSets {
			if existing.Intersection(emailSets).Len() > 0 {
				matchedIdx = append(matchedIdx, i)
			} else {
				newEmailSets = append(newEmailSets, existing)
			}
		}

		for _, idx := range matchedIdx {
			emailSets = emailSets.Union(existingEmailSets[idx])
		}

		newEmailSets = append(newEmailSets, emailSets)

		userAccounts[name] = newEmailSets
	}

	var results [][]string

	for name, emailSets := range userAccounts {
		for _, emailSet := range emailSets {
			emails := emailSet.Emails()
			sort.Strings(emails)
			userAccs := make([]string, len(emails)+1)
			userAccs[0] = name
			copy(userAccs[1:], emails)
			results = append(results, userAccs)
		}
	}

	return results
}
