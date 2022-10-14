package mqtt

import "strings"

const (
	TopicSep = "/"
)

func TopicJoin(prefix, sep string, topic ...string) string {
	return strings.Join(append([]string{prefix}, topic...), sep)
}
