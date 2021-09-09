### csvToExcel
this is a tool to convert many csv files to one excel file.

### set up
create file like this.
```
├── README.md
├── csv       <- csv files 
├── go.mod
├── go.sum
├── main.go
└── output    <- created excel files
```

### RUN 
```
go run main.go
```
or

```
go build && ./csvToExcel
```
>※ csv file name is to be sheet name. 