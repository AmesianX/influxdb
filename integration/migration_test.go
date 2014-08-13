package integration

/*
This is commented out because we can't generate old data automatically. The files for this test
are up on S3 so that we can run it later. Just trust that I've run it (this is Paul)
*/

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"github.com/influxdb/influxdb/datastore"
// 	"github.com/influxdb/influxdb/migration"

// 	. "github.com/influxdb/influxdb/integration/helpers"
// 	. "launchpad.net/gocheck"
// )

// type MigrationTestSuite struct {
// 	server *Server
// }

// var _ = Suite(&MigrationTestSuite{})

// func (self *MigrationTestSuite) SetUpSuite(c *C) {
// 	self.server = NewServer("integration/migration_test.toml", c)
// }

// func (self *MigrationTestSuite) TearDownSuite(c *C) {
// 	self.server.Stop()
// 	dataDir := "migration_data/data"
// 	infos, err := ioutil.ReadDir(filepath.Join(dataDir, migration.OLD_SHARD_DIR))
// 	if err != nil {
// 		fmt.Printf("Error Clearing Migration: ", err)
// 		return
// 	}
// 	for _, info := range infos {
// 		if info.IsDir() {
// 			os.Remove(filepath.Join(info.Name(), migration.MIGRATED_MARKER))
// 		}
// 	}
// 	fmt.Println("ClearMigration: ", filepath.Join(dataDir, datastore.SHARD_DATABASE_DIR))
// 	os.RemoveAll(filepath.Join(dataDir, datastore.SHARD_DATABASE_DIR))
// }

// func (self *MigrationTestSuite) TestMigrationOfPreviousDb(c *C) {
// 	time.Sleep(time.Second * 60)
// 	client := self.server.GetClient("test1", c)
// 	s, err := client.Query("select count(value) from cpu_idle")
// 	c.Assert(err, IsNil)
// 	c.Assert(s, HasLen, 1)
// 	c.Assert(s[0].Points, HasLen, 1)
// 	c.Assert(s[0].Points[0][1].(float64), Equals, float64(44434))

// 	s, err = client.Query("select count(type) from customer_events")
// 	c.Assert(err, IsNil)
// 	c.Assert(s, HasLen, 1)
// 	c.Assert(s[0].Points, HasLen, 1)
// 	c.Assert(s[0].Points[0][1].(float64), Equals, float64(162597))

// 	client = self.server.GetClient("test2", c)
// 	s, err = client.Query("list series")
// 	c.Assert(err, IsNil)
// 	c.Assert(s, HasLen, 1)
// 	c.Assert(s[0].Points, HasLen, 1000)

// 	s, err = client.Query("select count(value) from /.*/")
// 	c.Assert(s, HasLen, 1000)
// 	for _, series := range s {
// 		c.Assert(series.Points, HasLen, 1)
// 		c.Assert(series.Points[0][1].(float64), Equals, float64(1434))
// 	}
// }
