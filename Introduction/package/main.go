package main

import (
	// "learning/package/datasource/sql" // Pemanggilan normal tanpa alias
	_ "learning/package/datasource/sql" // Alias
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // 1 Core Processor

	// mysqlConfig := sql.MySQLConfig{}
	// mysqlConfig := msql.MySQLConfig{}
	// _ = msql.MySQLConfig{}
}
