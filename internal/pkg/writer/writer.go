package writer

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/letscloud-community/letscloud-go/domains"
)

//Writer represents tabbed writer
type Writer struct {
	wr *tabwriter.Writer
}

//New instantiates a new writer object
func New(out io.Writer) *Writer {
	writer := tabwriter.NewWriter(out, 0, 8, 1, '\t', 0)
	return &Writer{wr: writer}
}

//Flush flushes the buffered content to io
func (w *Writer) Flush() error {
	return w.wr.Flush()
}

//WriteHeader writes header/struct field names
func (w *Writer) WriteHeader(vals ...string) {
	var s string
	for i, v := range vals {
		s += fmt.Sprintf("%s", v)

		if i != len(vals)-1 {
			s += "\t"
		} else {
			s += "\n"
		}
	}

	w.wr.Write([]byte(s))
}

//WriteData writes data to writer
func (w *Writer) WriteData(vals ...interface{}) {
	var s string
	for i, v := range vals {
		if dt, ok := v.([]domains.IPAddress); ok {
			var ipaddr string
			for i, ip := range dt {
				ipaddr += ip.Address

				if i != len(dt)-1 {
					ipaddr += ","
				}
			}

			s += fmt.Sprintf("%v", ipaddr)
		} else {
			if v != "" {
				s += fmt.Sprintf("%v", v)
			} else {
				s += fmt.Sprintf("n/a")
			}
		}

		if i != len(vals)-1 {
			s += "\t"
		} else {
			s += "\n"
		}
	}

	w.wr.Write([]byte(s))
}
