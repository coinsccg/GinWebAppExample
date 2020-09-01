package redis

import (
	"time"

	"github.com/go-redis/redis"
)

const (
	OneWeekInSeconds = 7 * 24 * 3600
	VoteScore        = 432 // 200 个赞成票能给帖子续一天 86400/200=432
	PerAge           = 20
)

// QuestionVoteUp 为提问投票
func QuestionVoteUp(questionID, userID string) (err error) {
	votedKey := KeyQuestionVotedSetPrefix + questionID
	scoreKey := KeyQuestionScoreZSet
	questionInfoKey := KeyQuestionInfoHashPrefix + questionID
	// 记录投票相关数据
	// 确保当前用户没有给该问题投过票
	if !Client.SIsMember(votedKey, userID).Val() {
		// redis 事务操作
		pipe := Client.TxPipeline()
		pipe.SAdd(votedKey, userID)
		pipe.ZIncrBy(scoreKey, VoteScore, questionID)
		pipe.HIncrBy(questionInfoKey, "votes", 1)
		_, err = pipe.Exec()
		return
	}
	return
}

// CreateQuestion 使用hash存储提问信息
func CreateQuestion(questionID, userID, caption, content string) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyQuestionVotedSetPrefix + questionID
	infoKey := KeyQuestionInfoHashPrefix + questionID
	scoreKey := KeyQuestionScoreZSet
	timeKey := KeyQuestionTimeZSet

	pipeline := Client.Pipeline()
	pipeline.SAdd(votedKey, userID)
	pipeline.Expire(votedKey, time.Second*OneWeekInSeconds)
	fields := map[string]interface{}{
		"caption":     caption,
		"content":     content,
		"question:id": questionID,
		"user:id":     userID,
		"time":        now,
		"votes":       1,
		"comments":    0,
	}
	pipeline.HMSet(infoKey, fields) // 存hash
	pipeline.ZAdd(scoreKey, redis.Z{ // 往问题分数的ZSet添加当前的问题
		Score:  now + VoteScore,
		Member: questionID,
	})
	pipeline.ZAdd(timeKey, redis.Z{ // 往问题时间ZSet添加当前的问题
		Score:  now,
		Member: questionID,
	})
	_, err = pipeline.Exec()
	return
}

// 查询问题排行榜
// 1.按发布时间排序
// 2.按分数(热度)排序

// GetQuestion 从key中分页取出问题
func GetQuestion(key string, page int64) []map[string]string {
	start := (page - 1) * PerAge // 0
	end := start + PerAge - 1    // 19
	ids := Client.ZRevRange(key, start, end).Val()
	list := make([]map[string]string, 0, len(ids))
	pipe := Client.Pipeline()
	for _, id := range ids {
		pipe.HGetAll(KeyQuestionInfoHashPrefix + id)
	}
	c, err := pipe.Exec()
	if err != nil {
		return nil
	}
	for _, cm := range c {
		v, err := cm.(*redis.StringStringMapCmd).Result()
		if err != nil {
			continue
		}
		list = append(list, v)
	}
	return list
}

// GetCategoryQuestion 分时间或者分数取出分页的帖子
func GetCategoryQuestion(categoryName, orderKey string, page int64) []map[string]string {
	key := orderKey + categoryName // 创建缓存键
	categoryKey := KeyCategoryQuestionSetPrefix + categoryName
	// 对categoryKey 和 问题时间 进行取交集计算，并把结果缓存到key中(60秒)
	if Client.Exists(key).Val() < 1 {
		Client.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, categoryKey, orderKey)

		Client.Expire(key, 60*time.Second)
	}
	return GetQuestion(key, page)
}

//GetCategoryQuestion("xuexi", KeyQuestionTimeZSet)  // 学习分区按发布时间排序

//GetCategoryQuestion("qinggan", KeyQuestionScoreZSet)// 情感分区按分数排序
