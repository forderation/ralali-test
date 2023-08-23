package util

func GetPageCount(pageSize int, totalData int) int {
	pageCount := (1)
	if pageSize > 0 {
		if pageSize >= totalData {
			return pageCount
		}
		if totalData%pageSize == 0 {
			pageCount = totalData / pageSize
		} else {
			pageCount = (totalData / pageSize) + 1
		}
	}
	return pageCount
}
