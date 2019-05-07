package logy

func logParams(args []string) []Field {
	if len(args)%2 != 0 {
		Warn("log: the variadic must be plural, the last one will ignored",nil)
	}
	ds := make([]Field, 0, len(args)/2)
	for i := 0; i < len(args)-1; i = i + 2 {
		ds = append(ds, newField(args[i], args[i+1]))
	}
	return ds
}
