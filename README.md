# go-filepursuit
Filepursuit API Golang practice project based on PRO version of the FilePursuit APK

### Note about discover function
* Use query `""` to get results that would resemble the main page in the app
* Use the link of an entry you want to discover further
* If fileType = `"-"` it is a folder that can be further explored

### Note to API owner
This API is in dire need of cryptographic securing, while reversing this API I noticed a surprising lack of any kind of security.
If you want some pointers from someone who is into reverse engineering as to how to properly secure the API, just get in contact with me and I will lay out some ideas that would be compatible with your PHP server-side and java client-side.

### Future of this project
I don't plan on maintaining this project. 

Use at your own risk.

### Usage
You can install this library by using:
```
go get github.com/BRUHItsABunny/go-filepursuit
```
This library depends on my HTTP client abstraction (also not production ready yet) :
```
go get github.com/BRUHItsABunny/gOkHttp
```

You can find examples in the `_examples` folder.

This library can:
* Search for files
* Discover servers and files
* Submit files
* Get app news
* Suggest other terms to search for based on incomplete keyword