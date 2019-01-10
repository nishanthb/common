package common

import "regex"

// if r.MatchString(s) {...}
// for named regexps - regexp.MustCompile(`(?P<name1>\d+)(?P<name2>\D+)...
func SplitRegexMap(r *regexp, s string) map[string]string {
	match := r.FindStringSubmatch(s)
	results := map[string]string{}
	for i, v := range match {
		results[r.SubexpNames()[i]] = v
	}
	return results

}
