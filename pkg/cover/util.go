package cover

import "regexp"

func mergeMaps(maps ...map[string]string) map[string]string {
	merged := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}

func namedCaptures(re *regexp.Regexp, s string) map[string]string {
	matches := make(map[string]string)

	submatches := re.FindStringSubmatch(s)
	if submatches == nil {
		return matches
	}

	for i, name := range re.SubexpNames() {
		if len(name) == 0 {
			// skip the match that just confirms there was a match
			continue
		}
		matches[name] = submatches[i]
	}
	return matches
}
