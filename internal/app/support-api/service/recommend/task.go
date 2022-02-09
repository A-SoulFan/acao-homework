package recommend

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/A-SoulFan/acao-homework/internal/pkg/utility/bilibili"

	"github.com/A-SoulFan/acao-homework/internal/domain"
)

const (
	tempSaveFile = "temp/recommend_slice.csv"
)

type defaultRecommendTask struct {
	recommendRepo domain.RecommendRepo
}

func (rt *defaultRecommendTask) InitTask(ctx context.Context) {
	rt.setRecommendSliceCache()

	rt.startTick()
}

func (rt *defaultRecommendTask) setRecommendSliceCache() {
	// 推荐切片写入 csv 文件
	rt.saveRecommend2CsvFile()

	// 从 csv 文件获取切片排好序并缓存至本地
	rt.cacheRecommend2Local()
}

func (rt *defaultRecommendTask) startTick() {
	tk := time.NewTicker(15 * time.Minute)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				rt.setRecommendSliceCache()

			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(tk)
}

// 从 csv 文件获取切片
func (rt *defaultRecommendTask) getRecommendSliceFromCsvFile() []*domain.RecommendVideo {
	if _, err := os.Stat(tempSaveFile); err != nil {
		return nil
	}

	f, err := os.Open(tempSaveFile)
	if err != nil {
		return nil
	}
	defer f.Close()

	cR := csv.NewReader(f)
	cAll, err := cR.ReadAll()

	if err != nil {
		return nil
	}

	recommendSlice := make([]*domain.RecommendVideo, 0, len(cAll))
	for _, s := range cAll {
		recommendSlice = append(recommendSlice, rt.buildVideo(s))
	}

	return recommendSlice
}

// 从 csv 文件获取切片排好序并缓存至本地
func (rt *defaultRecommendTask) cacheRecommend2Local() {
	// 从 csv 文件获取切片
	recommendSlice := rt.getRecommendSliceFromCsvFile()

	// 按播放量降序
	sort.SliceStable(recommendSlice, func(i, j int) bool {
		iP, err := strconv.Atoi(recommendSlice[i].PlayVal)
		if err != nil {
			panic(err)
		}

		jP, err := strconv.Atoi(recommendSlice[j].PlayVal)
		if err != nil {
			panic(err)
		}

		return iP > jP
	})

	// 缓存
	rt.recommendRepo.SetCache(recommendSlice)
}

func (rt *defaultRecommendTask) buildVideo(csvLen []string) *domain.RecommendVideo {
	video := &domain.RecommendVideo{}

	for i, s := range csvLen {
		switch i {
		case 1:
			video.PlayVal = s
		case 2:
			video.ImageUrl = s
		case 3:
			video.Url = s
		case 4:
			video.Desc = s
		case 5:
			video.Title = s
		case 6:
			video.Auth = s
		case 9:
			video.Time = s
		case 10:
			video.TimeSecond = s
		case 11:
			video.Bid = s
			video.Url = fmt.Sprintf("https://www.bilibili.com/video/%s", s)
		}
	}

	return video
}

// 查询B站推荐切片写入 csv 文件
func (rt *defaultRecommendTask) saveRecommend2CsvFile() error {
	var hitMap = map[string]int{}
	var mids = []string{
		"393396916", // 贾布
		"351609538", // 珈乐
		"672328094", // 嘉然
		"672342685", // 乃琳
		"672346917", // 向晚
		"672353429", // 贝拉
	}

	cvFile, err := os.Create(tempSaveFile)
	if err != nil {
		return err
	}
	defer cvFile.Close()
	_, _ = cvFile.WriteString("\xEF\xBB\xBF") // UTF-8 BOM
	cW := csv.NewWriter(cvFile)

	b := &bilibili.BiliBili{}
	for _, mid := range mids {
		var end = false
		for i := 1; !end; i++ {
			var res *bilibili.SpaceSearchResponse
			var err error
			for n := 0; n < 3; n++ {
				time.Sleep(time.Second * time.Duration(n+1))

				res, err = b.SpaceSearch(mid, 30, i)
				if err != nil {
					if n == 2 {
						fmt.Printf("读取失败, mid:%s ps:%d pn:%d \n %v \n", mid, 30, i, err)
						break
					}
				} else {
					break
				}
			}

			if len(res.List.Vlist) < 30 {
				end = true
			} else if len(res.List.Vlist) == 0 {
				end = true
				break
			}

			for _, v := range res.List.Vlist {
				lengthS := parseTime(v.Length)
				if lengthS > (15 * 60) {
					continue
				}

				// 出道日之前的
				if int64(v.Created) <= time.Date(2020, 12, 11, 0, 0, 0, 0, time.Local).Unix() {
					continue
				}

				// 14天之前的
				if int64(v.Created) <= time.Now().Add(-14*24*time.Hour).Unix() {
					continue
				}

				if _, ok := hitMap[v.Bvid]; ok {
					continue
				} else {
					hitMap[v.Bvid] = 1
				}

				err := cW.Write([]string{
					strconv.Itoa(v.Comment),
					strconv.Itoa(v.Play),
					v.Pic,
					v.Subtitle,
					v.Description,
					v.Title,
					v.Author,
					strconv.Itoa(v.Mid),
					strconv.Itoa(v.Created),
					v.Length,
					strconv.Itoa(lengthS),
					v.Bvid,
				})

				if err != nil {
					fmt.Printf("写入失败, %+v %+v \n", v, err)
				}
			}
		}
	}

	cW.Flush()

	err = cW.Error()
	if err != nil {
		return err
	}
	return nil
}

func parseTime(sTime string) int {
	l := strings.Split(sTime, ":")
	t := 0
	k := 1

	for i := len(l); i > 0; i-- {
		n, _ := strconv.Atoi(l[i-1])
		t += n * k

		k = k * 60
	}

	return t
}
