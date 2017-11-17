Golang data structure
==========================
1.HashSet
<pre><code>
package main

import (
	"github.com/mhcvs2/godatastructure"
	"fmt"
)

func main() {
	exampleSet := godatastructure.NewHashSet()
	exampleSet.Add("a")
	exampleSet.Add("b")
	exampleSet.Add("c")
	fmt.Println(exampleSet.String())
	exampleSet.Remove("a")
	exampleSet.Add("b")
	fmt.Println(exampleSet.String())
}

//Set{a b c}
//Set{b c}
</pre></code>