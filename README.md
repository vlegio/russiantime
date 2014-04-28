[![GoDoc](https://godoc.org/github.com/t0pep0/russiantime?status.png)](https://godoc.org/github.com/t0pep0/russiantime)
#Russiantime#

Golang package for format time in Russain style

##Usage:##


import library:

    import "github.com/t0pep0/russiantime"

And call:

    russiantime.Russian(time.Now(), "%YYYY %YYY %YY %Y, %md %Md %m %M, %D %d, %Wd %w %W, %H %h, %N %n, %S %s")
        =>  2014 014 14 4, 4 04 Апр Апрель, 02 2, 3 Ср Среда, 14 14, 00 0, 56 56
        
##Package files:##

* russiantime.go
   
    > func Russian(time.Time, string) string
* const.go
    
    > Internal data


##Format string syntax:##
> %YYYY - full digit year (e.g. 2014)

> %YYY - thre digit year (e.g. 014)

> %YY - two digit year (e.g. 14)

> %Y - one digit year (e.g. 4)

> %md - month digit (e.g. 4)

> %Md - two digit month (e.g. 04)

> %m - short month name (e.g. Апр)

> %M - long month name (e.g. Апрель)

> %D - two digit day (e.g. 02)

> %d - digit day (e.g. 2)

> %Wd - digit day of week (e.g. 3)

> %w - short day of week name (e.g. Ср)

> %W - long day of week name (e.g. Среда)

> %H - two digit hour (e.g. 05)

> %h - digit hour (e.g. 5)

> %N - two digit minutes (e.g. 03)

> %n - digit minutes (e.g. 3)

> %S - two digit seconds (e.g. 04)

> %s - digit seconds (e.g. 4)
