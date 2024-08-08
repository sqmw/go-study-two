package user

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	// 创建 mock 数据库连接
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建 mock 失败: %s", err)
	}
	defer db.Close()

	// 创建 UserRepo 实例
	repo := NewUserRepo(db)

	// 测试用例
	tests := []struct {
		name     string
		userID   int
		mockFunc func()
		want     *User
		wantErr  bool
	}{
		{
			name:   "成功查询用户",
			userID: 1,
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "age"}).
					AddRow(1, "张三", 25)
				mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
					WithArgs(1).
					WillReturnRows(rows)
			},
			want:    &User{ID: 1, Name: "张三", Age: 25},
			wantErr: false,
		},
		{
			name:   "用户不存在",
			userID: 999,
			mockFunc: func() {
				mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
					WithArgs(999).
					WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: true,
		},
	}

	// 执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置 mock 预期
			tt.mockFunc()

			// 执行测试
			got, err := repo.GetUserByID(tt.userID)

			// 验证结果
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			// 确保所有预期都被满足
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("有未满足的预期: %s", err)
			}
		})
	}
	fmt.Println("测试结束")
}
