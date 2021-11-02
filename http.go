package shelper

import (
    "bytes"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
    "net/http"
    "time"
)

func HttpGet(url string) (string, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }

    client := http.Client{Timeout: time.Second * 5}
    res, err := client.Do(req)
    if err != nil {
        return "", err
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}

func GbkToUtf8(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}
