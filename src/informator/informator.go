package informator

const QuerySQL = `
	SELECT b.season, (CASE WHEN b.country = 'USA' THEN 'USA' ELSE 'Others' END) as is_usa, count(b.player_name)
	FROM basketball b
	GROUP BY b.season, is_usa
	ORDER BY season, is_usa DESC;
`
