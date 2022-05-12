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

type InfoFlow interface {
	checkParam() error
	prepareInfo() error
	packInfo() error
}

type FlowProcessor struct {
	InfoFlow
}

func (producer *FlowProcessor) SetInfoFlow(flow InfoFlow) {
	producer.InfoFlow = flow
}

func (producer FlowProcessor) Do() error {
	err := producer.checkParam()
	if err != nil {
		return err
	}
	err = producer.prepareInfo()
	if err != nil {
		return err
	}
	err = producer.packInfo()
	if err != nil {
		return err
	}
	return nil
}

//type InfoFlow interface {
//	checkParam() error
//	prepareInfo() error
//	packInfo() error
//}

//func DoFlow(flow *InfoFlow) error {
//	err := (*flow).checkParam()
//	if err != nil {
//		return err
//	}
//	err = (*flow).prepareInfo()
//	if err != nil {
//		return err
//	}
//	err = (*flow).packInfo()
//	if err != nil {
//		return err
//	}
//	return nil
//}
