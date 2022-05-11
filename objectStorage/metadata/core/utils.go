package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"metadata/service"
	"time"
)

// 发送到es
func sendToEs(name, hash string, size uint64, version int32) error {
	data := fmt.Sprintf(`{"name":"%s", "version":"%d","length":"%d", "hash":"%s", "time":"%d"}`, name, version+1, size, hash, time.Now().Unix())
	response, err := cli.Index().Index("meta_info").BodyJson(data).Do(context.Background())
	if err != nil {
		log.Println("send metadata to es failed:", err)
		return err
	}
	log.Printf("ID %s to index %s, type %s \n", response.Id, response.Index, response.Type)
	return nil
}

// 获取es的查询结果；
func getResult(result *elastic.SearchResult) (error, []*service.MetaData) {
	var metas []*service.MetaData
	for _, item := range result.Hits.Hits {
		var temp Meta
		var meta service.MetaData
		err := json.Unmarshal(item.Source, &temp)
		if err != nil {
			log.Println(err)
			return err, nil
		}
		meta.Hash = temp.Hash
		meta.Name = temp.Name
		meta.Version = temp.Version
		meta.Length = temp.Length
		metas = append(metas, &meta)
	}
	return nil, metas
}

// 获取最近版本；
func getLatestVersion(name string) []*service.MetaData {
	var metas []*service.MetaData
	qu := elastic.NewQueryStringQuery(fmt.Sprintf("name:%s", name))
	// 使用sort之后，第一条就是最近版本，下面的循环会直接返回结果；
	// bug: 快速请求会连续获得相同的版本号，而实际版本号已经因为PutVersion而改变；
	// es延迟？
	result, err := cli.Search("meta_info").Sort("version.keyword", false).Query(qu).Do(context.Background())
	if err != nil {
		log.Println(err)
		return metas
	}
	for _, item := range result.Hits.Hits {
		var temp Meta
		var meta service.MetaData
		_ = json.Unmarshal(item.Source, &temp)
		meta.Hash = temp.Hash
		meta.Name = temp.Name
		meta.Version = temp.Version
		meta.Length = temp.Length
		metas = append(metas, &meta)
		return metas
	}
	return metas
}

// 获取版本信息
// toDO: version参数的校验需要修改；
func getMetaFromEs(name string, version int32) *elastic.SearchResult {
	if !(version > 0) {
		query := elastic.NewQueryStringQuery(fmt.Sprintf("name:%s", name))
		result, err := cli.Search("meta_info").Query(query).Do(context.Background())
		if err != nil {
			log.Println(err)
			return nil
		}
		return result
	} else {
		query := elastic.NewQueryStringQuery(fmt.Sprintf("version:%d", version))
		result, err := cli.Search("meta_info").Query(query).Do(context.Background())
		if err != nil {
			log.Println(err)
			return nil
		}
		return result
	}
}
