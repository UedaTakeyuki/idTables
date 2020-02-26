package main

import "fmt"

//	f "github.com/fauna/faunadb-go/faunadb"

//	"github.com/speps/go-hashids"

/*
func CreateHashID() (h *(hashids.HashID), err error) {
	// hashid
	hd := hashids.NewData()
	hd.Salt = "For Magic Label"
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hd.MinLength = 8
	h, err = hashids.NewWithData(hd)
	return
}
*/

// idTable
type IDTable struct {
	//	count    *PersistentInt
	table    map[string]string
	invtable map[string]string
	path     string
	//	h        *(hashids.HashID)
}

func NewIDTable(path string) (p *IDTable, err error) {
	p = new(IDTable)
	//	p.count = count
	p.path = path
	p.table = make(map[string]string)
	p.invtable = make(map[string]string)
	readStringDataMapCSV(path+".csv", &(p.table))
	readStringDataMapCSV(path+"_inv.csv", &(p.invtable))
	//	p.h, err = CreateHashID()
	return
}

func (t *IDTable) Update(internalID string, externalID string) (err error) {
	//	fmt.Println("bf update", t.table)
	(t.table)[externalID] = internalID
	(t.invtable)[internalID] = externalID
	//	fmt.Println("af update", t.table)
	//	t.count.Inc()
	go addStringDataMapCSV(t.path+".csv", externalID, internalID)     // add to CSV file
	go addStringDataMapCSV(t.path+"_inv.csv", internalID, externalID) // add to CSV file
	return
}

func (t *IDTable) DeleteDeleteByExternalID(externalID string) (internalID string, err error) {
	//	externalID := c.Param("id")
	fmt.Println(t)
	//	fmt.Println("bf delete", t.table)
	internalID = (t.table)[externalID]
	delete(t.table, externalID)
	delete(t.invtable, internalID)
	//	fmt.Println("af delete", t.table)
	overwriteStringDataMapCSV(t.path+".csv", &(t.table))        // save to CSV file
	overwriteStringDataMapCSV(t.path+"_inv.csv", &(t.invtable)) // save to CSV file
	//	faunaDelete(externalID)
	//	faunaDelete(internalID)
	return
}
