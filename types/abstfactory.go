package types

import "fmt"

/*
Build a database driver system:

Two product interfaces — Connection with Connect() string and Query with Execute(sql string) string
Two families — MySQL and Postgres, each with their own Connection and Query implementations
DatabaseFactory interface with CreateConnection() Connection and CreateQuery() Query
MySQLFactory and PostgresFactory implementing DatabaseFactory
A DatabaseClient struct that takes a DatabaseFactory, has a Run(sql string) method that creates a connection and executes a query
In main, run the same SQL through both factories
*/

type Connection interface {
	Connect() string
}

type Query interface {
	Execute(sql string) string
}

type MysqlConnection struct{}
type PostgresConnection struct{}

func (m *MysqlConnection) Connect() string {
	return "connection string for mysql"
}

func (m *PostgresConnection) Connect() string {
	return "connection string for postgres"
}

type MysqlQuery struct{}
type PostgresQuery struct{}

func (m *MysqlQuery) Execute(s string) string {
	return "executed s query (mysql)"
}

func (m *PostgresQuery) Execute(s string) string {
	return "executed s query (postgres)"
}

type DatabaseFactory interface {
	CreateConnection() Connection
	CreateQuery() Query
}

type MySQLFactory struct{}

func (m *MySQLFactory) CreateConnection() Connection {
	return &MysqlConnection{}
}

func (m *MySQLFactory) CreateQuery() Query {
	return &MysqlQuery{}
}

type PostgresFactory struct{}

func (m *PostgresFactory) CreateConnection() Connection {
	return &PostgresConnection{}
}

func (m *PostgresFactory) CreateQuery() Query {
	return &PostgresQuery{}
}

type DatabaseClient struct {
	factory DatabaseFactory
}

func NewDatabaseClient(factory DatabaseFactory) *DatabaseClient {
	return &DatabaseClient{factory: factory}
}

func (db *DatabaseClient) Run(q string) {
	connection := db.factory.CreateConnection()
	query := db.factory.CreateQuery()

	fmt.Println(connection.Connect())
	fmt.Println(query.Execute(q))
}
