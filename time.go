/**
 * @brief
 * @file time
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package box_lib

import "time"

// GetBeginTimeAndEndTime
// 获取 00:00:00 和 23:59:59 的时间戳
func GetBeginTimeAndEndTime(t time.Time, location *time.Location) (int64, int64) {
	date := t.Format("2006/01/02")
	if len(date) == 0 {
		return 0, 0
	}
	//日期当天 00:00:00 时间戳
	startDate := date + "_00:00:00"
	//日期当天 23:59:59 时间戳
	endDate := date + "_23:59:59"
	if location == nil {
		begin, _ := time.Parse("2006/01/02_15:04:05", startDate)
		end, _ := time.Parse("2006/01/02_15:04:05", endDate)
		//返回当天 00:00:00 和 23:59:59 的时间戳
		return begin.Unix(), end.Unix()
	}
	begin, _ := time.ParseInLocation("2006/01/02_15:04:05", startDate, location)
	end, _ := time.ParseInLocation("2006/01/02_15:04:05", endDate, location)
	//返回当天 00:00:00 和 23:59:59 的时间戳
	return begin.Unix(), end.Unix()
}
