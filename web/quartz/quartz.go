package quartz

import (
	cs "QiqiLike/constants"
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools/redis"
	"encoding/json"
	"github.com/golang-module/carbon"
	"log"
	"sort"
	"strconv"
	"time"
)

const (
	// 权重
	REPLY_NUM  = 5 // 回复数
	OPEN_NUM   = 3 // 打开数
	CREATED_AT = 2 // 当天时间 - 创建时间
)

// 定时抓取热榜到redis
func top() {
	// 算法 评论数和热度0
	db := repositorys.Db
	var articles []domain.TbArticle
	var topSlice domain.TopSlice
	var topNum int

	if _, ok := redis.Exists("top_num"); !ok {
		_ = redis.SetString("top_num", "1")
		topNum = 1
	} else {
		topNum, _ = redis.IncrBy("top_num")
	}

	db.Table("tb_article a").Select("a.uuid, a.title,a.created_at, count(c.id) num ").Joins("LEFT JOIN `tb_comment` c ON a.uuid = c.article_uuid ").Group("a.uuid").Order("num desc, a.created_at desc ").Find(&articles)
	for _, article := range articles {
		now := time.Now()
		num := article.Num * REPLY_NUM
		open := article.OpenNum * OPEN_NUM
		days := carbon.Parse(article.CreatedAt.Format("2006-01-02 15:04:05")).DiffInDays(carbon.Parse(now.Format("2006-01-02 15:04:05"))) * CREATED_AT
		top := domain.TbTop{
			Description:  "这是第" + strconv.Itoa(topNum) + "期热榜",
			CreatedAt:    now,
			Num:          int32(topNum),
			ArticleUuid:  article.Uuid,
			ArticleTitle: article.Title,
			Weight:       int64(num) + int64(open) - days,
		}
		topSlice = append(topSlice, top)
	}
	// 排序
	sort.Sort(topSlice)
	// 取前十条
	if len(topSlice) >= 10 {
		topSlice = topSlice[:10]
	}
	for _, v := range topSlice {
		db.Table("tb_top ").Create(&v)
	}
	bytes, _ := json.Marshal(&topSlice)
	_ = redis.SetString("top", string(bytes))
	log.Printf(cs.INFO+"热榜数据统计完成！当前时间为：%s, 当前期数为：%d\n", time.Now().Format("2006-01-02 15:04:05"), topNum)

}

// 每天一次 吧龄统计
func age() {
	db := repositorys.Db
	var users []domain.TbUser
	db.Model(&domain.TbUser{}).Select("uuid, created_at").Find(&users)
	for _, user := range users {
		days := carbon.Parse(user.CreatedAt.Format("2006-01-02 15:04:05")).DiffInDays(carbon.Parse(time.Now().Format("2006-01-02 15:04:05")))
		age := days / 365
		db.Model(&domain.TbUser{}).Where("uuid = ?", user.Uuid).Update("age", age)
	}
	log.Printf(cs.INFO+"吧龄数据统计完成！当前时间为：%s, 修改条数为：%d\n", time.Now().Format("2006-01-02 15:04:05"), len(users))
}
