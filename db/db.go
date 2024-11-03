package db

import (
	"fmt"

	"xorm.io/xorm"
)

// provide thread / goroutine safe access to database, based on the path name
// when only one single thread call to NewDatabase, it should be thread safe for everything:
// - multiple call to the same NewDatabase / Close is safe because we check it and we only call them on 1 thread.
// - each of the database itself is threadsafe across threads:
//   - assuming that the request functions do not panic without recovery
// - the database object always hold a connection to the actual database, so even if a new database is opened using other method, they can't do anything with it.
// note that different path will be treated as different
// for now, the user database doesn't use this system for a few reason:
// - either the whole request processing has to be done in a request to this db
// - or it's possible to have partially handled user request data in database:
//   - User A and B send a request at roughly the same time, both are handled concurrently.
//   - Both handlers write some data to the database.
//   - User A request is more simple, so it's finished first, commiting A's data and some of B's data.
//   - Then user B request encountered a problem and need to be reverted, which doesn't revert B's data already written, leaving B in a bad state.
//   - Similarly, if A's request encountered a problem and reverted some of B's data, then B would also be in a bad state.
// - the problem can somewhat be eliviated if we just do all the necessary write at the end of request:
//   - if there is a problem, do not write anything.
//   - otherwise, acquire the lock and write everything and commit too.
// - howerver, the above can still lead to some weird outcome:
//   - The userdata are not totally separated by user.
//   - A's data can have an impact on how B's request is handled.
//   - For example, if both A and B send a friend request to each other, then we can potentially ended up with 2 out going requests.
// - TLDR it's easier to just keep each request atomic, and if that's the case then each request only need its own session.

type DatabaseSync struct {
	path           string
	engine         *xorm.Engine
	session        *xorm.Session
	requestChannel chan func(session *xorm.Session)
	syncChannel    chan struct{}
}

var databases = map[string]*DatabaseSync{}

func NewDatabase(path string) (d *DatabaseSync, err error) {
	var exists bool
	d, exists = databases[path]
	if exists {
		return
	}
	d = &DatabaseSync{
		path: path,
	}
	d.engine, err = xorm.NewEngine("sqlite", path)
	if err != nil {
		return
	}
	d.session = d.engine.NewSession()
	err = d.session.Begin()
	if err != nil {
		d.session.Close()
		return
	}
	d.requestChannel = make(chan func(session *xorm.Session))
	d.syncChannel = make(chan struct{})
	go d.serve()
	return
}

func (d *DatabaseSync) serve() {
	for {
		f := <-d.requestChannel
		if f == nil {
			break
		}
		f(d.session)
		d.syncChannel <- struct{}{}
	}
	// by default, commit nothing
	// so each function f has to commit if it wants the result to be recorded
	d.session.Rollback()
	d.session.Close()
	d.syncChannel <- struct{}{}
}

func (d *DatabaseSync) Do(f func(session *xorm.Session)) {
	if f == nil {
		fmt.Println("nil function ignored. If you want to stop, use .Close() instead!")
	}
	d.requestChannel <- f
	<-d.syncChannel
}

func (d *DatabaseSync) Close() error {
	d.requestChannel <- nil
	<-d.syncChannel
	d.engine.Close()
	delete(databases, d.path)
	return nil
}
