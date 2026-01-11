# dota的数据

## 查询语句

1. 查询与某个玩家的某个英雄次数（内嵌文档的查询）
db.match.history.find({"players":{"$elemMatch" :{"account_id": 116931565, "hero_id": 1}}.count()

2. 查询某个玩家二人黑次数
db.match.history.find({$and: [{"players.account_id": 116931565}, {"players.account_id": 116931565}]}}).count()

3. 查询某个玩家三人黑次数
db.match.history.find({$and: [{"players.account_id": 116931565}, {"players.account_id": 151533661}, {"players.account_id": 168689546}]}).count()

## 建立索引

db.match.details.ensureIndex({"match_id": 1})
db.players.summaries.ensureIndex({"account_id": 1})