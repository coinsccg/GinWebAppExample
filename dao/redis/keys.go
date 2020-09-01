package redis

const (
	AppPrefix = "xxxxxx:"

	KeyQuestionTimeZSet  = AppPrefix + "question:time:"  // 发布时间作为分数的问题ZSet
	KeyQuestionScoreZSet = AppPrefix + "question:score:" // 投票累计作为分数的问题ZSet

	KeyQuestionVotedSetPrefix    = AppPrefix + "question:voted:"    // 某个问题已经投票过的用户Set
	KeyQuestionInfoHashPrefix    = AppPrefix + "question:info:"     // hash,存储问题的基础信息
	KeyCategoryQuestionSetPrefix = AppPrefix + "question:category:" // 问题的分类
)
