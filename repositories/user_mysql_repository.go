package repositories

import (
	"database/sql"

	"github.com/go_sql_study2/db"
	"github.com/go_sql_study2/libs"
	"github.com/go_sql_study2/models"
)

const user_find_by_all_query = "SELECT id, nickname, create_at, update_at FROM user"
const user_find_by_name_query = user_find_by_all_query + " WHERE nickname = ?;"
const user_create_query = "INSERT INTO user (`id`,`nickname`,`create_at`, `update_at`) VALUES (?, ?, ?, ?);"
const user_today_create_count_query = "INSERT INTO dayJoinCnt (yyyymmdd, cnt) VALUES (?, ?) ON DUPLICATE KEY UPDATE cnt = cnt +1;"
const user_update_query = "UPDATE user SET `update_at` = ? WHERE `id` = ?;"

type userMySQLRepository struct {
	err  error
	stmt *sql.Stmt
	rows *sql.Rows
	tx   *sql.Tx
}

func UserMySQLRepo() UserRepo {
	return &userMySQLRepository{}
}

func (m *userMySQLRepository) FindByName(user *models.User) (*models.User, error) {
	m.stmt, m.err = db.DBConn.UserDB.Conn.Prepare(user_find_by_name_query)
	if m.err != nil {
		return nil, m.err
	}
	defer m.stmt.Close()
	userModel := &models.User{}
	m.err = m.stmt.QueryRow(user.NickName).Scan(&userModel.Id, &userModel.NickName, &userModel.Create_at, &userModel.Update_at)
	if m.err != nil && m.err != sql.ErrNoRows { // sql.ErrNoRows == sql: no rows in result set
		return nil, m.err
	}
	return userModel, nil
}

func (m *userMySQLRepository) FindByAllUser() ([]*models.User, error) {
	m.stmt, m.err = db.DBConn.UserDB.Conn.Prepare(user_find_by_all_query)
	if m.err != nil {
		return nil, m.err
	}
	defer m.stmt.Close()
	m.rows, m.err = m.stmt.Query()
	if m.err != nil {
		return nil, m.err
	}
	defer m.rows.Close()
	users := []*models.User{}
	for m.rows.Next() {
		user := &models.User{}
		m.err = m.rows.Scan(&user.Id, &user.NickName, &user.Create_at, &user.Update_at)
		users = append(users, user)
	}
	m.err = m.rows.Err()
	if m.err != nil {
		return nil, m.err
	}
	return users, nil
}

func (m *userMySQLRepository) CreateUser(user *models.User) (*models.User, error) {
	m.tx, m.err = db.DBConn.UserDB.Conn.Begin()
	if m.err != nil {
		return nil, m.err
	}
	ts, yyyymmdd := libs.GetTime()
	user.Id = libs.GetUUID()
	user.Create_at = ts
	today := yyyymmdd
	_, m.err = m.tx.Exec(user_create_query, user.Id, user.NickName, ts, 0)
	if m.err != nil {
		m.tx.Rollback()
		return nil, m.err
	}
	_, m.err = m.tx.Exec(user_today_create_count_query, today, 1)
	if m.err != nil {
		m.tx.Rollback()
		return nil, m.err
	}
	m.tx.Commit()
	return user, nil
}

func (m *userMySQLRepository) UpdateUser(user *models.User) (*models.User, error) {
	ts, _ := libs.GetTime()
	_, m.err = db.DBConn.UserDB.Conn.Exec(user_update_query, ts, user.Id)
	user.Update_at = ts
	return user, m.err
}
