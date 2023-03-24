<h1>Stupid Go Pool Worker</h1>

<h2>How to build a pool worker</h2>

```bash
1.  Create a type of Job and Result

2.  Create a func worker(jobs <-chan Job, results chan<- Result) {}
    chan<- is for sending value to this chan (sender)
    <-chan is for sending chan to some value (receiver)

3.  Create chan for jobs and results

4.  Create number of workers by using for-loop to do work and in this for loop must call the func worker()

5.  Assign a value into jobs chan and dont forget to close the chan when finish

6.  for-loop in range of jobs size and create some var to receive a result from results chan (or you can stack here)
```