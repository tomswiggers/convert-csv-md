build:
				go build convert-csv-md.go
				chmod u+x convert-csv-md
install:
				go build convert-csv-md.go
				chmod u+x convert-csv-md
				cp convert-csv-md /usr/local/bin/convert-csv-md
