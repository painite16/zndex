package code
import "testing"


type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "short"},
    {0,"zero"},
    {29,"long"},
    {202,"very long"}
  
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
//Дополните тест для функции Length из урока, чтобы покрытие кода составляло 100%.