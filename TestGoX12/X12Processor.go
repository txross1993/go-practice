package TestGoX12

import (
	"github.com/azoner/gox12"
	log "github.com/txross1993/go-practice/EdiParser/logwrapper"
	"strings"
)

type PreProcessor struct {
	fr *rawX12FileReader
}

func NewFileReader(in io.Reader) (*PreProcessor, error) {
	r, err := gox12.NewRawX12FileReader(in)

	if err !:= nil (
		return nil, err
	)

	p := &PreProcessor{p: &r, ""}

	return &p, nil

}

func getSegmentId(segmentId string, elemId int, subelemId int) fieldName string {
	if elemId < 10 {
		fieldName := segmentId + "_0" + strings.Iota(elemId)
		return fieldName
	}

	fieldName := segmentId + "_" + strings.Iota(elemId)
}

func parseTransaction(ch RawSegment) *pb.Transaction {

	
	txn := &pb.Transaction{
		txnHeader: 
		txn:
		txnFooter:
	}
}

func (p *PreProcessor) GetTransactions() {
	for rs := range p.fr.GetSegments()  {
		if strings.ToUpper(rs.Segment.SegmentId) == "ISA" {
			for v := range rs.Segment.GetAllValues()
				if v.ElementIdx == 13 {
					p.isa_13 = v.Value
				}

				if v.ElementIdx < 10 {
					id := "0" + strings.Itoa(v.ElementIdx)
				}
		}


	}
}
