package nonpersistent

import "net/http"

type RestService struct {
}

func (s *RestService) Init() (err error) {
	return
}

func (s *RestService) Port() (port int) {
	return 0
}

func (s *RestService) IsValid() (ret bool) {
	return false
}

func (s *RestService) StartServe() (err error) {
	dispatcher := s.NewDispatcher(10000, 10000)
	dispatcher.Run()
	return
}

// Job represents the job to be run
type Job struct {
	// Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			// case job := <-w.JobChannel:
			// 	we have received a work request.
			// 	if err := job.Payload.UploadToS3(); err != nil {
			// 		log.Errorf("Error uploading to S3: %s", err.Error())
			// 	}
			case <-w.JobChannel:

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the body into a string for json decoding
	// var content = &PayloadCollection{}
	// err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// Go through each payload and queue items individually to be posted to S3
	// for _, payload := range content.Payloads {

	// let's create a job with the payload
	// work := Job{Payload: payload}
	work := Job{}
	// Push the work onto the queue.
	JobQueue <- work
	// }

	w.WriteHeader(http.StatusOK)
}

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	MaxWorkers int
}

func (s *RestService) NewDispatcher(maxWorkers, maxQueue int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	JobQueue = make(chan Job, maxQueue)
	return &Dispatcher{MaxWorkers: maxWorkers, WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
