package main

import (
	"database/sql"
	"fmt"
	"math"

	"qbprojects/projects"
)

func main() {
	projects.ProductionLogger()
	defer projects.Logger.Sync()

	cnf := projects.ReadConf("prod")
	connstr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", cnf.DbUser, cnf.DbName, cnf.DbPassword, cnf.DbHost)

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		projects.Logger.Fatalw("Failed to open database connection", "error", err)
	}

	defer db.Close()

	ticket := projects.Auth(cnf.QbUsername, cnf.QbPassword)

	cnt := projects.CountQbProjects(ticket.Ticket, cnf.ProjectCentralTable)
	sz := 100
	pages := int(math.Ceil(float64(cnt.Total) / float64(sz)))

	projects.Logger.Debugw("Number of Pages", "pages", pages)
	i := 0
	var projectDoc projects.ProjectDoc
	for i < pages {
		projectDoc = projects.GetProjects(ticket.Ticket, cnf.ProjectCentralTable, cnf.QbAppToken, sz, i*sz)
		projects.Upsert(db, projectDoc)
		i++
	}
}
