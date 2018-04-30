//package main
package TestSqlDatabaseMssql

//這裡大量照抄
//https://github.com/denisenkom/go-mssqldb/blob/master/examples/tsql/tsql.go#L19
//go-mssqldb/examples/tsql/tsql.go
//作為練習

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	var (
		userid   = flag.String("U", "sa", "login_id")
		password = flag.String("P", "1111", "password")
		server   = flag.String("S", "MSI\\SQL2K14", "server_name[\\instance_name]")
		database = flag.String("d", "FET2017", "db_name")
	)
	flag.Parse()

	dsn := "server=" + *server + ";user id=" + *userid + ";password=" + *password + ";database=" + *database
	//這句open還是印象深刻   帶領我進介面大門的鑰匙!!!
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}

	//db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}

	defer db.Close()
	r := bufio.NewReader(os.Stdin)
	for {
		_, err = os.Stdout.Write([]byte("> "))
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return
			}
			fmt.Println(err)
			return
		}
		err = exec(db, cmd)
		if err != nil {
			fmt.Println(err)
		}
	}

}

//main end

func exec(db *sql.DB, cmd string) error {
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	if cols == nil {
		return nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}
		fmt.Print(cols[i])
	}
	fmt.Println()
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()

	}
	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}

func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}

func ExportTest() string {
	return "ttt"
}
