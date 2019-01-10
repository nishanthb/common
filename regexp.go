package common

import "regex"

// if r.MatchString(s) {...}
func SplitRegexMap(r *regexp, s string) map[string]string {
	match := r.FindStringSubmatch(s)
	results := map[string]string{}
	for i, v := range match {
		results[r.SubexpNames()[i]] = v
	}
	return results

}
