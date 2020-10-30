package main

import "fmt"


type Author struct {
	Name string             //名字
	VIP bool                //是否是高贵的带会员
	Icon string             //头像
	Signature string        //签名
	Focus int               //关注人数
}
type Title struct {
	Name string             //视频标题
	Frequency int           //播放量
	Barrage int             //弹幕
    Top int                 //排名
}
type Video struct {
	A Author
	T Title
}
func main()  {
	video:=Video{
		Author{
			Name:"Mi",
			VIP :true,
			Icon:"Mi Rabbit",
			Signature:"永远相信美好的事情即将发生",
			Focus:1000000,
			},
		Title{
			Name:"Are you ok?",             //视频标题
			Frequency:500000,           //播放量
			Barrage:10000,             //弹幕
			Top:1,                 //排名
		},

	}
	fmt.Println(video)

	//a:=Author{
	//	Name:"Mi",
	//	VIP :true,
	//	Icon:"Rabbit",
	//	Signature:"永远相信美好的事情即将发生",
	//	Focus:1000000,
	//}
	//b:=Title{
	//	Name:"Are you ok?",             //视频标题
	//	Frequency:500000,           //播放量
	//	Barrage:10000,             //弹幕
	//	Top:1,                 //排名
	//	}
}