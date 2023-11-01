# Transaction-Generator

This transaction tries to make multiple sql queries atomicited in **multiple different repositories**.

The transaction-generator should inherit all repositories.
```go
type transactionGenerator struct {
	db *sqlx.DB
	tx *sqlx.Tx

	repo1
	repo2
}
```
* `repo1` and `repo2` are repositories structs.

Also, This transaction must implement two interface:

```go
type TransactionGeneratpr interface{
    Begin() ConcreteTransaction
}

type Transaction interface{
    Commit() error
    Rollback() error
}
```

## Begin method

This method should take a transcation (`tx`) from `t.db.Beginx()` method and set repositories' transaction equal to it.

* `t` is transaction generator struct.

```go
t.tx = tx
t.repo1.tx = tx
t.repo2.tx = tx
```

This method should return an object that has all capabilities of all included repositories. 

This `Begin` method's signature should be same as `db.Beginx()`'s signature. So it must return an interface and an error.

The returned interface in this pattern is called `ConcreteTransaction` and is like this:

```go
type ConcreteTransaction interface {
	Repo1 
	Repo2 

	Transaction
}
```

`Repo1` and `Repo2` are the repository interfaces (Consists querying methods like Get methods and Update methods).

`Transaction` is a new interface that has `Rollback` and `Commit` methods.

## Commit method

When a transaction is Committed or Roled back, all transactions (all repositories transactions and the transaction-generator struct's transaction should be nill)

```go
if t.tx == nil {
    return ErrTransactionNotStarted
}

err := t.tx.Commit()
if err != nil {
    return err
}

t.tx = nil
t.repo1.tx = nil
t.repo2.tx = nil
```

## Rollback method

This method is very similar to commit method:
```go
if t.tx == nil {
		return ErrTransactionNotStarted
	}

	err := t.tx.Rollback()
	if err != nil {
		return err
	}

	t.tx = nil
	t.waitingScreen.tx = nil
	t.waitingScreenRevision.tx = nil

	return nil
```