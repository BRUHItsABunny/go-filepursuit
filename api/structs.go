package api

import (
	"net/url"
	"strconv"
)

//go:generate go-enum -f=$GOFILE --nocase --flag --names

// Make x ENUM(DateDESC, DateASC, SizeDESC, SizeASC, FileDESC, FileASC)
type SortType int

func (t SortType) ToParameter() string {
	switch t {
	case SortTypeDateDESC:
		return "datedesc"
	case SortTypeDateASC:
		return "dateasc"
	case SortTypeSizeDESC:
		return "sizedesc"
	case SortTypeSizeASC:
		return "sizeasc"
	case SortTypeFileDESC:
		return "filedesc"
	case SortTypeFileASC:
		return "fileasc"
	}
	return "datedesc"
}

// Make x ENUM(All, Archive, Mobile, Ebook, Audio, Video)
type FileType int

func (t FileType) ToParameter() string {
	switch t {
	case FileTypeAll:
		return "all"
	case FileTypeArchive:
		return "archive"
	case FileTypeMobile:
		return "mobile"
	case FileTypeEbook:
		return "ebook"
	case FileTypeAudio:
		return "audio"
	case FileTypeVideo:
		return "video"
	}
	return "all"
}

type Parameters struct {
	Sort     SortType
	FileType FileType
	Start    int
}

func NewParameters(sortType SortType, fileType FileType, start int) *Parameters {
	return &Parameters{
		Sort:     sortType,
		FileType: fileType,
		Start:    start,
	}
}

func (p *Parameters) Encode() url.Values {
	return url.Values{
		"start":    {strconv.Itoa(p.Start)},
		"sort":     {p.Sort.ToParameter()},
		"type":     {p.FileType.ToParameter()},
		"filetype": {""}, // I didn't make this up, the app does this
		"device":   {"mobile"},
	}
}

type FPResponse struct {
	Status     string     `json:"status"`
	Result     []*FPEntry `json:"result,omitempty"`
	FilesFound []*FPEntry `json:"files_found,omitempty"`
}

type FPEntry struct {
	Type          string `json:"type"`
	DomainName    string `json:"domain_name,omitempty"`
	DomainLink    string `json:"domain_link,omitempty"`
	FilesInside   int64  `json:"files_inside,omitempty"`
	FolderName    string `json:"folder_name,omitempty"`
	FolderLink    string `json:"folder_link,omitempty"`
	FileID        string `json:"file_id,omitempty"`
	FileType      string `json:"file_type,omitempty"`
	FileName      string `json:"file_name,omitempty"`
	FileLink      string `json:"file_link,omitempty"`
	DateAdded     string `json:"date_added,omitempty"`
	TimeAgo       string `json:"time_ago,omitempty"`
	FileSize      string `json:"file_size,omitempty"`
	FileSizeBytes string `json:"file_size_bytes,omitempty"`
	ReferrerLink  string `json:"referrer_link,omitempty"`
	ReferrerHost  string `json:"referrer_host,omitempty"`
	ReadablePath  string `json:"readable_path,omitempty"`
}
