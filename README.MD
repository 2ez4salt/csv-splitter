# CSV SPLITTER
The purpose of this project is to split a **large** CSV file into smaller pieces.

# Specs

- That tool can split CSV file into smalller pieces.
- Encrypts the desired column as desired.
- Can determine the sizes of small parts.
- Can determine where to extract small parts.
- Can identify the names of small parts to be extracted.
- Can identify the passwords of small parts to be extracted.

# Installation
After cloned this repo you need some requirements.
```bash 
go get  
``` 
## Peer Dependencies
<h5><i>IMPORTANT! You need install this package.</i></h5>

```
bash npm i -g secure-spreadsheet  
``` 
# Encryption

This tool uses  [hashids](https://hashids.org/), for custom encryption settings you must set enviroment variables.


> "U-" prefix is automaticly adding to the encrypted area for now.

### Bash
```bash 
echo 'export HASHID_SALT=12345' >> ~/.bash_profile  
echo 'export HASHID_ALPHABET=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890' >> ~/.bash_profile  
echo 'export HASHID_MINLENGTH=9' >> ~/.bash_profile  
```  
### Zsh

```bash 
echo 'export HASHID_SALT=SALT' >> ~/.zshenv  
echo 'export HASHID_ALPHABET=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890' >> ~/.zshenv  
echo 'export HASHID_MINLENGTH=9' >> ~/.zshenv  
```   


# Usage
After that installations in project you can run with

```bash 
go run index.go  
```

# Change Log
Change log will be here !