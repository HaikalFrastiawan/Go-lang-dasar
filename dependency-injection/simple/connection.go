package simple

import "fmt"

type Connection struct {
	*File
}



func (c *Connection) Close() {
	fmt.Println("CLose COnnection", c.File.Name)
}

func NewConnection(file *File) (*Connection, func()){
	Connection := &Connection{File: file}
	return	Connection, func() {
		Connection.Close()
	}
}