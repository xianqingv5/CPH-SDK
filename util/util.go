package util

import "net/url"

func AddurlParam(key, val string, v *url.Values)  {
	if len(val) != 0 {
		v.Add(key, val)
	}
}
