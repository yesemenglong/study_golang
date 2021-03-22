package main

type sub struct {
	// 把chan error 看成一个整体，作为关闭的通道
	closing chan chan error
	updates chan string
}

func (s *sub) Close() error {
	// Tip 核心逻辑：两层通知，第一层作为准备关闭的通知，第二层作为关闭结果的返回
	errc := make(chan error)
	// Tip 第一步：要关闭时，先传一个chan error过去，通知要关闭了
	s.closing <- errc
	// Tip 第三步：从chan error中读取错误，阻塞等待
	return <-errc
}

func (s *sub) loop() {
	var err error
	for {
		select {
		case errc := <-s.closing:
			// Tip 第二步：收到关闭后，进行处理，处理后把error传回去
			errc <- err
			close(s.updates)
			return
		}
	}
}
