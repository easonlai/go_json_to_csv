# Parsing JSON to CSV in Golang

This is a demo repo to demonstrate how to use Golang to parse JSON files into CSV file. I also used [Golang](https://go.dev/) to generate the volume of JSON files (10,000). It read 10,000 of rows data entries from a CSV file and created corresponding JSON files. I also prepared a [Python](https://www.python.org/) version to parse JSON files into CSV file for performance comparison.

<img src="https://github.com/easonlai/go_json_to_csv/blob/main/git-images/golang-gopher.png" width="300" height="300">

**Content:**
* /data/ <-- Folder to store input JSON files.
* to_json.csv <-- Source CSV files for JSON files generation.
* csv_to_json.go <-- Golang progrm to read 10,000 of rows data entries from a CSV file and created corresponding JSON files. It can directly run by "go run .\csv_to_json.go". Or build this Golang program by "go build .\csv_to_json.go" as exe to run.
* json_to_csv.go <-- Golang program to parse JSON files into CSV file. It can directly run by "go run .\json_to_csv.go". Or build this Golang program by "go build .\json_to_csv.go" as exe to run.
* json_to_csv.py <-- Python script to parse JSON files into CSV file.
* json_to_csv_parallel.py <-- Optimized Python script to parse JSON file into CSV file with [multiprocessing pool](https://docs.python.org/3/library/multiprocessing.html).


**Performance comparison with Golang and Python**

On average, compiled Golang program is 4 times faster than the Python script.
![alt text](https://github.com/easonlai/go_json_to_csv/blob/main/git-images/git-image-1.png)

**Performance reference of optimized Python with multiprocessing pool**

On average, optimized Python with a multiprocessing pool is 4 times faster than script without optimization. It is almost on par with compiled Golang program.
![alt text](https://github.com/easonlai/go_json_to_csv/blob/main/git-images/git-image-2.png)


**Enjoy!**
