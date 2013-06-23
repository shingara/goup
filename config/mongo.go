package config

import (
  "labix.org/v2/mgo"
  "fmt"
)

var(
  base_config *Config
)

type Config struct {
  session *mgo.Session
  database *mgo.Database
  Collection *mgo.Collection
}

// Get a mongoDB session to this app
func init(){
  fmt.Printf("init Mongo START\n")
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  /* defer session.Close() */
  // Optional. Switch the session to a monotonic behavior.
  base_config = &Config{
    session: session,
  }
  fmt.Printf("init Mongo DONE\n")
}

func CloseSession() {
  base_config.Close()
}

func (c *Config) Close() {
  c.session.Close()
}

func (c *Config) Clone() *Config{
  config_clone := &Config{
    session: c.session.Clone(),
    database: c.session.DB("goup"),
  }
  config_clone.session.SetMode(mgo.Monotonic, true)
  return config_clone
}

func Config_collection(name string) *Config{
  c := base_config.Clone()
  c.Collection = c.database.C(name)
  return c
}
