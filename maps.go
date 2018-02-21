var ages map[string]int = make(map[string]int)

ages["fede"] = 23
ages["lore"] = 28
delete(m, "fede")

if age, ok := ages["fede"]; !ok { /* ... */ } // if not found, age = 0 value


