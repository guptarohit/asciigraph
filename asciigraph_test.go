package asciigraph

import (
	"testing"
	"fmt"
)

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

			` 11.00 ┤      ╭╮   
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
				map[string]interface{}{"caption": "Plot using asciigraph."}},

			` 11.00 ┤      ╭╮              
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
          Plot using asciigraph.`},
		{
			input{
				[]float64{.2, .1, .2, 2, -.9, .7, .91, .3, .7, .4, .5},
				map[string]interface{}{"caption": "Plot using asciigraph."}},

			`  2.00 ┤           
  1.03 ┼  ╭╮ ╭╮    
  0.07 ┼──╯│╭╯╰─── 
 -0.90 ┤   ╰╯      
          Plot using asciigraph.`},
		{
			input{
				[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 1},
				map[string]interface{}{"height": 4, "offset": 3}},

			` 11.00 ┤           
  8.40 ┼      ╭╮   
  5.80 ┤    ╭─╯│╭╮ 
  3.20 ┤╮ ╭╮│  ╰╯│ 
  0.60 ┼╰─╯││    ╰ 
 -2.00 ┤   ╰╯      `},
		{
			input{
				[]float64{.453, .141, .951, .251, .223, .581, .771, .191, .393, .617, .478},
				map[string]interface{}{}},

			` 0.95 ┤           
 0.86 ┤ ╭╮        
 0.77 ┤ ││  ╭╮    
 0.68 ┤ ││  ││    
 0.59 ┤ ││ ╭╯│ ╭╮ 
 0.50 ┤ ││ │ │ │╰ 
 0.41 ┼╮││ │ │╭╯  
 0.32 ┤│││ │ ││   
 0.23 ┤││╰─╯ ╰╯   
 0.14 ┤╰╯         `},


		{
			input{[]float64{.01, .004, .003, .0042, .0083, .0033, 0.0079},
				map[string]interface{}{}},

			` 0.010 ┼╮      
 0.009 ┤│      
 0.008 ┤│  ╭╮╭ 
 0.007 ┤│  │││ 
 0.006 ┤│  │││ 
 0.005 ┤│  │││ 
 0.004 ┤╰╮╭╯││ 
 0.003 ┤ ╰╯ ╰╯ `},

		{
			input{
				[]float64{192, 431, 112, 449, -122, 375, 782, 123, 911, 1711, 172},
				map[string]interface{}{"height": 10}},

			` 1711 ┤           
 1544 ┼        ╭╮ 
 1378 ┤        ││ 
 1211 ┤        ││ 
 1044 ┤        ││ 
  878 ┤       ╭╯│ 
  711 ┤     ╭╮│ │ 
  545 ┤     │││ │ 
  378 ┤╭╮╭╮╭╯││ │ 
  211 ┼╯╰╯││ ╰╯ ╰ 
   45 ┤   ││      
 -122 ┤   ╰╯      `},
		{
			input{
				[]float64{0.3189989805, 0.149949026, 0.30142492354, 0.195129182935, 0.3142492354, 0.1674974513, 0.3142492354, 0.1474974513, 0.3047974513},
				map[string]interface{}{"width": 30, "height": 5, "caption": "Plot with custom height & width."}},

			` 0.32 ┤                              
 0.29 ┼╮            ╭─╮     ╭╮     ╭ 
 0.27 ┤╰╮    ╭─╮   ╭╯ │    ╭╯│     │ 
 0.24 ┤ │   ╭╯ ╰╮ ╭╯  ╰╮  ╭╯ ╰╮   ╭╯ 
 0.22 ┤ ╰╮ ╭╯   ╰╮│    ╰╮╭╯   ╰╮ ╭╯  
 0.19 ┤  ╰╮│     ╰╯     ╰╯     │╭╯   
 0.16 ┤   ╰╯                   ╰╯    
         Plot with custom height & width.`},
	}

	for _, c := range cases {
		got := Plot(c.in.data, c.in.conf)
		fmt.Println(got + "\n")
		if got != c.want {
			t.Errorf("Plot(%f, %q) == %q, want %q", c.in.data, c.in.conf, got, c.want)
		}
	}
}
