# HTTP MD5 listing
This application lists MD5 sums of responses from HTTP Requests. 

# Notes
* App will return MD5 sums from http Response Body
* If http request fails it will show the MD5 sum of the Error
* Program will fail fatally if the url's sent are not parseable
* Program will get parallel flag to determine max concurent HTTP calls if flag is not provided it will default to 10

# Usage

./HTTPMD5.exe -paralel 5 www.google.com www.amazon.com http://www.yahoo.com https://stackoverflow.com reddit.com