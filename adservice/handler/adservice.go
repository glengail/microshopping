package handler

import (
	pb "adservice/proto"
	"context"
	"math/rand"
)

// 最大展示广告数
const MAS_ADS_TO_SERVE = 2

var adsMap = createAdsMap()

type AdService struct{}

func (s *AdService) GetAds(ctx context.Context, in *pb.AdRequest) (out *pb.AdResponse, err error) {
	out = new(pb.AdResponse)
	allAds := make([]*pb.Ad, 0)
	//根据分类顺序获得广告
	if len(in.ContextKeys) > 0 {
		for _, v := range in.ContextKeys {
			ads := s.GetAdsByCategory(v)
			allAds = append(allAds, ads...)
		}
		//如果没有则随机获取
		if len(allAds) == 0 {
			allAds = s.GetRandomAds()
		}
	}

	out = new(pb.AdResponse)
	out.Ads = allAds

	return out, nil
}

// 根据分类获得广告
func (s *AdService) GetAdsByCategory(catagory string) []*pb.Ad {
	return adsMap[catagory]
}
func (s *AdService) GetRandomAds() []*pb.Ad {
	//展示的广告
	ads := make([]*pb.Ad, MAS_ADS_TO_SERVE)
	//获取所有广告
	allAds := make([]*pb.Ad, 0, 7)
	for _, ad := range adsMap {
		allAds = append(allAds, ad...)
	}
	for i := 0; i < MAS_ADS_TO_SERVE; i++ {
		ads = append(ads, allAds[rand.Intn(len(allAds))])
	}
	return ads
}

func createAdsMap() map[string][]*pb.Ad {
	hairdryer := &pb.Ad{RedirectUrl: "/product/93e7a3b3-5a46-4d5e-a5e0-fa0c72137907", Text: "出风机，5折热销"}
	tankTop := &pb.Ad{RedirectUrl: "/product/8c3ef1ef-1ec5-4569-b205-5686314ed802", Text: "背心8折热销"}
	candleHolder := &pb.Ad{RedirectUrl: "/product/7e891c5c-e4b7-4b09-a1e8-944cba78200a", Text: "烛台7折热销"}
	bambooGlassJar := &pb.Ad{RedirectUrl: "/product/d8f22057-ff59-4432-a876-7ea5c0068f59", Text: "竹玻璃罐9折"}
	watch := &pb.Ad{RedirectUrl: "/product/6a39ac09-7125-4a70-8653-e9b0544a3a3b", Text: "手表买一送一"}
	mug := &pb.Ad{RedirectUrl: "/product/ff3d49c1-7f0b-4bce-b274-8c2100518301", Text: "马克杯买二送一"}
	loafers := &pb.Ad{RedirectUrl: "/product/81671bd6-81f6-4514-b3c0-2ebf0aa1a818", Text: "平底鞋，买一送二"}

	return map[string][]*pb.Ad{
		"clothing":    {tankTop},
		"accessories": {watch},
		"footwear":    {loafers},
		"hair":        {hairdryer},
		"decor":       {candleHolder},
		"kitchen":     {bambooGlassJar, mug},
	}
}
