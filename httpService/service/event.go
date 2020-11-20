package service

var EventMgr *EventManager

func InitEventManager(eventPool chan map[string]interface{}) *EventManager {
	return &EventManager{eventPool: eventPool}
}

type EventManager struct {
	eventPool chan map[string]interface{}
}

//PushEvent 推送事件
func (e *EventManager) PushEvent(eventType, actionType string, actionValue interface{}) {
	data := make(map[string]interface{})
	data["EventType"] = eventType
	data["ActionType"] = actionType
	data["ActionValue"] = actionValue
	e.eventPool <- data
}
