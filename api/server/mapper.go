package server

type Mapper interface {
	Single(interface{}) interface{}
	Many(interface{}) interface{}
}
