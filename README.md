<div align="center">
<h1>GoGet</h1>
Easy way to hold a http(s) request<br>
</div>

```powershell
goget -u https://www.google.com
```

# TODO
**Goget will send a http request, and show the request time, status, response, and save response to a file**  

# Feature
✔ CLI Tool  
✔ http and https request  
✔ Save response to file

# HOW TO USE
First clone the file from github:  
```powershell
git clone https://github.com/banahaker/Goget.git
```
Go into the file and run go install:
```powershell
cd Goget
go install
```
Finally you can use Goget now:
```powershell
goget -u https://www.google.com -loc index.html
```
# Explain
**Flags:**
1. **-u** : (must) enter the url you want to request
2. **-loc**: (optional) enter a filename(like index.html, goget will create a data folder and create index.html inside, then the response will write into this file).

# Update Log
## 2022.02.09
 - First version public