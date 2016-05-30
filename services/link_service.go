package services

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"

	"api.link/models"
)

func GetLink(linkID int) (models.Link, error) {

	c := models.Pool.Get()
	defer c.Close()

	link := models.Link{}
	value, err := redis.Values(c.Do("HGETALL", linkID))
	if err != nil {
		return link, err
	}

	if len(value) > 0 {
		err = redis.ScanStruct(value, &link)
		return link, err
	}

	rows, err := models.DB.Query("select * from link where id = $1", linkID)
	defer rows.Close()
	if err != nil {
		return link, err
	}

	if rows.Next() {
		rows.Scan(&link.Id, &link.WebURL, &link.AppstoreURL, &link.PlayStoreURL)
	} else {
		return link, errors.New("Link with given id not found")
	}

	_, err = c.Do("HMSET", linkID, "id", link.Id, "web_url", link.WebURL, "appstore_url", link.AppstoreURL, "playstore_url", link.PlayStoreURL)
	if err != nil {
		fmt.Print(err)
	}
	return link, err
}

func PostLink(link models.Link) (int, error) {
	var insertID int
	err := models.DB.QueryRow("insert into link (weburl, appstore_url, playstore_url) values($1,$2, $3) returning id", link.WebURL, link.AppstoreURL, link.PlayStoreURL).Scan(&insertID)
	if err != nil {
		return insertID, err
	}

	return insertID, err
}

func PutItem(link models.Link) (bool, error) {
	fmt.Println(link)
	_, err := models.DB.Exec("update link set weburl = $1, appstore_url=$2,playstore_url=$3 where id = $4", link.WebURL, link.AppstoreURL, link.PlayStoreURL, link.Id)
	if err != nil {
		return false, err
	}

	return true, err
}

func DeleteItem(linkID int) (bool, error) {
	_, err := models.DB.Exec("delete from todo where id = $1", linkID)
	if err != nil {
		return false, err
	}
	return true, err
}
