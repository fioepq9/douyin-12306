package service

import (
	"douyin-12306/logger"
)

func init() {
	logger.L.Info("init service success", map[string]interface{}{
		"package":  "service",
		"function": "init",
	})
}

// InfoFlow 数据流接口，使用模板方法模式
type InfoFlow interface {
	checkParam() error
	prepareInfo() error
	packInfo() error
}

// FlowProcessor 流处理器，调用它的Do方法执行下面的校验参数、准备数据和封装结果的流程
type FlowProcessor struct {
	InfoFlow
}

func (producer *FlowProcessor) SetInfoFlow(flow InfoFlow) {
	producer.InfoFlow = flow
}

func (producer FlowProcessor) Do() error {
	// 校验参数
	err := producer.checkParam()
	if err != nil {
		return err
	}
	// 准备数据
	err = producer.prepareInfo()
	if err != nil {
		return err
	}
	// 封装结果
	err = producer.packInfo()
	if err != nil {
		return err
	}
	return nil
}
