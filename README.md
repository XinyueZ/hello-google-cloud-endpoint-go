#  Learning coding on google cloud with endpoint by golang 

## Reference

- https://github.com/GoogleCloudPlatform/go-endpoints
- https://github.com/googlesamples/cloud-polymer-go

## API

##### This API ```POST```s following struct as ```BODY``` to backend:

```POST https://appname.appspot.com/_ah/api/documents/v1/documents```

```json
{
    "author": "Tester",
    "title": "Jimmy Block has a wunder",
    "content": "hello,world"
}
```

##### This API  ```GET```s a list of recent ```POST```ed struct:


```GET https://appname.appspot.com/_ah/api/documents/v1/documents?limit=[>=0]```

if ```limit<=0``` returns *10* objects as default. 

##### This API  can ```POST``` following ```json``` to search list of objects:

```json
{
  "keyword": "Trump"
}
```

## List of objects if we search

```json
{
 "items": [
  {
   "id": "ahNzfm9yYml0YWwtc3RhZ2UtNjQ4chULEghEb2N1bWVudBiAgICA-MKECgw",
   "author": "TesterTesterTesterTesterTesterTesterTester",
   "content": "TesterTesterTesterTesterTesterTester",
   "title": "Jimmy Block has a wunder",
   "date": "2017-01-26T21:28:02.94061Z",
   "kind": "documents#documentsItem"
  },
  {
   "id": "ahNzfm9yYml0YWwtc3RhZ2UtNjQ4chULEghEb2N1bWVudBiAgICAgICACgw",
   "author": "Tester",
   "content": "hello,world",
   "title": "Jimmy Block has a wunder",
   "date": "2017-01-26T21:24:26.166387Z",
   "kind": "documents#documentsItem"
  }
 ],
 "kind": "documents#documents",
 "etag": "\"MATiH3Txu9Crd2jEnZXdcWIFNBI/-nJWsrajthXliPOORBUw6KiGnvk\""
}
```

##### This API  can also do *update* with body in ```json``` :

*From* data-entry:

```json 
{
   "id": "ag1zfndlYmRldi1wbGF5chULEghEb2N1bWVudBiAgICA67iPCgw",
   "author": "asdf",
   "title": "asdfasdfsadfsadf",
   "date": "2017-01-28T22:31:50.665367Z"
  }
```
*to*

```json
{
   "id": "ag1zfndlYmRldi1wbGF5chULEghEb2N1bWVudBiAgICA67iPCgw",
   "author": "TrumpTrumpTrumpTrump",
   "title": "TrumpwunderTrumpTrumpTrumpTrumpTrump",
   "date": "2017-01-28T22:31:50.665367Z"
  }
```

##### This API  can also do *delete* with body in ```json``` :

```json 
  {
   "id": "ag1zfndlYmRldi1wbGF5chULEghEb2N1bWVudBiAgICA67iPCgw" 
  }
```

