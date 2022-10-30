package main

import "fmt"

func errType(x V) E {
	return E("bad type: `" + x.Type())
}

func errwType(w V) E {
	return E("left argument: bad type: `" + w.Type())
}

func errs(s string) E {
	return E(s)
}

func errsw(s string) E {
	return E("left argument:" + s)
}

func errf(format string, a ...interface{}) E {
	return E(fmt.Sprintf(format, a...))
}
