package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"task/model"
	"task/service"
)

// CreateTask 创建Task,将Task信息放入消息队列中；
func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ channel err:" + err.Error())
	}
	queue, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	body, _ := json.Marshal(req)
	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ publish err:" + err.Error())
	}
	return nil
}

// GetTasksList 实现备忘录服务接口 获取备忘录列表
func (*TaskService) GetTasksList(ctx context.Context, req *service.TaskRequest, resp *service.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []model.Task
	var count uint32
	// 查找备忘录
	err := model.DB.Offset(req.Start).Limit(req.Limit).Where("uid=?", req.Uid).Find(&taskData).Error
	if err != nil {
		return errors.New("mysql find : " + err.Error())
	}
	// 统计数量
	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)
	// 返回proto里面定义的类型
	var taskRes []*service.TaskModel
	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = count
	return nil
}

// GetTask 获取详细的备忘录
func (*TaskService) GetTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.First(&taskData, req.Id)
	taskRes := BuildTask(taskData)
	resp.TaskDetail = taskRes
	return nil
}

// UpdateTask 修改备忘录
func (*TaskService) UpdateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskData := model.Task{}
	// 查找该用户的这条信息
	model.DB.Model(&model.Task{}).Where("id= ? AND uid=?", req.Id, req.Uid).First(&taskData)
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content
	model.DB.Save(&taskData)
	resp.TaskDetail = BuildTask(taskData)
	return nil
}

// DeleteTask 删除备忘录
func (*TaskService) DeleteTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	err := model.DB.Model(&model.Task{}).Where("id =? AND uid=?", req.Id, req.Uid).Delete(&model.Task{}).Error
	if err != nil {
		return errors.New("删除失败：" + err.Error())
	}
	return nil
}
