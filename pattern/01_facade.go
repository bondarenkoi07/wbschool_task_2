package pattern

// Фасад — это простой интерфейс для работы со сложной подсистемой, содержащей множество классов.
// Плюсы: Изолирует клиентов от компонентов сложной подсистемы.
// Минусы:Фасад рискует стать божественным объектом, привязанным ко всем классам программы.

// В данном примере структура DB является фасадам к библиотеке pgx.
import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
	"os"
)

const (
	ErrorZeroRowAffected = "zero rows affected"
)

var dsn string

type foobar interface {
	GetFields() ([]interface{}, string)
}

// DB : Фасад к библиотеке pgx. Упрощает работу с таблицами в pg, необходимыми нам в соответствии с ТЗ.
// Предоставляет возможность создавать в таблице строки, ассоциируемые с моделями,
// удовлетворяющими интерфейсу foobar в заданной таблице с помощью метода Create
// Также позволяет считывать данные с помощью readRaw and ReadOrdersRaw
// Begin, End являются обёртками для начала и завершения транзакции (pgx.Tx).
type DB struct {
	// Пул подключений
	pool *pgxpool.Pool
	ctx  context.Context
}

func init() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s",
		user,
		password,
		"wb-go-team-dev.dev.wb.ru",
		database,
	)
}

func _() (DB, error) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)

	return DB{
		pool: pool,
		ctx:  ctx,
	}, err
}

func (d *DB) Begin() (*pgx.Tx, error) {
	begin, err := (*d).pool.Begin((*d).ctx)

	return &begin, err
}

func (d *DB) End(err error, transaction *pgx.Tx) error {

	if err != nil {
		superErr := (*transaction).Rollback((*d).ctx)
		if superErr != nil {
			return errors.New(superErr.Error() + "\nwhile handling\n" + err.Error())
		}
	} else {
		err = (*transaction).Commit((*d).ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) Create(tableName string, model foobar, transaction *pgx.Tx) error {
	values, cols := model.GetFields()
	valuesPlaceholders := ""
	count := len(values) - 1
	index := 0

	for ; index < count; index++ {
		valuesPlaceholders += fmt.Sprintf("$%d", index+1)
		valuesPlaceholders += ","
	}
	valuesPlaceholders += fmt.Sprintf("$%d", index+1)

	SQLStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, cols, valuesPlaceholders)

	ct, err := (*transaction).Exec(d.ctx, SQLStatement, values...)

	if err != nil {
		return errors.New(fmt.Sprintf("%s insertion error:%v", tableName, err))
	} else if ct.RowsAffected() == 0 {
		return errors.New(ErrorZeroRowAffected)
	} else {
		return nil
	}
}

func (d *DB) readRaw(tableName, orderCol string) (*pgx.Rows, error) {
	SQLStatement := fmt.Sprintf("SELECT * FROM %s ORDER BY %s", tableName, orderCol)
	conn, err := (*d).pool.Acquire(d.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(d.ctx, SQLStatement)

	if err != nil {
		return nil, err
	} else {
		return &rows, nil
	}
}

func (d *DB) ReadOrdersRaw() (pgx.Rows, error) {
	SQLStatement := ` 
	SELECT o.order_uid, o.entry,SUM(i.total_price)+delivery_cost as total_price, o.customer_id,
		o.track_number, o.delivery_service
	FROM task_order AS o
		JOIN items AS i
			ON o.order_uid = i.orderid
		JOIN payment as p
			ON o.order_uid = p.orderid
	GROUP BY o.order_uid,p.delivery_cost
`
	return d.pool.Query((*d).ctx, SQLStatement)
}
