package core

import (
	"encoding/json"
	"time"
)

const dateFormat = "2006-01-02"
const timeFormat = "2006-01-02 15:04:05"

//TsDate 自定义日期格式
type TsDate int64

//TsTime 自定义时间格式
type TsTime int64

func (tsd *TsDate) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}
	tt, _ := time.ParseInLocation(dateFormat, date, time.Local)
	*tsd = TsDate(tt.Unix())
	return nil
}

func (tsd TsDate) MarshalJSON() ([]byte, error) {
	tt := time.Unix(int64(tsd), 0).Format(dateFormat)
	return json.Marshal(tt)
}

func (tst *TsTime) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}
	tt, _ := time.ParseInLocation(timeFormat, date, time.Local)
	*tst = TsTime(tt.Unix())
	return nil
}

func (tst TsTime) MarshalJSON() ([]byte, error) {
	tt := time.Unix(int64(tst), 0).Format(timeFormat)
	return json.Marshal(tt)
}
