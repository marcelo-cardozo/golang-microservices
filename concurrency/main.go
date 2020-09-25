package main

import (
	"fmt"
	"github.com/marcelo-cardozo/golang-microservices/src/api/domain/repositories"
	"github.com/marcelo-cardozo/golang-microservices/src/api/services"
	"github.com/marcelo-cardozo/golang-microservices/src/api/utils"
	"math/rand"
	"strconv"
	"sync"
)

type createRepoResult struct {
	request repositories.CreateRepoRequest
	response *repositories.CreateRepoResponse
	err utils.ApiError
}

var (
	success = map[string]string{}
	failed  = map[string]utils.ApiError{}
)


func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)
	for i := 0; i < 20; i++ {
		request := repositories.CreateRepoRequest{Name: strconv.Itoa(rand.Int())}

		result = append(result, request)
	}
	return result
}

func main() {
	requests := getRequests()
	input := make(chan createRepoResult)
	buffer := make(chan struct{}, 5)

	wg := &sync.WaitGroup{}
	go handleResult(input, wg)

	for _, request := range requests {
		fmt.Println(request.Name)
		wg.Add(1)
		buffer <- struct{}{}
		go createRepo(request, input, buffer)
	}
	wg.Wait()
	close(input)
}

func handleResult(input chan createRepoResult, wg *sync.WaitGroup){
	for result := range input {
		if result.err != nil {
			failed[result.request.Name]=result.err
		}else {
			success[result.request.Name]=result.response.Name
		}
		fmt.Println("*****"+result.request.Name)
		wg.Done()
	}
}

func createRepo(request repositories.CreateRepoRequest, input chan createRepoResult, buffer chan struct{}){
	fmt.Println("-----"+request.Name)

	response, err := services.RepoService.CreateRepo(request)

	input <- createRepoResult{
		request: request,
		response: response,
		err:      err,
	}

	<- buffer
}