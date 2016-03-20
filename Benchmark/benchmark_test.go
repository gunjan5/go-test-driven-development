package bench

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkSample(b *testing.B) {
	temp, _ := template.New("").Parse("Gopher")

	b.ResetTimer()
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		temp.Execute(&buf, nil)
		buf.Reset()
	}

}
