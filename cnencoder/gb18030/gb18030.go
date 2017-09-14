package gb18030

import (
    "bytes"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
)
func Encode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GB18030.NewDecoder()))
    if err == nil {
        dst = string(data)
    }
    return
}
func Decode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GB18030.NewDecoder()))
    if err == nil {
        dst = string(data)
    }
    return
}
