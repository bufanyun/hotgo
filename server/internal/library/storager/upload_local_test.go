package storager

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMergePartFile(t *testing.T) {
	srcFile, _ := os.Open("1.zip")
	defer srcFile.Close()
	fileInfo, err := srcFile.Stat()
	if err != nil {
		t.Error(err)
	}
	hash := md5.New()
	_, err = io.Copy(hash, srcFile)
	sum := hash.Sum(nil)
	fileHash := fmt.Sprintf("%x", sum)

	var singleSize = int64(1024 * 1024 * 100)
	chunks := int(fileInfo.Size() / singleSize)
	if fileInfo.Size()%singleSize != 0 {
		chunks += 1
	}
	srcFile.Seek(0, io.SeekStart)
	for j := 0; j < chunks; j++ {
		partSize := singleSize
		if j == chunks-1 {
			partSize = fileInfo.Size() - int64(j)*singleSize
		}
		partData := make([]byte, partSize)
		_, err = io.ReadFull(srcFile, partData)
		pf, _ := os.Create(fmt.Sprintf("tmp/%d", j+1))
		_, err = pf.Write(partData)
		if err != nil {
			t.Error(err)
			return
		}
		pf.Close()
	}

	err = MergePartFile("tmp/", "2.zip")
	if err != nil {
		t.Error(err)
		return
	}

	f2, _ := os.Open("2.zip")
	hash2 := md5.New()
	_, err = io.Copy(hash2, f2)
	sum2 := hash.Sum(nil)
	fileHash2 := fmt.Sprintf("%x", sum2)
	if fileHash != fileHash2 {
		t.Error(errors.New("hash mismatch"))
	}

}
