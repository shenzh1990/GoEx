package common

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	. "github.com/nntaoli-project/goex/v2"
	"github.com/nntaoli-project/goex/v2/internal/logger"
	"github.com/spf13/cast"
)

type RespUnmarshaler struct {
}

func (u *RespUnmarshaler) UnmarshalDepth(data []byte) (*Depth, error) {
	//TODO implement me
	panic("implement me")
}

func (u *RespUnmarshaler) UnmarshalTicker(data []byte) (*Ticker, error) {
	var tk = &Ticker{}

	var open float64
	_, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		err = jsonparser.ObjectEach(value, func(key []byte, val []byte, dataType jsonparser.ValueType, offset int) error {
			valStr := string(val)
			switch string(key) {
			case "last":
				tk.Last = cast.ToFloat64(valStr)
			case "askPx":
				tk.Sell = cast.ToFloat64(valStr)
			case "bidPx":
				tk.Buy = cast.ToFloat64(valStr)
			case "vol24h":
				tk.Vol = cast.ToFloat64(valStr)
			case "high24h":
				tk.High = cast.ToFloat64(valStr)
			case "low24h":
				tk.Low = cast.ToFloat64(valStr)
			case "ts":
				tk.Timestamp = cast.ToInt64(valStr)
			case "open24h":
				open = cast.ToFloat64(valStr)
			}
			return nil
		})
	})

	if err != nil {
		logger.Errorf("[UnmarshalTicker] %s", err.Error())
		return nil, err
	}

	tk.Percent = (tk.Last - open) / open * 100

	return tk, nil
}

func (u *RespUnmarshaler) UnmarshalGetKlineResponse(data []byte) ([]Kline, error) {
	var klines []Kline
	_, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		var (
			k Kline
			i int
		)
		_, err = jsonparser.ArrayEach(value, func(val []byte, dataType jsonparser.ValueType, offset int, err error) {
			valStr := string(val)
			switch i {
			case 0:
				k.Timestamp = cast.ToInt64(valStr)
			case 1:
				k.Open = cast.ToFloat64(valStr)
			case 2:
				k.High = cast.ToFloat64(valStr)
			case 3:
				k.Low = cast.ToFloat64(valStr)
			case 4:
				k.Close = cast.ToFloat64(valStr)
			case 5:
				k.Vol = cast.ToFloat64(valStr)
			}
			i += 1
		})
		klines = append(klines, k)
	})

	return klines, err
}

func (u *RespUnmarshaler) UnmarshalResponse(data []byte, res interface{}) error {
	return json.Unmarshal(data, res)
}