package main

import "flag"

func main() {
	newmonth := flag.String("newmonth", "", "Neue Datei für Monat X anlegen")
	monthTpl := flag.String("monthtpl", "", "Pfad zu Monats-Datei Template")

	flag.Parse()

	if *newmonth != "" && *monthTpl != "" {
		NewMonthFile(*newmonth, *monthTpl)
	}
}
