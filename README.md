# gocov-merge
Command to merge two or more JSON coverage profiles.

	go install github.com/cloud9-tools/gocov-merge
	{ for pkg in a b c; do gocov test "$pkg"; done } | gocov-merge
