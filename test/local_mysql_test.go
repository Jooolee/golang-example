package test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// User 定义用户结构体
type User struct {
	ID    int
	Name  string
	Email string
}

// 数据库连接信息
const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "public"
)

// 构建数据库连接字符串
func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

// 创建用户
func createUser(db *sql.DB, name, email string) (int64, error) {
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 获取所有用户
func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// 根据 ID 获取用户
func getUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
		return User{}, err
	}
	return user, nil
}

// 更新用户信息
func updateUser(db *sql.DB, id int, name, email string) error {
	_, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", name, email, id)
	return err
}

// 删除用户
func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func TestLocalMysql(*testing.T) {
	// 打开数据库连接
	db, err := sql.Open("mysql", getDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 创建用户
	newUserID, err := createUser(db, "John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created user with ID: %d\n", newUserID)

	// 获取所有用户
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	// 根据 ID 获取用户
	user, err := getUserByID(db, int(newUserID))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User with ID %d: Name: %s, Email: %s\n", user.ID, user.Name, user.Email)

	// 更新用户信息
	err = updateUser(db, int(newUserID), "Jane Doe", "jane.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User updated successfully")

	// 删除用户
	err = deleteUser(db, int(newUserID))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully")
}
