package pbstor

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"pbstor/pb"
	"testing"
)

func TestRWPb(t *testing.T) {
	e1 := &pb.Entry{Term:1,LogIdx:1,Key:"aaa",Value:"bbb"}
	e2 := & pb.Entry{Term:2,LogIdx:1,Key:"ccc",Value:"bbb"}
	entries := & pb.Entries{}
	entries.E = append(entries.E, e1)
	entries.E = append(entries.E, e2)

	data, _ := proto.Marshal(entries)
	ioutil.WriteFile("./log", data, os.ModePerm)

	data, _ = ioutil.ReadFile("./log")
	log := &pb.Entries{}

	proto.Unmarshal(data, log)

	for _, v := range log.E {
		fmt.Printf("term:%v, logIdx:%v, key:%v, value:%v\n", v.Term, v.LogIdx, v.Key, v.Value)

	}
}