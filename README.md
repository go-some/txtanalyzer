# txtanalyzer
NER, 문서요약 API에 접근하기 위한 go client입니다.

## Usage Example 
```go
package main

import (
        "fmt"
        "github.com/go-some/crawler"
        "github.com/go-some/txtanalyzer"
        "go.mongodb.org/mongo-driver/bson"
        "log"
)

func main() {
        reader := crawler.NewMongoDBReader()

        err := reader.Init()
        if err != nil {
                log.Fatal(err)
        }

        filter := bson.D{{}}
        docs, err := reader.ReadDocs(filter, 10)
        if err != nil {
                log.Fatal(err)
                return
        }

        for _, doc := range docs {
                body := doc.Body
                entList, err := txtanalyzer.RequestNER(body)
                if err != nil {
                        log.Fatal(err)
                        return
                }
                for _, ent := range entList {
                        fmt.Println(ent.Label, ent.Text)
                }
        }

        err = reader.Destroy()
        if err != nil {
                log.Fatal(err)
                return
        }
}
```
