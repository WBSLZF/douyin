package service

import "github.com/RaymondCode/simple-demo/dao"

func RelationAction(user_id int64, to_user_id int64, actionType int64) error {
	if actionType == 1 {
		return dao.Relations{}.Add(user_id, to_user_id)
	}
	return dao.Relations{}.Delete(user_id, to_user_id)
}
