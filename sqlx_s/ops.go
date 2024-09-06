package sqlx_s

import (
	queries "github.com/nam2184/generic-queries"
  util "github.com/nam2184/generic-queries/utils"
  "fmt"

)

func Insert[T queries.QueryTypes, I any](getID bool) queries.QueryHandlerFunc[T, I] {
    return func(q *queries.Query[T, I]) {
        if q.A != nil {
            // Handle slice of T
            if q.Tx != nil {
                fmt.Println("Handling insert with transaction")
                for _, item := range q.A {
                    fields, _ := util.Fields[T](item)
                    placeholders, _ := util.GenerateNamedParams[T](item)
                    query := fmt.Sprintf("INSERT INTO %s (%s) VALUES( %s ) RETURNING id", 
                                        item.TableName(), 
                                        fields,
                                        placeholders, 
                                        )
                    // Execute the query for each item in the slice
                    var err error
                    if getID == true {
                      string, args, berr := q.Tx.BindNamed(query, &item)
                      q.Args = args
                      q.Q = append(q.Q, string)
                      err = berr
                    } else {
                      _, err = q.Tx.NamedExec(query, &item)
                    }
                    if err != nil {
                        fmt.Println("Error executing query, rolling back transaction:", err)
                        fmt.Println(query)
                        if rollbackErr := q.Tx.Rollback(); rollbackErr != nil {
                            fmt.Println("Failed to rollback transaction:", rollbackErr)
                        }
                        return
                    }
                }
            } else {
                fmt.Println("Should not handle slice without transaction") 
            }
        }
        fmt.Printf("Insert operation completed successfully\n")
    }
}

func Select[T queries.QueryTypes, I any](constraint string) queries.QueryHandlerFunc[T, I] {
    return func(q *queries.Query[T, I]) {
        // Handle slice of T
        if q.Tx != nil {
            var item T
            fmt.Println("Handling select with transaction")
            fields, _ := util.AllFields[T](item)
            query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", 
                                    fields, 
                                    item.TableName(),
                                    constraint, 
                                    )
            // Execute the query for each item in the slice

            err := q.Tx.Select(&q.A, query)
            if err != nil {
                fmt.Println("Error executing query, rolling back transaction:", err)
                if rollbackErr := q.Tx.Rollback(); rollbackErr != nil {
                    fmt.Println("Failed to rollback transaction:", rollbackErr)
                }
                return
            }
        } else {
                fmt.Println("Should not handle slice without transaction") 
            }
        fmt.Printf("Insert operation completed successfully\n")
        }
}



func Delete[T queries.QueryTypes, I any](constraint string) queries.QueryHandlerFunc[T, I] {
    return func(q *queries.Query[T, I]) {
      if q.A != nil {
            // Handle slice of T
            if q.Tx != nil {
                fmt.Println("Handling slice with transaction")
                for _, item := range q.A {
                    query := fmt.Sprintf("Delete FROM %s WHERE %s", 
                                        item.TableName(), 
                                        constraint, 
                                        )
                    // Execute the query for each item in the slice
                    _, err := q.Tx.NamedExec(query, &item)
                    if err != nil {
                        fmt.Println("Error executing query, rolling back transaction:", err)
                        fmt.Println(query)
                        if rollbackErr := q.Tx.Rollback(); rollbackErr != nil {
                            fmt.Println("Failed to rollback transaction:", rollbackErr)
                        }
                        return
                    }
                }
            } else {
                fmt.Println("Should not handle slice without transaction") 
            }
        }
        fmt.Printf("Insert operation completed successfully\n")
    }
}

func Update[T queries.QueryTypes, I any](constraint string) queries.QueryHandlerFunc[T, I] {
    return func(q *queries.Query[T, I]) {
      if q.A != nil {
            // Handle slice of T
            if q.Tx != nil {
                fmt.Println("Handling slice with transaction")
                for _, item := range q.A {
                    fields, _ := util.Fields[T](item)

                    query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", 
                                        item.TableName(),
                                        fields,
                                        constraint, 
                                        )
                    // Execute the query for each item in the slice
                    _, err := q.Tx.NamedExec(query, &item)
                    if err != nil {
                        fmt.Println("Error executing query, rolling back transaction:", err)
                        fmt.Println(query)
                        if rollbackErr := q.Tx.Rollback(); rollbackErr != nil {
                            fmt.Println("Failed to rollback transaction:", rollbackErr)
                        }
                        return
                    }
                }
            } else {
                fmt.Println("Should not handle slice without transaction") 
            }
        }
        fmt.Printf("Insert operation completed successfully\n")
    }
}
