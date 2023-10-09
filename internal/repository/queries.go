package repository

const (
	insertNewCostumer  = `INSERT INTO users(NIK, full_name, born_place, born_date,is_admin, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)`
	getCostumerByEmail = `
        SELECT
            id,
            NIK,
            full_name,
            born_place,
            born_date,
            is_admin,
            email,
            password,
            created_at,
            updated_at,
            deleted_at
        FROM
            users
        WHERE
            email = ?
    `
	getUserBirthdayByDate = `
 SELECT
            id,
            NIK,
            full_name,
            born_place,
            born_date,
            is_admin,
            email,
            password,
            created_at,
            updated_at,
            deleted_at
        FROM
            users
        WHERE
            born_date = ?
`
	createCampaign = `INSERT INTO campaigns (code, name, amount, start_date, end_date, quota)
		VALUES(?,?,?,?,?,?)`
	createUserCampaign = `INSERT INTO user_campaigns (user_id, campaign_id)
		VALUES
			(?, ?)`
	getCampaignByCode = `SELECT id, code, name, amount, start_date, end_date, quota, created_at, updated_at, deleted_at
		FROM campaigns
		WHERE code = '?';`

	getCampaignByUserID = `SELECT campaigns.*
		FROM campaigns
		INNER JOIN user_campaigns ON campaigns.id = user_campaigns.campaign_id
		WHERE user_campaigns.user_id = ?;`
	getProductById = `SELECT id, name, type, model_product, price, profit, product_manager, release_date, created_at, updated_at, deleted_at
		FROM products
		WHERE id = ?;
		`
	getAllProduct = `SELECT id, name, type, model_product, price, profit, product_manager, release_date, created_at, updated_at, deleted_at
		FROM products`

	createTransactionWithPromo = `INSERT INTO transactions (user_id, product_id, campaign_id, total, amount, status)
						VALUES
							(?, ?, ?, ?, ?, ?);`
	createTransactionWithoutPromo = `INSERT INTO transactions (user_id, product_id, total, amount, status)
						VALUES
							(?, ?, ?, ?, ?);`

	updateQuotaVoucher = `
UPDATE campaigns
SET quota = ?
WHERE id = ?
`
)
