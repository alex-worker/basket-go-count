package informator

import "database/sql"

const QuerySQL = `
	SELECT b.season, (CASE WHEN b.country = 'USA' THEN 'USA' ELSE 'Others' END) as is_usa, count(b.player_name)
	FROM basketball b
	GROUP BY b.season, is_usa
	ORDER BY season, is_usa DESC;
`

type SeasonRecord struct {
	Season string
	IsUSA  string
	Count  int
}

type SeasonRecords []SeasonRecord

func DoSeasonRecordsCount(db *sql.DB) (SeasonRecords, error) {
	rows, err := db.Query(QuerySQL)
	if err != nil {
		return nil, err
	}

	var resp SeasonRecords

	for rows.Next() {
		var season string
		var isUSA string
		var count int

		err := rows.Scan(&season, &isUSA, &count)
		switch err {
		case sql.ErrNoRows:
			return resp, nil
		case nil:
			resp = append(resp, SeasonRecord{
				season,
				isUSA,
				count,
			})
		default:
			return resp, err
		}
	}

	err = rows.Close()
	if err != nil {
		return resp, err
	}
	return resp, nil
}
