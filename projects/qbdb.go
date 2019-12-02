package projects

import (
	"database/sql"
	"fmt"
)

func byId(db *sql.DB, id int64) bool {
	q := fmt.Sprintf("select project_id from quickbase_project_central where project_id = %d", id)
	rows, err := db.Query(q)
	if err != nil {
		Logger.Fatalw("Error happened while querying quickbase projects", "id", id, "error", err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			Logger.Fatalw("Error closing Query", "error", err, "query", q)
		}
	}()

	exists := rows.Next()
	Logger.Debugw("Project exists?", "id", id, "exists", exists)
	return exists
}

func Upsert(db *sql.DB, project ProjectDoc) {
	insertStmt := `INSERT INTO quickbase_project_central (project_id, project_status, partner, client, 
project_name, product, related_client, start_date, end_date, contract_status, client_billable, associated_contract_ids,
vertical, hs_product, prebill_media, prebill_management, contract_type, deal_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
$11, $12, $13, $14, $15, $16, $17, $18)`
	updateStmt := `UPDATE quickbase_project_central SET(project_status, partner, client, project_name, product,
related_client, start_date, end_date, contract_status, client_billable, associated_contract_ids, vertical,
hs_product, prebill_media, prebill_management, contract_type, deal_type) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,
$12, $13, $14, $15, $16, $17) WHERE project_id=$18`

	projects := project.Projects

	for i := 0; i < len(projects); i++ {
		project := projects[i]

		if byId(db, project.ProjectId) == true {
			//Update
			pstmt, err := db.Prepare(updateStmt)
			if err != nil {
				Logger.Fatalw("Preparing upsert statement failed", "error", err, "statement", updateStmt)
			}

			res, err := pstmt.Exec(project.ProjectStatus, project.Partner, project.Client, project.ProjectName,
				project.Product, project.RelatedClient, project.StartDate(), project.EndDate(), project.ContractStatus,
				project.ClientBillable, project.AssociatedContractIds, project.Vertical, project.HsProduct,
				project.PrebillMedia, project.PrebillManagement, project.ContractType, project.DealType, project.ProjectId)

			if err != nil {
				Logger.Fatalw("Executing upsert statement failed", "error", err, "statement", updateStmt)
			}

			rowCnt, err := res.RowsAffected()
			if err != nil {
				Logger.Fatalw("Retrieving affected rows of updating project failed", "error", err, "statement", updateStmt)
			}

			Logger.Debugw("Updated Project", "rowCount", rowCnt, "projectId", project.ProjectId)
		} else {
			//Insert
			istmt, err := db.Prepare(insertStmt)
			if err != nil {
				Logger.Fatalw("Preparing insert statement failed", "error", err, "statement", insertStmt)
			}

			res, err := istmt.Exec(project.ProjectId, project.ProjectStatus, project.Partner, project.Client,
				project.ProjectName, project.Product, project.RelatedClient, project.StartDate(), project.EndDate(),
				project.ContractStatus, project.ClientBillable, project.AssociatedContractIds, project.Vertical,
				project.HsProduct, project.PrebillMedia, project.PrebillManagement, project.ContractType,
				project.DealType)

			if err != nil {
				Logger.Fatalw("Inserting project failed", "error", err, "insertStatement", insertStmt)
			}

			rowCnt, err := res.RowsAffected()
			if err != nil {
				Logger.Fatalw("Failed to retrieve rows for new projects", "error", err, "statement", insertStmt)
			}
			Logger.Debugw("New Project Inserted", "rowCnt", rowCnt)
		}
	}
}
