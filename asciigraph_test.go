package asciigraph

import "testing"

func TestPlot(t *testing.T) {
	type input struct {
		data []float64
		conf map[string]interface{}
	}

	cases := []struct {
		in   input
		want string
	}{
		{
			input{[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 1}, map[string]interface{}{}},

			`   11.00 ┤      ╭╮   
   10.00 ┤      ││   
    9.00 ┼      ││   
    8.00 ┤      ││   
    7.00 ┤     ╭╯│╭╮ 
    6.00 ┤     │ │││ 
    5.00 ┤    ╭╯ │││ 
    4.00 ┤    │  │││ 
    3.00 ┤    │  ╰╯│ 
    2.00 ┼╮ ╭╮│    │ 
    1.00 ┤╰─╯││    ╰ 
    0.00 ┤   ││      
   -1.00 ┤   ││      
   -2.00 ┤   ╰╯      `},
		{
			input{
				[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 4, 5, 6, 9, 4, 0, 6, 1, 5, 3, 6, 2},
				map[string]interface{}{"caption": "Plot using asciigraph"}},

			`   11.00 ┤      ╭╮              
   10.00 ┤      ││              
    9.00 ┼      ││    ╭╮        
    8.00 ┤      ││    ││        
    7.00 ┤     ╭╯│╭╮  ││        
    6.00 ┤     │ │││ ╭╯│ ╭╮  ╭╮ 
    5.00 ┤    ╭╯ │││╭╯ │ ││╭╮││ 
    4.00 ┤    │  ││╰╯  ╰╮││││││ 
    3.00 ┤    │  ╰╯     ││││╰╯│ 
    2.00 ┼╮ ╭╮│         ││││  ╰ 
    1.00 ┤╰─╯││         ││╰╯    
    0.00 ┤   ││         ╰╯      
   -1.00 ┤   ││                 
   -2.00 ┤   ╰╯                 
           Plot using asciigraph`},
		{
			input{
				[]float64{.2, .1, .2, 2, -.9, .7, .91, .3, .7, .4, .5},
				map[string]interface{}{"caption": "Plot using asciigraph"}},

			`    2.00 ┤           
    1.03 ┼  ╭╮ ╭╮    
    0.07 ┼──╯│╭╯╰─── 
   -0.90 ┤   ╰╯      
           Plot using asciigraph`},
		{
			input{
				[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 1},
				map[string]interface{}{"height": 4}},

			`   11.00 ┤           
    8.40 ┼      ╭╮   
    5.80 ┤    ╭─╯│╭╮ 
    3.20 ┤╮ ╭╮│  ╰╯│ 
    0.60 ┼╰─╯││    ╰ 
   -2.00 ┤   ╰╯      `},
	}

	for _, c := range cases {
		got := Plot(c.in.data, c.in.conf)
		if got != c.want {
			t.Errorf("Plot(%f, %q) == %q, want %q", c.in.data, c.in.conf, got, c.want)
		}
	}
}
