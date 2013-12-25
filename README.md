# Mail

(A Go Email Utility, still under development...)

## USAGE

```go
mail.From("from@example.com").
To("to1@example.com", "to2@example.com").
Subject("hello").
Body("test").
Body(mail.Body{Value: "<div>hello world</div>", ContentType: "text/html; charset=UTF-8"}).
Attach("report.csv").
Attach(mail.Attachment{FileName: "report2.csv", Content: filebytes}).  // filebytes, _ := ioutil.ReadFile("report.csv")
Send()
```
