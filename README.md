# Go File Pursuit
[FilePursuit](https://filepursuit.com/) API Golang practice project based on [PRO version](https://play.google.com/store/apps/details?id=com.filepursuit.filepursuitpro) of the [FilePursuit APK](https://play.google.com/store/apps/details?id=com.filepursuit.filepursuit)

### Note about discover function
* Use query `""` to get results that would resemble the main page in the app
* Use the link of an entry you want to discover further

### Note to API owner
This API is **still** in dire need of cryptographic securing, while reversing this API I noticed a surprising lack of any kind of security other than the built-in security measures that come with the Flutter SDK (eg: Flutter only uses its own CA keystore instead of the system one).
If you want some pointers from someone who is into reverse engineering as to how to properly secure the API, just get in contact with me, and I will lay out some ideas that would be compatible with your PHP server-side and ~~java~~ Flutter client-side.

### Future of this project
I don't plan on maintaining this project. 

This code is able to work `with` **and** `without` API key.

Use at your own risk without a valid API key, you can get one [here](https://rapidapi.com/azharxes/api/filepursuit/).

### Usage
You can install this library by using:
```bash
go get github.com/BRUHItsABunny/go-filepursuit
```

Refer to [lib_test.go](https://github.com/BRUHItsABunny/go-filepursuit/blob/main/lib_test.go) to see how to use the library

This library can:
* Search for files
* Discover servers and files
* Suggest other terms to search for based on incomplete keyword