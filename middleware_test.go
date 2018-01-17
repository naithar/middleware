package middleware

import (
	"testing"
)

type cleanValue struct {
	unclean string
	clean string
}

var cleanValues = []cleanValue {
	{ "/", "/" },
	{ "////", "/" },
	{ "/a", "/a" },
	{ "/a/", "/a" },
	{ "///a/", "/a" },
	{ "///a///", "/a" },
	{ "/a/b", "/a/b" },
	{ "///a/b", "/a/b" },
	{ "///a///b", "/a/b" },
}

var cleanValuesTrailing = []cleanValue {
	{ "/", "/" },
	{ "////", "/" },
	{ "/a", "/a" },
	{ "/a/", "/a/" },
	{ "///a/", "/a/" },
	{ "///a///", "/a/" },
	{ "/a/b", "/a/b" },
	{ "///a/b", "/a/b" },
	{ "///a///b", "/a/b" },
	{ "///a///b/", "/a/b/" },
	{ "///a///b//", "/a/b/" },
	{ "///a///b//c/", "/a/b/c/" },
}
func TestCleanPath(t *testing.T) {
	for _, value := range cleanValues {
		if result := CleanPath(value.unclean, false); result != value.clean {
			t.Errorf("CleanPath(%q, false) = %q. Expected %q", value.unclean, result, value.clean)
		}
	}
}
func TestCleanPathTrailing(t *testing.T) {
	for _, value := range cleanValuesTrailing {
		if result := CleanPath(value.unclean, true); result != value.clean {
			t.Errorf("CleanPath(%q, false) = %q. Expected %q", value.unclean, result, value.clean)
		}
	}
}