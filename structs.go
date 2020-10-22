package go_filepursuit

import gokhttp "github.com/BRUHItsABunny/gOkHttp"

type FilePursuitClient struct {
	Client *gokhttp.HttpClient
}

type FileEntry struct {
	ID       string `json:"id"`
	Link     string `json:"link"`
	File     string `json:"file"`
	Filetype string `json:"filetype"`
	RegDate  string `json:"reg_date"`
	Filename string `json:"filename"`
	Filesize string `json:"filesize"`
	Referrer string `json:"referrer"`
	Linkname string `json:"linkname"`
	Hostname string `json:"hostname"`
}

type NewsItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Link      string `json:"link"`
	Imagelink string `json:"imglink"`
	Button    string `json:"button"`
}

type DiscoveryEntry struct {
	Link        string  `json:"link"`
	FilesInside *int64  `json:"files_inside,omitempty"`
	File        string  `json:"file"`
	RegDate     string  `json:"reg_date"`
	Filesize    string  `json:"filesize"`
	ID          string  `json:"id"`
	FileType    string  `json:"filetype"`
	Hostname    string  `json:"hostname"`
	Filename    *string `json:"filename,omitempty"`
	Referrer    *string `json:"referrer,omitempty"`
	Linkname    *string `json:"linkname,omitempty"`
}

func (de *DiscoveryEntry) IsDirectory() bool {
	return de.FileType == "-"
}
