package actions

import (
	"github.com/photoview/photoview/api/graphql/models"
	"gorm.io/gorm"
)

//func MyMediaNew(db *gorm.DB, user *models.User, order *models.Ordering, paginate *models.Pagination) ([]*models.Media, error) {
//	database := "photoview"
//	regionId := "cn-hangzhou"
//	secretArn := "acs:rds:cn-hangzhou:2304982093480287:rds-db-credentials/aurora-taKgr1"
//	accountId := "2304982093480287"
//
//	var o client.ExecuteStatementRequest
//	o.AccountId = &accountId
//	o.Database = &database
//	o.RegionId = &regionId
//	o.SecretArn = &secretArn
//
//	var client client.Client
//	var config openapi.Config
//	config.SetAccessKeyId("ACSTQDkNtSMrZtwL")
//	config.SetAccessKeySecret("zXJ7QF79Oz")
//	config.SetEndpoint("rds-data-daily.aliyuncs.com")
//	client.Init(&config)
//}

func MyMedia(db *gorm.DB, user *models.User, order *models.Ordering, paginate *models.Pagination) ([]*models.Media, error) {

	//db的对象表示有一个对数据库的操作，这里是将用户的相册添加到自己的目录中
	//应该就是将album_id和user_id添加到user_albums中
	if err := user.FillAlbums(db); err != nil {
		return nil, err
	}

	query := db.Where("media.album_id IN (SELECT user_albums.album_id FROM user_albums WHERE user_albums.user_id = ?)", user.ID)
	query = models.FormatSQL(query, order, paginate)

	var media []*models.Media
	if err := query.Find(&media).Error; err != nil {
		return nil, err
	}

	return media, nil
}
