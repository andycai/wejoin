package body

type Group struct {
	ID     uint   `json:"id" xml:"id" form:"id"`
	Name   string `json:"name" xml:"name" form:"name"`
	Logo   string `json:"logo" xml:"logo" form:"logo"`
	Notice string `json:"notice" xml:"notice" form:"notice"`
	Addr   string `json:"addr" xml:"addr" form:"addr"`
	Action uint   `json:"action" xml:"action" form:"action"`
}
