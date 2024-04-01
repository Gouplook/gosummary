package main

import "fmt"

type Job struct {
	ID   int
	Data []int
}

type Result struct {
	JobID  int
	Result int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		// 执行任务逻辑，这里仅仅是一个示例
		sum := 0
		for _, v := range job.Data {
			sum += v
		}

		// 将结果发送到结果通道
		result := Result{
			JobID:  job.ID,
			Result: sum,
		}
		results <- result
	}
}

//numJobs := 10
//numWorkers := 3
//
//// 创建任务和结果通道 10个人
//jobs := make(chan Job, numJobs)       // 10任务
//results := make(chan Result, numJobs) // 需要10次结果
//
//// 创建等待组
//var wg sync.WaitGroup
//
//// 启动协程池中的工作协程
//for i := 0; i < numWorkers; i++ { // 开启3个携程
//	wg.Add(1)
//	go func(workerID int) {
//		defer wg.Done()
//		worker(workerID, jobs, results)
//	}(i)
//}
//
//// 发送任务到任务通道
//for i := 0; i < numJobs; i++ {
//	job := Job{
//		ID:   i,
//		Data: []int{1, 2, 3, 4, 5},
//	}
//	jobs <- job
//}
//close(jobs) // 关闭任务通道
//
//// 等待所有工作协程完成任务
//go func() {
//	wg.Wait()
//	close(results) // 关闭结果通道
//}()
//
//// 读取并打印任务结果
//for result := range results {
//	fmt.Printf("Job %d Result: %d\n", result.JobID, result.Result)
//}

func main() {
	//var newPrimSum, newInteSum, newSenSum, newSum int
	var upInteSum, upSenSum, upSum int
	upSum = 3
	upSenSum = 0
	upInteSum = 3
	var msg string
	switch {
	case upInteSum != 0 && upSenSum == 0:
		msg = fmt.Sprintf("升级报警: %v处(中级%v处)", upSum, upInteSum)
	case upInteSum == 0 && upSenSum != 0:
		msg = fmt.Sprintf("升级报警: %v处(高级%v处)", upSum, upSenSum)
	case upInteSum != 0 && upSenSum != 0:
		msg = fmt.Sprintf("升级报警: %v处(中级%v处、高级%v处)", upSum, upInteSum, upSenSum)
	}

	fmt.Println(msg)
}
