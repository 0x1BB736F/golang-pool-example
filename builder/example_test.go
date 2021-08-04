package pool

import (
	"context"
	"fmt"
	"time"
)

func ExampleNew() {
	// Create new strings.Builder pool with 10 items
	examplePool := New(10)

	// If all the elements of the pool are involved and
	// there is no time to wait, then you need to move on.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	item, err := examplePool.Get(ctx)
	if err != nil {
		// if err is not nil it will be an ErrTimeoutDone
		return
	}

	defer item.Close()

	_, _ = item.WriteString("hello")
	fmt.Println(item.String())

	//output: hello
}
