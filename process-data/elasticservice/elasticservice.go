package elasticservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
)

var tag string = "[Elastic Service]"
var client *elastic.Client 
var index string = "vendas-analysis"
var url string = "https://c01160965fdb47918c3d06933a6df089.us-east-1.aws.found.io:9243"
var user string = "elastic"
var passwd string = "YgcFDLqEuhvxafts9Ulv8KHJ"
var ctx = context.Background()


//Vendas struct insert elasticsearch
type Vendas struct {
	Data string `csv:"data"`
	Escrv int `csv:"escrv"`
	Material int64 `csv:"material"`
	GrpMerc int64 `csv:"grp.merc."`
	QtdFaturd int `csv:"qnt.faturd"`
}

// Init connect ES
func Init(){
	for {

		// Instantiate a new Elasticsearch client object instance
		clientES, err := elastic.NewClient(
			elastic.SetURL(url),
			elastic.SetBasicAuth(user, passwd),
			elastic.SetHealthcheckInterval(1*time.Minute),
			elastic.SetSniff(false),
		)

		if err != nil {

			fmt.Println(tag, "Error trying to create client, trying in 10 seconds...")
			time.Sleep(10 * time.Second)

			continue

		} else {

			// returns information about ElasticSearch
			info, code, err := clientES.Ping(url).Do(ctx)

			if err != nil {
				fmt.Println(tag, "Error trying to get status, trying in 10 seconds...")
				time.Sleep(10 * time.Second)
				continue
			}

			fmt.Println(tag, "ElasticClient initialized - code: ", code, " - version: ", info.Version.Number)

			client = clientES

			if indexError := checkIndex(); indexError == nil {
				// Successfuly initialzed and index checked
				break
			} else {
				fmt.Println(tag, "Error tryin to create indexes ", indexError)
				time.Sleep(5 * time.Second)
				continue
			}

		}
	}


}

// checkIndex checks if an index with the name
/// huawei-log-analysis created in elasticsearch already exists,
// if not it automatically creates it
func checkIndex() error {
	// Use the indexExists service to check if a specified index exists
	exists, err := client.IndexExists(index).Do(ctx)

	if err != nil {
		return err
	}

	// if index does not exist, create a new one
	if !exists {

		createIndex, err := client.CreateIndex(index).Do(ctx)

		if err != nil {
			return err
		}

		if !createIndex.Acknowledged {
			fmt.Println(tag, createIndex)
			return nil
		}

		fmt.Println(tag, "Elasticsearch Successfully created index")
		return nil
	}

	fmt.Println(tag, "Elasticsearch Index already exist")
	return nil

}

//Insertion - insere dados no elastic
func Insertion(data *Vendas) error {

	ctx := context.Background()

	dataJSON, _ := json.Marshal(data)
	js := string(dataJSON)

	if client == nil {
		return errors.New("elastic search client is nil")
	}
	_, err := client.Index().Index(index).BodyJson(js).Do(ctx)

	if err != nil {
		fmt.Println(tag, err)
		return err
	}

	return nil
}
