package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type config struct {
	fields         int
	delimiter      string
	hideWrongLines bool
}

func parseFlags() config {
	cfg := config{}

	flag.IntVar(&cfg.fields, "f", 0, "показать все столбцы или только этот столбец")
	flag.StringVar(&cfg.delimiter, "d", "\t", "разделитель")
	flag.BoolVar(&cfg.hideWrongLines, "s", true, "скрывать ли строки без разделителя")
	flag.Parse()

	return cfg
}

func cut(sc *bufio.Scanner, cfg config) {
	// принцип работы похож на сканнер результатов запроса в
	// постгрес
	for sc.Scan() {
		line := sc.Text()
		if !cfg.hideWrongLines && !strings.Contains(line, cfg.delimiter) {
			continue
		}
		columns := strings.Split(line, cfg.delimiter)

		if cfg.fields == 0 {
			fmt.Println(strings.Join(columns, " "))
		} else {
			if len(columns) >= cfg.fields {
				fmt.Println(columns[cfg.fields-1])
			}
		}
	}
}

func main() {
	cfg := parseFlags()
	// используем интерфейс из пакета bufio,
	// предназначенный для работы с файлами, где строки
	// разделяются с помощью \n
	sc := bufio.NewScanner(os.Stdin)

	cut(sc, cfg)
}
