package exporter

func Export(exptr, name string, data any, pretty bool) {
	switch exptr {
	case "json":
		FmtJson(data, pretty)
	case "csv":
		FmtCsv(name, data, pretty)
	case "excel":
		FmtExcel(name, data, pretty)
	default:
		FmtTable(data, pretty)
	}
}
