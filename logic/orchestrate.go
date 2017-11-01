package logic


func StartPipeline(maxChannelSize int, dop int) {
	raw_data := make(chan string, maxChannelSize)
	normalized_data := make(chan string, maxChannelSize)

	route := CreateRoute(normalized_data)
	transport := CreateTransport(raw_data)
	normalizer := CreateNormalizer(raw_data, normalized_data)

	go route()
	go transport()
	startWorkers(normalizer, dop)
}

func StopPipeline(){

}

func startWorkers(f func(), numberOfWorkers int){
	for i := 0; i < numberOfWorkers; i++ {
		go f()
	}
}
