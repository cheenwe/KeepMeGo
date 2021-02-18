package upload

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
)

// FileHeader 解析多个文件上传中，每个具体的文件的信息
type FileHeader struct{
	ContentDisposition string
	Name string
	FileName string			///< 文件名
	ContentType string
	ContentLength int64
}

// ParseFileHeader 解析描述文件信息的头部
// @return FileHeader 文件名等信息的结构体
// @return bool 解析成功还是失败
func ParseFileHeader(h []byte) (FileHeader, bool){
	arr := bytes.Split(h, []byte("\r\n"))
	var outHeader FileHeader
	outHeader.ContentLength = -1
	const (
		CONTENTDISPOSITION = "Content-Disposition: "
		NAME = "name=\""
		FILENAME = "filename=\""
		CONTENTTYPE = "Content-Type: "
		CONTENTLENGTH = "Content-Length: "
	)
	for _,item := range arr{
		if bytes.HasPrefix(item, []byte(CONTENTDISPOSITION)){
			l := len(CONTENTDISPOSITION)
			arr1 := bytes.Split(item[l:], []byte("; "))
			outHeader.ContentDisposition = string(arr1[0])
			if bytes.HasPrefix(arr1[1], []byte(NAME)){
				outHeader.Name = string(arr1[1][len(NAME):len(arr1[1])-1])
			}
			l = len(arr1[2])
			if bytes.HasPrefix(arr1[2], []byte(FILENAME)) && arr1[2][l-1]==0x22{
				outHeader.FileName = string(arr1[2][len(FILENAME):l-1])
			}
		} else if bytes.HasPrefix(item, []byte(CONTENTTYPE)){
			l := len(CONTENTTYPE)
			outHeader.ContentType = string(item[l:])
		} else if bytes.HasPrefix(item, []byte(CONTENTLENGTH)){
			l := len(CONTENTLENGTH)
			s := string(item[l:])
			contentLength,err := strconv.ParseInt(s, 10, 64)
			if err!=nil{
				log.Printf("content length error:%s", string(item))
				return outHeader, false
			} else {
				outHeader.ContentLength = contentLength
			}
		} else {
			log.Printf("unknown:%s\n", string(item))
		}
	}
	if len(outHeader.FileName)==0{
		return outHeader,false
	}
	return outHeader,true
}

// ReadToBoundary 从流中一直读到文件的末位
/// @return []byte 没有写到文件且又属于下一个文件的数据
/// @return bool 是否已经读到流的末位了
/// @return error 是否发生错误
func ReadToBoundary(boundary []byte, stream io.ReadCloser, target io.WriteCloser)([]byte, bool, error){
	readData := make([]byte, 1024*8)
	readDataLen := 0
	buf := make([]byte, 1024*4)
	bLen := len(boundary)
	reachEnd := false
	for ;!reachEnd; {
		readLen, err := stream.Read(buf)
		if err != nil {
			if err != io.EOF && readLen<=0 {
				return nil, true, err
			}
			reachEnd = true
		}
		//todo: 下面这一句很蠢，值得优化
		copy(readData[readDataLen:], buf[:readLen])  //追加到另一块buffer，仅仅只是为了搜索方便
		readDataLen += readLen
		if (readDataLen<bLen+4){
			continue
		}
		loc := bytes.Index(readData[:readDataLen], boundary)
		if loc>=0{
			//找到了结束位置
			target.Write(readData[:loc-4])
			return readData[loc:readDataLen],reachEnd, nil
		}

		target.Write(readData[:readDataLen-bLen-4])
		copy(readData[0:], readData[readDataLen-bLen-4:])
		readDataLen = bLen + 4
	}
	target.Write(readData[:readDataLen])
	return nil, reachEnd, nil
}

// ParseFromHead 解析表单的头部
/// @param readData 已经从流中读到的数据
/// @param readTotal 已经从流中读到的数据长度
/// @param boundary 表单的分割字符串
/// @param stream 输入流
/// @return FileHeader 文件名等信息头
///			[]byte 已经从流中读到的部分
///			error 是否发生错误
func ParseFromHead(readData []byte, readTotal int, boundary []byte, stream io.ReadCloser)(FileHeader, []byte, error){
	buf := make([]byte, 1024*4)
	foundBoundary := false
	boundaryLoc := -1
	var fileHeader FileHeader
	for {
		readLen, err := stream.Read(buf)
		if err!=nil{
			if err!=io.EOF{
				return fileHeader, nil, err
			}
			break
		}
		if readTotal+readLen>cap(readData){
			return fileHeader, nil, fmt.Errorf("not found boundary")
		}
		copy(readData[readTotal:], buf[:readLen])
		readTotal += readLen
		if !foundBoundary {
			boundaryLoc = bytes.Index(readData[:readTotal], boundary)
			if -1 == boundaryLoc {
				continue
			}
			foundBoundary = true
		}
		startLoc := boundaryLoc+len(boundary)
		fileHeadLoc := bytes.Index(readData[startLoc:readTotal], []byte("\r\n\r\n"))
		if -1==fileHeadLoc{
			continue
		}
		fileHeadLoc += startLoc
		ret := false
		fileHeader,ret = ParseFileHeader(readData[startLoc:fileHeadLoc])
		if !ret{
			return fileHeader,nil,fmt.Errorf("ParseFileHeader fail:%s", string(readData[startLoc:fileHeadLoc]))
		}
		return fileHeader, readData[fileHeadLoc+4:readTotal], nil
	}
	return fileHeader,nil,fmt.Errorf("reach to sream EOF")
}