package main

const KeyID = "123456"

func main() {

}

/*
// WaitGroup的用法
func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Task 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Task 2")
	}()

	wg.Wait()
	fmt.Println("Over!")
}

*/

/*
// WithTimeout,请求超时自动取消
func main() {
	parent := context.Background()
	ctx, _ := context.WithTimeout(parent, 1*time.Millisecond)

	start := time.Now()
	req, _ := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err, " time ", time.Now().Sub(start))
		return
	}
	fmt.Println("Response received, status code:", res.StatusCode)
}
*/

/*
func main() {
	exit := make(chan string)
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	go worker(ctx, "01", exit)
	go worker(ctx, "02", exit)
	go worker(ctx, "03", exit)
	time.Sleep(time.Second * 5)
	cancel()
	<-exit
}

// CancelCtx取消同级及其下级的goroutine
func worker(ctx context.Context, name string, exit chan string) {
	for {
		select {
		case <-ctx.Done():
			defer close(exit)
			fmt.Printf("%s closer.\n", name)
			return
		default:
			fmt.Printf("%s is working.\n", name)
			time.Sleep(time.Second * 1)
		}
	}
}
*/

// ValueContext值传递
/*
func ContextWithValue() {
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), KeyID, rand.Int())
	operation1(ctx)
}

func operation1(ctx context.Context) {
	fmt.Println("operation1 for id:", ctx.Value(KeyID), " completed")
	operation2(ctx)
}

func operation2(ctx context.Context) {
	fmt.Println("operation2 for id:", ctx.Value(KeyID), " completed")
}

*/
