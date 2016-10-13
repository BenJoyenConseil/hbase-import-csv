package main

import (
	"github.com/tsuna/gohbase"
	"encoding/csv"
	"bufio"
	"io"
	"strconv"
	"time"
	"github.com/golang/example/stringutil"
	"github.com/tsuna/gohbase/hrpc"
	"fmt"
	"context"
	"os"
)

func main() {
		f, _ := os.Open("MOCK_DATA.csv")
	client := gohbase.NewClient("localhost:2181")


	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// clientid, crid, eventid, typeevent, date, phrase, raw
		epoch, _ := strconv.Atoi(record[4])
		dateTime := time.Unix(int64(epoch), 0)
		dateFormatted, _ := strconv.Atoi(dateTime.Format("20060102"))
		invertEpoch := strconv.Itoa(99999999 - dateFormatted)
		//fmt.Println(invertEpoch)
		rowkey := stringutil.Reverse(record[0]) + record[1] + record[2] + record[3] + invertEpoch

		getRequest, err := hrpc.NewGetStr(context.Background(), "ldt_table", "r1")
		getRsp, err := client.Get(getRequest)
		fmt.Println(getRsp)

		values := map[string]map[string][]byte{"cf": map[string][]byte{"phrase": []byte(record[5]), "raw": []byte(record[6])}}
		putRequest, err := hrpc.NewPutStr(context.Background(), "ldt_table", rowkey, values)
		_, err = client.Put(putRequest)

		if err != nil {
			fmt.Println(err)
		}
	}
}
