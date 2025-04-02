package exporter

func Export(exptr string, data any, pretty bool) {
	switch exptr {
	case "json":
		FmtJson(data, pretty)
	case "csv":
		FmtCsv(data, pretty)
	case "excel":
		FmtExcel(data, pretty)
	default:
		FmtTable(data, pretty)
	}
}
