package biz

import (
	"fmt"

	pb "github.com/emart.io/zbay/service/go"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

func UUID() string {
	now := types.TimestampNow()
	return fmt.Sprint(now.Seconds) + "-" + fmt.Sprint(now.Nanos)
}

func Connection() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

func DistanceSQL(lon string, lat string) string {
	sqlQuery := `
(SELECT data,
    ROUND(
        6378.138 * 2 * ASIN(
            SQRT(
                POW(
                    SIN(
                        (
                            ` + fmt.Sprint(lat) + ` * PI() / 180 - substring_index(data->>'$.from.location',',', -1) * PI() / 180
                        ) / 2
                    ),
                    2
                ) + COS(` + fmt.Sprint(lat) + ` * PI() / 180) * COS(substring_index(data->>'$.from.location',',', -1) * PI() / 180) * POW(
                    SIN(
                        (
                            ` + fmt.Sprint(lon) + ` * PI() / 180 - substring_index(data->>'$.from.location',',', 1) * PI() / 180
                        ) / 2
                    ),
                    2
                )
            )
        ) * 1000
    ) AS juli
FROM
   orders WHERE data->'$.status' is null 
ORDER BY
	juli ASC ) as orders_juli`
	fmt.Println(sqlQuery)
	return sqlQuery
}

func IsAdmin(user *pb.User) bool {
	if (user.Id == "15901251201") || (user.Id == "13488762245") {
		return true
	}
	return false
}
