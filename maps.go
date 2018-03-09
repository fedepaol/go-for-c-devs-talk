var values = make(map[string]int)

values["pr"] = 23
values["bdu"] = 28
delete(m, "bdu")

if age, ok := values["pr"]; !ok { /* ... */ } // if not found, age = 0 value


