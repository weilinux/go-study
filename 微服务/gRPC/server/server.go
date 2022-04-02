package main

import (
	"context"
	"github.com/my/repo/微服务/gRPC/router"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	"time"
)

// 实现routeGuideServer的方法
type routeGuideServer struct {
	// 模拟数据库
	features []*router.Feature
	// 未实现的集成器必须嵌入，才能具有前向兼容的实现
	router.UnimplementedRouteGuideServer
}

// GetFeature unary
func (s *routeGuideServer) GetFeature(ctx context.Context, point *router.Point) (*router.Feature, error) {
	// 循环数据库
	for _, feature := range s.features {
		// 判断两条信息是否相等（pb.point）
		if proto.Equal(feature.Location, point) {
			// 如果相等则返回具体的数据
			return feature, nil
		}
	}

	return nil, nil
}

// check if a point is inside a rectangle
func inRange(point *router.Point, rect *router.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left && float64(point.Longitude) <= right && float64(point.Latitude) >= bottom && float64(point.Latitude) <= top {
		return true
	}
	return false
}

// ListFeatures server side streaming
func (s *routeGuideServer) ListFeatures(rectangle *router.Rectangle, stream router.RouteGuide_ListFeaturesServer) error {
	// 循环数据库
	for _, feature := range s.features {
		// 如果坐标在范围之内
		if inRange(feature.Location, rectangle) {
			// 使用stream将数据发送给客户端
			return stream.Send(feature)
		}
	}
	return nil
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// calcDistance calculates the distance between two points using the "haversine" formula.（计算经纬度）
// The formula is based on http://mathforum.org/library/drmath/view/51879.html.
func calcDistance(p1 *router.Point, p2 *router.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

// RecordRoute client side streaming
func (s *routeGuideServer) RecordRoute(stream router.RouteGuide_RecordRouteServer) error {
	// 总耗时（开始时间）
	startTime := time.Now()
	// 计时器，路线总距离
	var pointCount, distance int32
	// 因为是计算两个点，所以需要保存上一个点
	var prevPoint *router.Point
	for {
		// 等待客户端发来的流式响应(阻塞等待状态)
		point, err := stream.Recv()

		// 如果已发送完毕则将计算数据返回
		if err == io.EOF {
			// 总耗时（结束时间）
			endTime := time.Now()
			// 向客户端发送数据并关闭流
			return stream.SendAndClose(&router.RouteSummary{
				PointCount:  pointCount,
				Distance:    distance,
				ElapsedTime: int32(endTime.Sub(startTime).Seconds()),
			})
		}

		if err != nil {
			return err
		}

		// 客户端每发送一个记录一个
		pointCount++
		// 如果上一个点不为空则计算该点和上一个点的距离
		if prevPoint != nil {
			// 总距离相加
			distance += calcDistance(prevPoint, point)
		}
		// 将该点保存（赋值给上一个点）
		prevPoint = point
	}
}

// 计算距离返回最远或最近的地址（根据request.Mod判断）
func (s *routeGuideServer) recommendOnce(request *router.RecommendationRequest) (*router.Feature, error) {
	var nearest, farthest *router.Feature
	var nearestDistance, farthestDistance int32

	for _, feature := range s.features {
		distance := calcDistance(feature.Location, request.Point)
		if nearest == nil || distance < nearestDistance {
			nearestDistance = distance
			nearest = feature
		}
		if farthest == nil || distance > farthestDistance {
			farthestDistance = distance
			farthest = feature
		}
	}
	if request.Mode == router.RecommendationMode_GetFarthest {
		return farthest, nil
	} else {
		return nearest, nil
	}
}

// Recommend bidirectional streaming
func (s *routeGuideServer) Recommend(stream router.RouteGuide_RecommendServer) error {
	for {
		// 等待客户端发来的流式响应(阻塞等待状态)
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		// 计算距离返回最远和最近的地址
		recommend, err := s.recommendOnce(request)
		if err != nil {
			log.Fatalln(err)
		}

		// 使用stream将数据发送给客户端
		return stream.Send(recommend)
	}
}

func NewServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*router.Feature{
			{
				Name: "上海交通大学闵行校区 上海市闵行区东川路800号",
				Location: &router.Point{
					Latitude:  310235000,
					Longitude: 121437403,
				},
			},
			{
				Name: "复旦大学 上海市杨浦区五角场邯郸路220号",
				Location: &router.Point{
					Latitude:  312978870,
					Longitude: 121503457,
				},
			},
			{
				Name: "华东理工大学 上海市徐汇区梅陇路130号",
				Location: &router.Point{
					Latitude:  311416130,
					Longitude: 121424904,
				},
			},
		},
	}
}

func main() {
	// 服务端创建监听
	lis, err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", "5000"))
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}

	// 创建server
	grpcServer := grpc.NewServer()

	// 将routeGuideServer注册到grpcServer
	router.RegisterRouteGuideServer(grpcServer, NewServer())

	// 将监听传给grpcServer
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
