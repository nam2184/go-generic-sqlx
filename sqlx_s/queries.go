package sqlx_s

import (
	queries "github.com/nam2184/generic-queries"

	"github.com/jmoiron/sqlx"
)

/*
These generic queries create a transaction struct if none is passed, thus defining the query for you,
suitable for one specific query done routes and routes that handle multiple queries.QueryTypes
*/

//If you want to save memory by having specific *queries.Transactions passed, define InsertQuery[T](nil, tran, data)

func InsertQuery[T queries.QueryTypes, I any](tx *sqlx.Tx, tran *queries.Transaction[T, I], data []T) *queries.Query[T, I] {
    if tran == nil {
        // Create a new transaction if not provided
        tran = queries.NewTransaction[T, I](Insert[T, I](false), tx)
    }
    
    qs := queries.NewQueryMany[T, I](data, tran.Tx)
    tran.Handler.HandleQuery(qs)
    return qs
}

func InsertQueryID[T queries.QueryTypes, I any](tx *sqlx.Tx, tran *queries.Transaction[T, I], data []T) *queries.Query[T, I] {
    if tran == nil {
        // Create a new transaction if not provided
        tran = queries.NewTransaction[T, I](Insert[T,I](true), tx)
    }
    
    qs := queries.NewQueryMany[T, I](data, tran.Tx)
    tran.Handler.HandleQuery(qs)   
 
    return qs
}


func DeleteQuery[T queries.QueryTypes, I any](tx *sqlx.Tx, constraint string, tran *queries.Transaction[T, I], data []T)  *queries.Query[T, I] {
     if tran == nil {
        // Create a new transaction if not provided
        tran = queries.NewTransaction[T, I](Delete[T, I](constraint), tx)
    }

    qs := queries.NewQueryMany[T, I](data, tran.Tx)
    tran.Handler.HandleQuery(qs) 
    return qs
}

func SelectQuery[T queries.QueryTypes, I any](tx *sqlx.Tx, constraint string, tran *queries.Transaction[T, I], data []T) *queries.Query[T, I] {
    if tran == nil {
        // Create a new transaction if not provided
        tran = queries.NewTransaction[T, I](Select[T, I](constraint), tx)
    }
    
    qs := queries.NewQueryMany[T, I](data, tran.Tx)
    tran.Handler.HandleQuery(qs) 
    return qs
}

func UpdateQuery[T queries.QueryTypes, I any](tx *sqlx.Tx, constraint string, tran *queries.Transaction[T, I], data []T) *queries.Query[T, I] {
    if tran == nil {
        tran = queries.NewTransaction[T, I](Update[T, I](constraint), tx)
    }

    qs := queries.NewQueryMany[T, I](data, tran.Tx)
    tran.Handler.HandleQuery(qs) 
    return qs
}

