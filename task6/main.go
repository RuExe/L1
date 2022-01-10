package task6

//quit := make(chan struct{})
//go func() {
//	for {
//		select {
//		case <-quit:
//			return
//		default:
//			// …
//		}
//	}
//}()
//close(quit)

//func Generator() chan int {
//	ch := make(chan int)
//	go func() {
//		n := 1
//		for {
//			select {
//			case ch <- n:
//				n++
//			case <-ch:
//				return
//			}
//		}
//	}()
//	return ch
//}
//func main() {
//	number := Generator()
//	fmt.Println(<-number)
//	fmt.Println(<-number)
//	close(number)
//}

//forever := make(chan struct{})
//ctx, cancel := context.WithCancel(context.Background())
//go func(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():  // if cancel() execute
//			forever <- struct{}{}
//			return
//		default:
//			fmt.Println("for loop")
//		}
//
//		time.Sleep(500 * time.Millisecond)
//	}
//}(ctx)
//
//go func() {
//	time.Sleep(3 * time.Second)
//	cancel()
//}()
//
//<-forever
//fmt.Println("finish")
