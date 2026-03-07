package db

import (
	"context"
	"database/sql"
	"fmt"
)

type PreparedStatements struct {
	Statements map[string]*sql.Stmt
}

func NewPreparedStatements() *PreparedStatements {
	return &PreparedStatements{
		Statements: make(map[string]*sql.Stmt),
	}
}

var PrepStmts = NewPreparedStatements()

func InitPreparedStatements(ctx context.Context, db *sql.DB) error {
	var err error
	defer func() {
		if err != nil {
			err := PrepStmts.Close()
			if err != nil {
				return
			}
		}
	}()

	createPreparedStatement := func(name, statement string) error {
		// Initialize and add a prepared statement
		stmt, err_ := db.PrepareContext(ctx, statement)
		if err_ != nil {
			err_ = fmt.Errorf("failed to prepare [%s] statement: %w", name, err_)
			err = err_
			return err_
		}

		// Check and add the prepared statement for later use
		if _, ok := PrepStmts.Statements[name]; ok {
			err_ = fmt.Errorf("[%s] statement already exists", name)
			err = err_
			return err_
		}

		PrepStmts.Statements[name] = stmt

		return nil
	}
	_ = createPreparedStatement

	err_ := createPreparedStatement("createPostsTable", "CREATE TABLE IF NOT EXISTS public.posts (id SERIAL PRIMARY KEY, title TEXT, content TEXT)")
	if err_ != nil {
		return err_
	}

	err_ = createPreparedStatement("createUsersTable", "CREATE TABLE IF NOT EXISTS public.users (id SERIAL PRIMARY KEY, username TEXT, email TEXT)")
	if err_ != nil {
		return err_
	}

	err_ = createPreparedStatement("getPostById", "SELECT json_build_object('id', id, 'title', title, 'content', content) FROM public.posts WHERE id = $1")
	if err_ != nil {
		return err_
	}

	return nil
}

func (self *PreparedStatements) GetPreparedStatements() []string {
	keys := make([]string, 0, len(self.Statements))

	for k := range self.Statements {
		keys = append(keys, k)
	}

	return keys
}

func (self *PreparedStatements) Close() error {
	for _, stmt := range self.Statements {
		if err := stmt.Close(); err != nil {
			return err
		}
	}
	self.Statements = nil
	return nil
}
